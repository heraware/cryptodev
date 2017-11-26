FROM alpine:latest

MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

RUN adduser -D -g "" cryptodevuser

WORKDIR /home/cryptodevuser

RUN apk update \
	&& apk --no-cache add wget ca-certificates \
	&& wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/sgerrand.rsa.pub \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk \
 	&& apk --no-cache add glibc-${GLIBC_VERSION}.apk \
	&& apk --no-cache add glibc-bin-${GLIBC_VERSION}.apk \
	&& wget https://github.com/dashpay/dash/releases/download/v0.12.1.4/dashcore-0.12.1.4-linux64.tar.gz \
	&& tar -xvf dashcore-*.tar.gz \
	&& cp dashcore*/bin/*  /usr/local/bin \
	&& rm -rf dashcore* \
	&& mkdir .dashcore

COPY dash.conf .dashcore

RUN chown -R cryptodevuser:cryptodevuser /home/cryptodevuser/.dashcore

USER cryptodevuser

EXPOSE 27000 27001

CMD ["dashd"]
