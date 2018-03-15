FROM golang:1.10.0 AS buildbridge

ADD . /go/src/github.com/bonsaiai/nats-streaming-grpc-bridge
WORKDIR /go/src/github.com/bonsaiai/nats-streaming-grpc-bridge
RUN go get github.com/google/uuid && \	
    go get github.com/nats-io/go-nats-streaming && \
	go get github.com/sirupsen/logrus && \
	go get google.golang.org/grpc && \
    CGO_ENABLED=0 GOOS=linux go install

FROM alpine:3.7
COPY --from=buildbridge /go/bin/nats-streaming-grpc-bridge /nats-streaming-grpc-bridge
CMD [ /nats-streaming-grpc-bridge ]
