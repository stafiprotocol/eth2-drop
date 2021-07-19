# ETH2-DROP


## use

```sh
make build
# after config dropper_conf.toml
./build/dropperd
# after config ledger_conf.toml
./build/ledgerd
```

## gen keystore
```sh
./build/dropperd accounts geneth --keypath ./
```

## bind abi

```sh
abigen --abi ./contract/fis_drop/FisDropREth.json --pkg contract_fis_drop --type FisDropREth --out ./contract/fis_drop/FisDropREth.go
```
