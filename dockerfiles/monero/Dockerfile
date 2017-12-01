FROM alpine:latest

# Original Dockerfile https://hub.docker.com/r/szyhf/bitcoin-alpine/~/dockerfile/
MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

ENV MONERO_VERSION=0.11.1.0
ENV GLIBC_VERSION=2.25-r0

RUN adduser -D -g "" cryptodevuser

WORKDIR /home/cryptodevuser

RUN apk update \
	&& apk --no-cache add bash wget ca-certificates bzip2 \
	&& wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/sgerrand.rsa.pub \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk \
 	&& apk --no-cache add glibc-${GLIBC_VERSION}.apk \
	&& apk --no-cache add glibc-bin-${GLIBC_VERSION}.apk \
	&& wget https://github.com/monero-project/monero/releases/download/v${MONERO_VERSION}/monero-linux-x64-v${MONERO_VERSION}.tar.bz2 \
	&& tar -jxvf monero-linux-x64-v${MONERO_VERSION}.tar.bz2 \
	&& mkdir .bitmonero \
	&& apk del wget ca-certificates bzip2


RUN mv monero*/* /usr/local/bin/ \
	&& rm -rf monero* \
	&& rm -rf glibc-${GLIBC_VERSION}.apk \
	&& rm -rf glibc-bin-${GLIBC_VERSION}.apk

COPY bitcoin.conf .bitmonero

COPY entrypoint.sh /usr/local/bin

RUN chown -R cryptodevuser:cryptodevuser /home/cryptodevuser/.bitmonero

USER cryptodevuser

EXPOSE 28000 28001

CMD ["entrypoint.sh"]
