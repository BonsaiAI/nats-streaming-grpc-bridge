package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/bonsaiai/nats-streaming-grpc-bridge/bridge"
	"github.com/google/uuid"
	"github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	defaultNATSServer = "nats://nats:4222"
	defaultClusterID  = "queue"
)

type configuration struct {
	serverURL      string
	clusterID      string
	clientID       string
	verboseLogging bool
	port           int
}

type bridgeService struct {
	bridge.BridgeServer
	connection stan.Conn
}

func randomClientID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	clientID := fmt.Sprintf("bridge-%s", u.String())
	return clientID
}

func newBridgeService(cfg *configuration) (*bridgeService, error) {
	bridge := bridgeService{}
	log.Debugf("Connecting to NATS at %s...", cfg.serverURL)
	conn, err := stan.Connect(cfg.clusterID, cfg.clientID, stan.NatsURL(cfg.serverURL))
	if err != nil {
		log.Errorf("Connection failed - %s", err.Error())
		return nil, err
	}
	bridge.connection = conn

	return &bridge, nil
}

func (s *bridgeService) close() {
	s.connection.Close()
}

var emptyPublishResponse = bridge.PublishResponse{}

func (s *bridgeService) Publish(ctx context.Context, req *bridge.PublishRequest) (*bridge.PublishResponse, error) {
	err := s.connection.Publish(req.Subject, req.Data)
	if err != nil {
		return nil, err
	}
	return &emptyPublishResponse, nil
}

func (s *bridgeService) Subscribe(req *bridge.SubscribeToStreamRequest, stream bridge.Bridge_SubscribeToStreamServer) error {
	log.Debugf("Subscription requested to %s...", req.Subject)
	subscriptionOptions := []stan.SubscriptionOption{}
	switch req.SubscriptionType {
	case bridge.SubscribeToStreamRequest_DELIVER_ALL_AVAILABLE:
		subscriptionOptions = append(subscriptionOptions, stan.DeliverAllAvailable())
	case bridge.SubscribeToStreamRequest_START_WITH_LAST_RECEIVED:
		subscriptionOptions = append(subscriptionOptions, stan.StartWithLastReceived())
	case bridge.SubscribeToStreamRequest_START_AT_SEQUENCE:
		subscriptionOptions = append(subscriptionOptions, stan.StartAtSequence(req.StartAtSequence))
	case bridge.SubscribeToStreamRequest_START_AT_TIME:
		startTime := time.Unix(req.StartAtTime.Seconds, int64(req.StartAtTime.Nanos))
		subscriptionOptions = append(subscriptionOptions, stan.StartAtTime(startTime))
	case bridge.SubscribeToStreamRequest_START_AT_TIME_DELTA:
		subscriptionOptions = append(subscriptionOptions, stan.StartAtTimeDelta(time.Duration(req.StartAtTimeDeltaNs)))
	}

	var subscription stan.Subscription
	var err error

	handler := func(queueMsg *stan.Msg) {
		msg := bridge.Msg{
			Sequence:    queueMsg.Sequence,
			Subject:     queueMsg.Subject,
			Reply:       queueMsg.Reply,
			Data:        queueMsg.Data,
			Timestamp:   queueMsg.Timestamp,
			Redelivered: queueMsg.Redelivered,
			Crc32:       queueMsg.CRC32,
		}
		e := stream.Send(&msg)
		if e != nil {
			log.Warnf("Error sending message %d on %s - %s", queueMsg.Sequence, queueMsg.Subject, e.Error())
		}
	}

	if len(req.QueueGroup) > 0 {
		subscription, err = s.connection.QueueSubscribe(req.Subject, req.QueueGroup, handler, subscriptionOptions...)
		if err != nil {
			return err
		}
	} else {
		subscription, err = s.connection.Subscribe(req.Subject, handler, subscriptionOptions...)
		if err != nil {
			return err
		}
	}

	defer subscription.Unsubscribe()

	<-stream.Context().Done()

	return nil
}

func parseCommandLine() *configuration {
	cfg := configuration{}
	flag.StringVar(&cfg.serverURL, "server", defaultNATSServer, "Comma separated list of NATS URIs to connect to")
	flag.StringVar(&cfg.clusterID, "clusterID", defaultClusterID, "NATS Cluser ID to connect to")
	flag.StringVar(&cfg.clientID, "clientID", randomClientID(), "Client ID used to identify the bridge")
	flag.BoolVar(&cfg.verboseLogging, "verbose", false, "Whether or not to enable verbose logging")
	flag.IntVar(&cfg.port, "port", 5000, "Port to bind to")

	flag.Parse()

	return &cfg
}

func main() {

	cfg := parseCommandLine()
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	if cfg.verboseLogging {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Staring nats-streaming-grpc-bridge...")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.port))
	if err != nil {
		panic(err)
	}

	service, err := newBridgeService(cfg)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	bridge.RegisterBridgeServer(server, service)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}

	log.Info("Bridge Out.")

}
