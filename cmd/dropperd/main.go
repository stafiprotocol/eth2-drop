package main

import (
	"context"
	"crypto/ecdsa"
	"drop/contract/fis_drop"
	"drop/pkg/config"
	"drop/pkg/log"
	"drop/pkg/utils"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	"github.com/stafiprotocol/chainbridge/utils/keystore"
	"github.com/urfave/cli/v2"
)

const reTryLimit = 30
const waitTime = time.Second * 10

func _main() error {
	cfg, err := config.Load("dropper_conf.toml")
	if err != nil {
		fmt.Printf("loadConfig err: %s", err)
		return err
	}
	log.InitLogFile(cfg.LogFilePath + "/dropper")
	logrus.Infof("config info:%+v ", cfg)

	kpI, err := keystore.KeypairFromAddress(cfg.From, keystore.EthChain, cfg.KeystorePath, false)
	if err != nil {
		return err
	}
	kp, ok := kpI.(*secp256k1.Keypair)
	if !ok {
		return fmt.Errorf("keypair failed")
	}
	logrus.Info("open wallet ok")
	logrus.Info("dropper is running...")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			newDaySeconds := utils.GetNewDayUtc8Seconds()
			if newDaySeconds < cfg.DropTime {
				fmt.Print("waiting droptime...")
				continue
			}
			client, err := ethclient.Dial(cfg.EthApi)
			if err != nil {
				return err
			}
			fisDropContract, err := contract_fis_drop.NewFisDropREth(common.HexToAddress(cfg.DropContract), client)
			if err != nil {
				return err
			}
			callOpts := bind.CallOpts{
				Pending:     false,
				From:        [20]byte{},
				BlockNumber: nil,
				Context:     context.Background(),
			}

			//check date drop,skip if has drop
			nowDate := utils.GetNowUTC8Date()
			dateHash := crypto.Keccak256Hash([]byte(nowDate))
			dateHasDrop, err := fisDropContract.DateDrop(&callOpts, dateHash)
			if err != nil {
				return err
			}
			if dateHasDrop {
				continue
			}
			//check dropflow, skip if no dropflow yesterday
			retry := 0
			dropFlowLatestDate := ""
			for {
				if retry > reTryLimit {
					return fmt.Errorf("getDropFLowLatest  reach retry")
				}
				dropFlowLatestDate, err = getDropFLowLatest(cfg.LedgerApi)
				if err != nil {
					logrus.Warnf("getDropFLowLatest failed: %s", err)
					time.Sleep(waitTime)
					continue
				}
				break
			}
			if dropFlowLatestDate < utils.GetYesterdayUTC8Date() {
				continue
			}

			//get claim state
			claimOpen, err := fisDropContract.ClaimOpen(&callOpts)
			if err != nil {
				return err
			}
			claimRoundOnchain, err := fisDropContract.ClaimRound(&callOpts)
			if err != nil {
				return err
			}

			//check gasprice
			gasPriceMaxLimit := big.NewInt(cfg.MaxGasPrice * 1e9)
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				gasPrice = nil
			} else {
				gasPrice = gasPrice.Add(gasPrice, big.NewInt(20e9))
			}
			if gasPrice.Cmp(gasPriceMaxLimit) > 0 {
				gasPrice = gasPriceMaxLimit
			}

			//txopts
			from := kp.CommonAddress()
			txOpts := &bind.TransactOpts{
				From: from,
				Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
					return signTx(tx, kp.PrivateKey(), cfg.ChainId)
				},
				GasPrice: gasPrice,
				Context:  context.Background(),
			}

			// close claim
			if claimOpen {
				tx, err := fisDropContract.CloseClaim(txOpts)
				if err != nil {
					return err
				}

				//wait close claim tx onchain
				retry := 0
				for {
					if retry > reTryLimit {
						return fmt.Errorf("check ClaimClose tx  reach retry")
					}
					_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
					if err == nil && !isPending {
						break
					} else {
						logrus.Warn("check CloseClaim tx failed ,watting...", " isPending ", isPending, " err ", err)
						time.Sleep(waitTime)
						retry++
						continue
					}
				}

				//wati claim close
				retry = 0
				for {
					if retry > reTryLimit {
						return fmt.Errorf("check ClaimClose tx  reach retry")
					}
					open, err := fisDropContract.ClaimOpen(&callOpts)
					if err == nil && !open {
						break
					} else {
						logrus.Warn("check ClaimOpen failed ,watting...", " err ", err)
						time.Sleep(waitTime)
						retry++
						continue
					}
				}
			}

			//get root hash of new round
			retry = 0
			willUseRootHash := ""
			for {
				if retry > reTryLimit {
					return fmt.Errorf("get root hash  reach retry")
				}
				willUseRootHash, err = getRootHash(cfg.LedgerApi, claimRoundOnchain.Int64()+1)
				if err != nil {
					logrus.Warnf("getRootHash failed: %s", err)
					time.Sleep(waitTime * 2)
					continue
				}
				break
			}

			//send tx set root hash
			tx, err := fisDropContract.SetMerkleRoot(txOpts, dateHash, common.HexToHash(willUseRootHash))
			if err != nil {
				return err
			}
			//wait root hash set tx onchain
			retry = 0
			for {
				if retry > reTryLimit {
					return fmt.Errorf("check SetMerkleRoot tx  reach retry")
				}
				_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
				if err == nil && !isPending {
					break
				} else {
					logrus.Warn("check SetMerkleRoot tx failed ,watting...", " isPending ", isPending, " err ", err)
					time.Sleep(waitTime)
					retry++
					continue
				}
			}

			//wait date drop
			retry = 0
			for {
				if retry > reTryLimit {
					return fmt.Errorf("check dateDrop  reach retry")
				}
				dateDrop, err := fisDropContract.DateDrop(&callOpts, dateHash)
				if err == nil && dateDrop {
					break
				} else {
					logrus.Warn("check dateDrop failed ,watting...", " err ", err)
					time.Sleep(waitTime)
					retry++
					continue
				}
			}

			logrus.Infof("set merkle root hash success, round %d ,root_hash: %s", claimRoundOnchain.Int64()+1, willUseRootHash)

		}
	}
}

func signTx(rawTx *types.Transaction, prv *ecdsa.PrivateKey, chainId int64) (signedTx *types.Transaction, err error) {
	// Sign the transaction and verify the sender to avoid hardware fault surprises
	signer := types.NewEIP155Signer(big.NewInt(chainId))
	signedTx, err = types.SignTx(rawTx, signer, prv)
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
}

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2020 Stafi Protocol Authors"
	app.Name = "dropperd"
	app.Usage = "dropperd"
	app.Authors = []*cli.Author{{Name: "Stafi Protocol 2021"}}
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		&accountCommand,
	}

	app.Flags = append(app.Flags, cliFlags...)
}

func run(ctx *cli.Context) error {
	return _main()
}
