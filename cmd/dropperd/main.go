// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"context"
	"crypto/ecdsa"
	"drop/pkg/config"
	"drop/pkg/log"
	"drop/pkg/utils"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	"github.com/stafiprotocol/chainbridge/utils/keystore"
	"github.com/urfave/cli/v2"
)

const reTryLimit = math.MaxInt32
const waitTime = time.Second * 10

var callOpts = bind.CallOpts{
	Pending:     false,
	From:        [20]byte{},
	BlockNumber: nil,
	Context:     context.Background(),
}

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

	ticker := time.NewTicker(time.Duration(cfg.TaskTicker) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			//1 wait droptime, skip if now < droptime
			newDaySeconds := utils.GetNewDayUtc8Seconds()
			if newDaySeconds < cfg.DropTime {
				continue
			}

			//dial client
			client := waitClient(cfg.EthApi)

			//create fisDropContract
			fisDropContract := waitFisDropContract(cfg.DropContract, client)
			nowDate := utils.GetNowUTC8Date()
			dateHash := crypto.Keccak256Hash([]byte(nowDate))

			//2 check date drop, skip if has drop today
			dateHasDrop := waitDateHashDrop(fisDropContract, dateHash)
			if dateHasDrop {
				continue
			}

			//3 check dropflowLatestDate, skip if no dropflow yesterday
			dropFlowLatestDate := waitDropFlowLatestDate(cfg.LedgerApi)
			if dropFlowLatestDate < utils.GetYesterdayUTC8Date() {
				continue
			}
			//4 check skip date, skip if skipDate==today
			skipDate := waitToGetSkipDate(cfg.LedgerApi)
			if skipDate == utils.GetNowUTC8Date() {
				continue
			}

			//get claim open state
			claimOpen := waitToGetClaimOpen(fisDropContract)

			//get claim round
			claimRoundOnchain := waitToGetClaimRound(fisDropContract)

			//check gasprice
			gasPriceMaxLimit := big.NewInt(cfg.MaxGasPrice * 1e9)
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				gasPrice = gasPriceMaxLimit
			} else {
				gasPrice = gasPrice.Add(gasPrice, big.NewInt(10e9))
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
				//send close claim tx
				sendCloseClaimTxAndWait(client, fisDropContract, txOpts)
				//wait until claim close
				waitUntilClaimClose(fisDropContract)
			}

			//get root hash of new round
			willUseRootHash, shouldSkipToday := waitToGetRootHash(cfg.LedgerApi, claimRoundOnchain.Int64()+1)
			if shouldSkipToday {
				//send tx to open claim
				sendOpenClaimTxAndWait(client, fisDropContract, txOpts)
			} else {
				//send tx to set root hash
				sendSetRootHashTxAndWait(client, fisDropContract, txOpts, dateHash, common.HexToHash(willUseRootHash))
			}

			//wait claim open
			waitUntilClaimOpen(fisDropContract)
			if shouldSkipToday {
				logrus.Infof("no need set root hash, open claim success, round %d", claimRoundOnchain.Int64())
			} else {
				logrus.Infof("set merkle root hash success, round %d ,root_hash: %s", claimRoundOnchain.Int64()+1, willUseRootHash)
			}

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
