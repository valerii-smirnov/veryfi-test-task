package gorm

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	logLevel logger.LogLevel
}

type Option func(cfg *config)

func WithLogLevel(logLevel logger.LogLevel) Option {
	return func(cfg *config) {
		cfg.logLevel = logLevel
	}
}

func NewMockedGorm(opts ...Option) (*gorm.DB, sqlmock.Sqlmock, error) {
	cfg := config{
		logLevel: logger.Silent,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	mysqlDialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gdb, err := gorm.Open(mysqlDialector, &gorm.Config{
		Logger: logger.Default.LogMode(cfg.logLevel),
	})

	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}
