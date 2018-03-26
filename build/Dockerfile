FROM golang:1.10.0 AS buildbridge

ADD . /go/src/github.com/bonsaiai/nats-streaming-grpc-bridge
WORKDIR /go/src/github.com/bonsaiai/nats-streaming-grpc-bridge
ENV CGO_ENABLED 0
ENV GOOS linux
RUN go install ./pkg/... ./cmd/...

FROM alpine:3.7
RUN mkdir -p /usr/local/bin
COPY --from=buildbridge /go/bin/bridge /usr/local/bin/bridge
CMD [ /usr/local/bin/bridge ]
