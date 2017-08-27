#!/usr/bin/env bash

geth --datadir . --password password.txt --unlock 0 --networkid 99 --rpc \
 --rpcaddr 0.0.0.0 --rpcapi admin,debug,eth,miner,net,personal,shh,txpool,web3 \
 js geth_mine.js
