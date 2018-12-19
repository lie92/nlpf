
FROM google/golang:1.11

RUN go get github.com/revel/cmd/revel

VOLUME ["/gopath/src"]

WORKDIR /gopath/src

CMD revel run github.com/revel/revel/samples/chat

EXPOSE 9000