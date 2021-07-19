package main

import (
	"context"
	"drop/contract/fis_drop"
	"drop/pkg/config"
	"drop/pkg/log"
	"drop/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func _main() error {
	cfg, err := config.Load("dropper_conf.toml")
	if err != nil {
		fmt.Printf("loadConfig err: %s", err)
		return err
	}
	log.InitLogFile(cfg.LogFilePath + "/dropper")
	logrus.Infof("config info:%+v ", cfg)

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

			//check date drop,skip if drop
			nowDate := utils.GetNowUTC8Date()
			dateHash := crypto.Keccak256Hash([]byte(nowDate))

			dateHasDrop, err := fisDropContract.DateDrop(&callOpts, dateHash)
			if err != nil {
				return err
			}
			if dateHasDrop {
				continue
			}

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
			privKeyBts, err := hexutil.Decode(cfg.Seed)
			if err != nil {
				return err
			}
			_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBts)
			from := crypto.PubkeyToAddress(*pubKey.ToECDSA())

			txOpts := &bind.TransactOpts{
				From: from,
				Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
					return signTx(tx, privKeyBts, cfg.ChainId)
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

				//wait claim closed
				retry := 0
				for {
					if retry > 20 {
						return fmt.Errorf("check ClaimClose tx  reach retry")
					}
					_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
					if err == nil && !isPending {
						break
					} else {
						logrus.Warn("check CloseClaim tx failed ,watting...", " isPending ", isPending, " err ", err)
						time.Sleep(15 * time.Second)
						retry++
						continue
					}
				}

				retry = 0
				for {
					if retry > 20 {
						return fmt.Errorf("check ClaimClose tx  reach retry")
					}
					open, err := fisDropContract.ClaimOpen(&callOpts)
					if err == nil && !open {
						break
					} else {
						logrus.Warn("check ClaimOpen failed ,watting...", " err ", err)
						time.Sleep(15 * time.Second)
						retry++
						continue
					}
				}
			}

			//get new round root hash
			retry := 0
			willUseRootHash := ""
			for {
				if retry > 20 {
					return fmt.Errorf("get root hash  reach retry")
				}
				time.Sleep(10 * time.Second)
				realUrl := fmt.Sprintf("%s/api/v1/root_hash?round=%d", cfg.LedgerApi, claimRoundOnchain.Int64()+1)
				rsp, err := http.Get(realUrl)
				if err != nil {
					logrus.Warnf("get root hash err %s", err)
					continue
				}
				if rsp.StatusCode != http.StatusOK {
					logrus.Warnf("get root hash status: %d", rsp.StatusCode)
					continue
				}
				rspBodyBts, err := ioutil.ReadAll(rsp.Body)
				if err != nil {
					logrus.Warnf("get root hash err %s", err)
					continue
				}
				if len(rspBodyBts) <= 0 {
					logrus.Warnf("get root hash err %s", fmt.Errorf("body zero err"))
					continue
				}
				rspRootHash := RspRootHash{}
				err = json.Unmarshal(rspBodyBts, &rspRootHash)
				if err != nil {
					logrus.Warnf("get root hash err %s", err)
					continue
				}
				if rspRootHash.Status != "80000" {
					logrus.Warnf("get root hash err %s", fmt.Errorf("status err:%s", rspRootHash.Status))
					continue
				}
				willUseRootHash = rspRootHash.Data.RootHash
				break

			}

			//set root hash

			tx, err := fisDropContract.SetMerkleRoot(txOpts, dateHash, common.HexToHash(willUseRootHash))
			if err != nil {
				return err
			}
			//wait root hash set
			retry = 0
			for {
				if retry > 20 {
					return fmt.Errorf("check SetMerkleRoot tx  reach retry")
				}
				_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
				if err == nil && !isPending {
					break
				} else {
					logrus.Warn("check SetMerkleRoot tx failed ,watting...", " isPending ", isPending, " err ", err)
					time.Sleep(15 * time.Second)
					retry++
					continue
				}
			}

			retry = 0
			for {
				if retry > 20 {
					return fmt.Errorf("check dateDrop  reach retry")
				}
				dateDrop, err := fisDropContract.DateDrop(&callOpts, dateHash)
				if err == nil && dateDrop {
					break
				} else {
					logrus.Warn("check dateDrop failed ,watting...", " err ", err)
					time.Sleep(15 * time.Second)
					retry++
					continue
				}
			}

			logrus.Infof("set merkle root hash success, round %d ,root_hash: %s", claimRoundOnchain.Int64()+1, willUseRootHash)

		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	debug.SetGCPercent(40)
	err := _main()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func signTx(rawTx *types.Transaction, privateKeyBts []byte, chainId int64) (signedTx *types.Transaction, err error) {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBts)
	// Sign the transaction and verify the sender to avoid hardware fault surprises
	signer := types.NewEIP155Signer(big.NewInt(chainId))
	signedTx, err = types.SignTx(rawTx, signer, privKey.ToECDSA())
	return
}

type RspRootHash struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		RootHash string `json:"root_hash"`
	} `json:"data"`
}
