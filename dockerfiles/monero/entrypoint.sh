#!/usr/bin/env bash

monerod --testnet --detach --no-igd --hide-my-port --testnet-data-dir ~/.bitmonero/node_02 \
--p2p-bind-ip 0.0.0.0 --log-level 1 \
--testnet-p2p-bind-port 38000 --confirm-external-bind \
--add-exclusive-node 127.0.0.1:28000 --allow-local-ip --non-interactive

monerod --testnet --no-igd --hide-my-port --testnet-data-dir ~/.bitmonero/node_01 \
--p2p-bind-ip 0.0.0.0 --rpc-bind-ip 0.0.0.0 --log-level 1 --testnet-p2p-bind-port 28000 \
--testnet-rpc-bind-port 28001 --confirm-external-bind \
--add-exclusive-node 127.0.0.1:38000 --allow-local-ip --non-interactive \
--start-mining A2m71Lp4dbbcJSDGefheW3a2LBWBddFZfGSqLfVMhin7fYcTGwC2d7eKSBqdGySQShS17RNoAFLtYhtYhcpwRrFs1fwhQbx --mining-threads 1
