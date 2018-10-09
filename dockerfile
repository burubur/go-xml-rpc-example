FROM golang:1.9-alpine as builder

WORKDIR /go/src/rpc-server
COPY ./rpc-server .

# RUN go build

WORKDIR /go/src/rpc-client
COPY ./rpc-client .

# RUN go build

EXPOSE 8800

ENTRYPOINT [ "./rpc-server", "./rpc-client" ]
