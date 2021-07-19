# ETH2-DROP


## use

```sh
make build
# after config dropper_conf.toml
./dropperd
# after config ledger_conf.toml
./ledgerd
```

## bind abi

```sh
abigen --abi ./contract/fis_drop/FisDropREth.json --pkg contract_fis_drop --type FisDropREth --out ./contract/fis_drop/FisDropREth.go
```
