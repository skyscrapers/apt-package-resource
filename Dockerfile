FROM debian:jessie

RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 9E3E53F19C7DE460 && \
    apt-get update && \
    apt-get install -y aptly && \
    rm -rf /var/lib/apt/lists/*

COPY check /opt/concourse/
COPY in /opt/concourse/
COPY out /opt/concourse/
COPY apt-package-resource /opt/concourse/
