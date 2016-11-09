FROM alpine

COPY . /src

EXPOSE 8080

ENTRYPOINT ["/src/main"]
