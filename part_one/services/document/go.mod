module github.com/valerii-smirnov/veryfi-test-task/part_one/document

go 1.21.2

require (
	github.com/caarlos0/env/v9 v9.0.0
	github.com/fsnotify/fsnotify v1.6.0
	github.com/nats-io/nats.go v1.31.0
	github.com/sirupsen/logrus v1.9.3
	github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/grpc v0.0.1
	github.com/veryfi/veryfi-go v1.2.2
	go.uber.org/fx v1.20.0
	google.golang.org/grpc v1.58.3
	gorm.io/driver/postgres v1.5.3
	gorm.io/gorm v1.25.5
)

require (
	github.com/creasty/defaults v1.5.1 // indirect
	github.com/go-resty/resty/v2 v2.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/valerii-smirnov/veryfi-test-task/part_one/pkg/grpc v0.0.1 => ../../pkg/grpc
