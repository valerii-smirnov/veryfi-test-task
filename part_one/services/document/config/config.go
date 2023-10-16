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
	FolderPath            string `env:"FOLDER_PATH,required"`
	WatcherWorkerPoolSize int    `env:"WATCHER_WORKER_POOL_SIZE,required"`
	GRPCServerPort        uint   `env:"GRPC_SERVER_PORT,required"`
	JetStreamHost         string `env:"JET_STREAM_HOST,required"`
	JetStreamPort         uint   `env:"JET_STREAM_PORT,required"`
	PostgresDSN           string `env:"POSTGRES_DSN,required"`

	StreamConfig StreamConfig
	VerifyConfig VeryfiConfig
}
