package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bonsaiai/nats-streaming-grpc-bridge/internal/metrics"
	"github.com/bonsaiai/nats-streaming-grpc-bridge/internal/service"
	"github.com/bonsaiai/nats-streaming-grpc-bridge/internal/start"

	"github.com/bonsaiai/nats-streaming-grpc-bridge/internal/bridge"
	"github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	cfg := start.ParseCommandLine()
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	if cfg.VerboseLogging {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Starting bridge...")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}

	conn, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(cfg.ServerURL))
	if err != nil {
		panic(err)
	}

	subscribePerformanceMetrics := metrics.ConstructSubscribePerformanceHistogram(cfg.MetricsEnabled)
	subscribeCountMetrics := metrics.ConstructSubscriberCountGauge(cfg.MetricsEnabled)
	publishPerformanceMetrics := metrics.ConstructPublishPerformanceHistogram(cfg.MetricsEnabled)
	metrics.Serve(cfg.MetricsEnabled, cfg.MetricsPort)

	service := service.NewBridgeService(conn, subscribePerformanceMetrics, subscribeCountMetrics, publishPerformanceMetrics)

	server := grpc.NewServer()
	bridge.RegisterBridgeServer(server, service)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}

	log.Info("Bridge Out.")

}
