package start

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
)

const (
	defaultNATSServer = "nats://nats:4222"
	defaultClusterID  = "queue"
)

// Configuration is the configuration parameters for the daemon application.
type Configuration struct {
	ServerURL      string
	ClusterID      string
	ClientID       string
	VerboseLogging bool
	Port           int
	MetricsEnabled bool
	MetricsPort    int
}

func randomClientID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	clientID := fmt.Sprintf("bridge-%s", u.String())
	return clientID
}

// ParseCommandLine generates a Configuration from command line parameters.
func ParseCommandLine() *Configuration {
	cfg := Configuration{}
	flag.StringVar(&cfg.ServerURL, "server", defaultNATSServer, "Comma separated list of NATS URIs to connect to")
	flag.StringVar(&cfg.ClusterID, "clusterID", defaultClusterID, "NATS Cluser ID to connect to")
	flag.StringVar(&cfg.ClientID, "clientID", randomClientID(), "Client ID used to identify the bridge")
	flag.BoolVar(&cfg.VerboseLogging, "verbose", false, "Whether or not to enable verbose logging")
	flag.IntVar(&cfg.Port, "port", 5000, "Port to bind to")
	flag.BoolVar(&cfg.MetricsEnabled, "metricsEnabled", true, "Whether or not the metrics scrape endpoint is enabled")
	flag.IntVar(&cfg.MetricsPort, "metricsPort", 9000, "Port to bind to for metrics scrapers")

	flag.Parse()

	return &cfg
}
