FROM golang:alpine

# Original Dockerfile https://hub.docker.com/r/ethereum/client-go/~/dockerfile/
MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

RUN adduser -D -g "" cryptodevuser

RUN mkdir -p /go/src/github.com/ethereumproject/

WORKDIR /go/src/github.com/ethereumproject/

RUN set -ex \
	&& apk add --no-cache git build-base gcc abuild bash  && \
	git clone --depth 1 https://github.com/ethereumproject/go-ethereum.git && \
  cd go-ethereum && \
  go get -t -d ./... && \
  go build ./cmd/geth && \
	cp geth /usr/local/bin/ && \
  cd .. && \
  rm -rf go-ethereum

WORKDIR /home/cryptodevuser

USER cryptodevuser

COPY genesis.json .

COPY genesis_alloc.json .

COPY password.txt .

COPY set_genesis.sh .

COPY geth_mine.js .

COPY entrypoint.sh .

RUN mkdir cryptodev/

# RUN mv genesis.json cryptodev/chain.json


# RUN bash set_genesis.sh

# RUN rm -rf cryptodev/
#
# RUN mkdir cryptodev/
#
# RUN cat genesis_alloc.json

RUN mv genesis_alloc.json cryptodev/chain.json

RUN geth --datadir . --chain=cryptodev --networkid 99 --password password.txt account new

RUN geth --datadir . --chain=cryptodev --networkid 99 --exec "eth.getBalance(eth.coinbase)" console

EXPOSE 8546 30003

ENTRYPOINT ["./entrypoint.sh"]
