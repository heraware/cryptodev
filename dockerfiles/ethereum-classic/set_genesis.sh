#!/usr/bin/env bash

coinbase=$(geth  --datadir . --chain=cryptodev --rpc --rpcaddr 0.0.0.0 --exec eth.coinbase console)
echo "Coinbase: $coinbase"
sed -i s/\"ADDRESS\"/$coinbase/g genesis_alloc.json
echo "Coinbase written in genesis.json file."
