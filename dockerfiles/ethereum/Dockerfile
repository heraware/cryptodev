FROM ethereum/client-go

# Original Dockerfile https://hub.docker.com/r/ethereum/client-go/~/dockerfile/
MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

RUN adduser -D -g "" cryptodevuser

RUN apk update \
	&& apk --no-cache add bash

WORKDIR /home/cryptodevuser

USER cryptodevuser

COPY genesis.json .

COPY genesis_alloc.json .

COPY password.txt .

COPY set_genesis.sh .

COPY geth_mine.js .

COPY entrypoint.sh .

RUN geth  --datadir . init genesis.json

RUN geth --datadir . --networkid 99 --password password.txt account new

RUN bash set_genesis.sh

RUN cat genesis_alloc.json

RUN rm -rf geth/

RUN geth  --datadir . init genesis_alloc.json

RUN geth --datadir . --networkid 99 --exec "eth.getBalance(eth.coinbase)" console

ENTRYPOINT ["./entrypoint.sh"]
