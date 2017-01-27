FROM ubuntu:16.04

ENV GOPATH /root/GO

RUN apt-get update
RUN apt-get install -y golang-go git

RUN mkdir -p /root/GO
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/hashicorp/consul/api

COPY app.go /root
RUN cd /root && go build app.go

CMD ["/root/app"]
