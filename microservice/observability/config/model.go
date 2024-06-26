package config

type Configuration struct {
	App   AppConfig
	Kafka KafkaConfig
	OTLP  OTLPConfig
}

type AppConfig struct {
	Name     string `env:"APP_NAME"`
	Port     string `env:"APP_PORT"`
	Env      string `env:"ENVIRONMENT"`
	LogLevel string `env:"LOG_LEVEL"`
}

type KafkaConfig struct {
	Brokers []string `env:"KAFKA_BROKERS" envSeparator:","`
}

type OTLPConfig struct {
	OTLPEndpoint string `env:"OTLP_ENDPOINT"`
}
