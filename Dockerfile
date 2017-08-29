FROM debian:jessie

RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 9E3E53F19C7DE460 && \
    apt-get update && \
    apt-get install -y jq aptly && \
    rm -rf /var/lib/apt/lists/* && \
    gpg --keyring /usr/share/keyrings/debian-archive-keyring.gpg --export | gpg --keyring trustedkeys.gpg --import

COPY check /opt/resource/
COPY in /opt/resource/
COPY out /opt/resource/
COPY apt-package-resource /opt/resource/
