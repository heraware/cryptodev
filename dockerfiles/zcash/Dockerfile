FROM ubuntu:16.04

MAINTAINER Yohan Graterol <yohangraterol92@gmail.com>

RUN adduser --disabled-password --gecos "" cryptodevuser

WORKDIR /home/cryptodevuser

RUN apt-get update \
		&& apt-get upgrade -y \
		&& apt-get install -y apt-transport-https wget \
		&& wget -qO - https://apt.z.cash/zcash.asc | apt-key add - \
		&& echo "deb [arch=amd64] https://apt.z.cash/ jessie main" | tee /etc/apt/sources.list.d/zcash.list \
		&& apt-get update && apt-get install -y zcash \
		&& zcash-fetch-params

RUN mkdir .zcash

COPY zcash.conf .zcash

RUN mv /root/.zcash-params /home/cryptodevuser

RUN chown -R cryptodevuser:cryptodevuser /home/cryptodevuser

USER cryptodevuser

EXPOSE 25000 25001

CMD ["zcashd"]
