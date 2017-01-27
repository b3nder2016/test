FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get install -y wget unzip

RUN mkdir -p /opt/consul

RUN cd /opt/consul && \
    wget https://releases.hashicorp.com/consul/0.7.2/consul_0.7.2_linux_amd64.zip && \
    unzip consul_0.7.2_linux_amd64.zip && \
    rm -f unzip consul_0.7.2_linux_amd64.zip

COPY consul.conf /etc

EXPOSE 8500

ENTRYPOINT ["/opt/consul/consul", "agent", "-config-file=/etc/consul.conf"]
