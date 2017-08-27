#!/usr/bin/env bash

geth --datadir . --password password.txt --chain cryptodev --unlock 0 \
 --port 30003 --networkid 99 --rpc --rpcaddr 0.0.0.0 --rpcport 8546 \
 --rpcapi admin,debug,eth,miner,net,personal,shh,txpool,web3 \
 js geth_mine.js
