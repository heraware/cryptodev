FROM alpine:latest

# Original Dockerfile https://hub.docker.com/r/szyhf/bitcoin-alpine/~/dockerfile/
MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

ENV BITCOIN_VERSION=0.15.0
ENV GLIBC_VERSION=2.25-r0

RUN adduser -D -g "" cryptodevuser

WORKDIR /home/cryptodevuser

RUN apk update \
	&& apk --no-cache add wget ca-certificates unzip \
	&& wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/sgerrand.rsa.pub \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk \
 	&& apk --no-cache add glibc-${GLIBC_VERSION}.apk \
	&& apk --no-cache add glibc-bin-${GLIBC_VERSION}.apk \
	&& wget https://github.com/BTCGPU/BTCGPU/releases/download/0.15.0.1/bitcoingold-0.15.0-x86_64-pc-linux-gnu.zip \
	&& unzip -d bitcoingold-${BITCOIN_VERSION} bitcoingold-${BITCOIN_VERSION}-x86_64-pc-linux-gnu.zip \
	&& mkdir .bitcoin \
	&& apk del wget ca-certificates \
	&& mv bitcoingold-${BITCOIN_VERSION}/bin/* /usr/local/bin/ \
	&& rm -rf bitcoingold-${BITCOIN_VERSION}/ \
	&& rm -rf bitcoingold-${BITCOIN_VERSION}-x86_64-pc-linux-gnu.zip \
	&& rm -rf glibc-${GLIBC_VERSION}.apk \
	&& rm -rf glibc-bin-${GLIBC_VERSION}.apk

COPY bitcoin.conf .bitcoin

RUN chown -R cryptodevuser:cryptodevuser /home/cryptodevuser/.bitcoin

USER cryptodevuser

EXPOSE 23000 23001

CMD ["bitcoind"]
