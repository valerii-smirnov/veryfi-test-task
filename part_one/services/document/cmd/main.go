package main

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/config"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/repositories"
	grpctransport "github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/transport/grpc"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/usecases"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/lib/event"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/lib/fswatcher"

	"go.uber.org/fx"
)

func main() {
	buildFXApp().Run()
}

func buildFXApp() *fx.App {
	return fx.New(
		// Register dependencies constructors.
		fx.Provide(newApplicationContext),
		fx.Provide(config.New),
		fx.Provide(newGormConn),
		fx.Provide(newFileWatcher),
		fx.Provide(newJetStreamClient),
		fx.Provide(fx.Annotate(event.NewProducer, fx.As(new(usecases.EventProducer)))),
		fx.Provide(fx.Annotate(fswatcher.NewFSWatcher, fx.As(new(usecases.FileSystemWatcher)))),
		fx.Provide(fx.Annotate(newVeryfiClient, fx.As(new(usecases.VeryfiProcessor)))),
		fx.Provide(fx.Annotate(repositories.NewFile, fx.As(new(usecases.DocumentRepository)))),
		fx.Provide(usecases.NewProcessor),
		fx.Provide(fx.Annotate(usecases.NewDocument, fx.As(new(grpctransport.DocumentUseCase)))),
		fx.Provide(grpctransport.NewDocument),

		// Register decorators.
		fx.Decorate(migrateDB),
		fx.Decorate(documentsStreamInitiator),

		// Register invokers.
		fx.Invoke(func(ctx context.Context, processor *usecases.Processor) error {
			return processor.Run(ctx)
		}),
		fx.Invoke(runGRPCServer),
	)
}
