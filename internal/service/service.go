package service

import (
	"time"

	//"github.com/bonsaiai/nats-streaming-grpc-bridge/pkg/bridge"
	"github.com/bonsaiai/nats-streaming-grpc-bridge/internal/bridge"
	"github.com/nats-io/go-nats-streaming"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	context "golang.org/x/net/context"
)

// Service is a structure containing the required interfaces for bridging a nats-streaming server and GRPC streams.
type Service struct {
	bridge.BridgeServer
	connection                     stan.Conn
	subscriptionPerformanceMetrics prometheus.Histogram
	subscriptionCountMetrics       prometheus.Gauge
	publishPerformanceMetrics      prometheus.Histogram
}

// NewBridgeService returns an implementation of the BridgeServer service.
func NewBridgeService(conn stan.Conn, subscribePerformanceMetrics prometheus.Histogram, subscribeCountMetrics prometheus.Gauge, publishPerformanceMetrics prometheus.Histogram) *Service {
	bridge := Service{
		connection:                     conn,
		subscriptionPerformanceMetrics: subscribePerformanceMetrics,
		subscriptionCountMetrics:       subscribeCountMetrics,
		publishPerformanceMetrics:      publishPerformanceMetrics,
	}

	return &bridge
}

func (s *Service) close() {
	s.connection.Close()
}

var emptyPublishResponse = bridge.PublishResponse{}

func recordTimingObservation(start time.Time, histogram prometheus.Histogram) {
	elapsed := time.Since(start)
	histogram.Observe(elapsed.Seconds())
}

// Publish pushes a message onto a nat-streaming queue.
func (s *Service) Publish(ctx context.Context, req *bridge.PublishRequest) (*bridge.PublishResponse, error) {
	defer recordTimingObservation(time.Now(), s.publishPerformanceMetrics)
	err := s.connection.Publish(req.Subject, req.Data)
	if err != nil {
		return nil, err
	}
	return &emptyPublishResponse, nil
}

// SubscribeToStream subscribes to a nats-streaming queue.
func (s *Service) SubscribeToStream(req *bridge.SubscribeToStreamRequest, stream bridge.Bridge_SubscribeToStreamServer) error {
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
	msgCount := 0

	handler := func(queueMsg *stan.Msg) {

		defer recordTimingObservation(time.Now(), s.subscriptionPerformanceMetrics)

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
		msgCount++
		if msgCount%10000 == 0 {
			log.Debugf("%d messages received so far for subject %s", msgCount, req.Subject)
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

	s.subscriptionCountMetrics.Inc()

	<-stream.Context().Done()

	s.subscriptionCountMetrics.Dec()

	log.Debugf("Subscription %s finished with %d messages received", req.Subject, msgCount)

	return nil
}
