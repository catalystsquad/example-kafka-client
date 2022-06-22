package main

import (
	"os"

	"github.com/catalystsquad/app-utils-go/env"
)

// KafkaConfig contains all configuration options for producers and consumers
// to pass to example impelemtation functions
type KafkaConfig struct {
	BootstrapServers         string
	Topic                    string
	GroupID                  string
	OffsetReset              bool
	TLSCertPath              string
	TLSKeyPath               string
	CaCertPath               string
	SecurityProtocol         string
	EnableSSLCerterification bool
	MessageCount             int
}

func main() {
	cfg := KafkaConfig{
		BootstrapServers:         env.GetEnvOrDefault("MYKAFKA_BOOTSTRAP_SERVERS", "localhost:9092"),
		Topic:                    env.GetEnvOrDefault("MYKAFKA_TOPIC", "tuttopic"),
		GroupID:                  env.GetEnvOrDefault("MYKAFKA_GROUPID", "tutgroup"),
		OffsetReset:              env.GetEnvAsBoolOrDefault("MYKAFKA_OFFSET_RESET", "true"),
		TLSCertPath:              os.Getenv("MYKAFKA_TLSCERT_PATH"),
		TLSKeyPath:               os.Getenv("MYKAFKA_TLSKEY_PATH"),
		CaCertPath:               os.Getenv("MYKAFKA_CACRT_PATH"),
		MessageCount:             env.GetEnvAsIntOrDefault("MYKAFKA_MESSAGE_NUM", "3"),
		SecurityProtocol:         "SSL",
		EnableSSLCerterification: false,
	}

	// connect as producer
	KafkaProducerExample(cfg)

	// connect as consumer
	KafkaConsumerExample(cfg)
}