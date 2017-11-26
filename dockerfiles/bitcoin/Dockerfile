FROM alpine:latest

# Original Dockerfile https://hub.docker.com/r/szyhf/bitcoin-alpine/~/dockerfile/
MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

ENV BITCOIN_VERSION=0.15.1
ENV GLIBC_VERSION=2.25-r0

RUN adduser -D -g "" cryptodevuser

WORKDIR /home/cryptodevuser

RUN apk update \
	&& apk --no-cache add wget ca-certificates \
	&& wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/sgerrand.rsa.pub \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk \
 	&& apk --no-cache add glibc-${GLIBC_VERSION}.apk \
	&& apk --no-cache add glibc-bin-${GLIBC_VERSION}.apk \
	&& wget https://bitcoin.org/bin/bitcoin-core-${BITCOIN_VERSION}/bitcoin-${BITCOIN_VERSION}-x86_64-linux-gnu.tar.gz \
	&& tar xzvf bitcoin-${BITCOIN_VERSION}-x86_64-linux-gnu.tar.gz \
	&& mkdir .bitcoin \
	&& apk del wget ca-certificates \
	&& mv bitcoin-${BITCOIN_VERSION}/bin/* /usr/local/bin/ \
	&& rm -rf bitcoin-${BITCOIN_VERSION}/ \
	&& rm -rf bitcoin-${BITCOIN_VERSION}-x86_64-linux-gnu.tar.gz \
	&& rm -rf glibc-${GLIBC_VERSION}.apk \
	&& rm -rf glibc-bin-${GLIBC_VERSION}.apk

COPY bitcoin.conf .bitcoin

RUN chown -R cryptodevuser:cryptodevuser /home/cryptodevuser/.bitcoin

USER cryptodevuser

EXPOSE 20000 20001

CMD ["bitcoind"]
