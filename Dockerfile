FROM golang:1.7.3


RUN go get github.com/Sirupsen/logrus &&\
    go get github.com/docker/docker/client &&\
    go get github.com/sakeven/go-env &&\
    go get golang.org/x/net/context

COPY . /src
WORKDIR /src

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]
