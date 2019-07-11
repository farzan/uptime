
FROM golang:1.12.7-stretch

RUN go get github.com/astaxie/beego \
    && go get github.com/beego/bee \
    && go get golang.org/x/crypto/bcrypt

WORKDIR /go/src

COPY . UptimeMonitor/

WORKDIR /go/src/UptimeMonitor

EXPOSE 8088

CMD ["bee", "run"]

