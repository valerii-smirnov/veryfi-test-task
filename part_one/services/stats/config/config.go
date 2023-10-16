package config

import "github.com/caarlos0/env/v9"

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

type Config struct {
	RestServerPort                uint   `env:"REST_SERVER_PORT,required"`
	JetStreamHost                 string `env:"JET_STREAM_HOST,required"`
	JetStreamPort                 uint   `env:"JET_STREAM_PORT,required"`
	DocumentServiceHost           string `env:"DOCUMENT_SERVICE_HOST,required"`
	DocumentServicePort           uint   `env:"DOCUMENT_SERVICE_PORT,required"`
	EventProcessorWorkersPoolSize int    `env:"EVENT_PROCESSOR_WORKERS_POOL_SIZE,required"`
	GoogleMapsAPIKey              string `env:"GOOGLE_MAPS_API_KEY,required"`
	PostgresDSN                   string `env:"POSTGRES_DSN,required"`

	StreamConfig StreamConfig
}
