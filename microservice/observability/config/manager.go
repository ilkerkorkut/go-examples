package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type ConfigurationManager interface {
	GetAppConfig() AppConfig
	GetKafkaConfig() KafkaConfig
	GetOTLPConfig() OTLPConfig
}

type configurationManager struct {
	app   AppConfig
	kafka KafkaConfig
	otlp  OTLPConfig
}

func NewConfigurationManager() ConfigurationManager {
	var configuration Configuration

	if err := env.Parse(&configuration); err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	c := &configurationManager{
		app:   configuration.App,
		kafka: configuration.Kafka,
		otlp:  configuration.OTLP,
	}

	return c
}

func (c *configurationManager) GetAppConfig() AppConfig {
	return c.app
}

func (c *configurationManager) GetKafkaConfig() KafkaConfig {
	return c.kafka
}

func (c *configurationManager) GetOTLPConfig() OTLPConfig {
	return c.otlp
}
