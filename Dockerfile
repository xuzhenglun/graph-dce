FROM golang:1.7.3

ADD . $GOPATH/src

RUN go get graph-dce/...
RUN go install graph-dce

EXPOSE 8080

ENTRYPOINT ["graph-dce"]
