FROM ubuntu:16.04

MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

RUN adduser --disabled-password --gecos "" cryptodevuser

WORKDIR /home/cryptodevuser

RUN apt-get update \
		&& apt-get upgrade -y \
		&& apt-get install -y apt-transport-https wget \
		&& apt-get install -y build-essential pkg-config libc6-dev m4 g++-multilib autoconf libtool ncurses-dev unzip git python zlib1g-dev wget bsdmainutils automake && apt-get autoremove -y \
		&& git clone https://github.com/ZencashOfficial/zen.git \
		&& ./zen/zcutil/fetch-params.sh

RUN mkdir .zen

COPY zen.conf .zen

RUN ./zen/zcutil/build.sh

RUN mv /root/.zcash-params /home/cryptodevuser

RUN chown -R cryptodevuser:cryptodevuser /home/cryptodevuser

USER cryptodevuser

EXPOSE 26000 26001

CMD ["/home/cryptodevuser/zen/src/zend"]
