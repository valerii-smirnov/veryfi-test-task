package main

import (
	"context"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/config"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/repositories"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/transport/rest"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/internal/usecases"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/stats/lib/event"

	"go.uber.org/fx"
)

func main() {
	buildFXApp().Run()
}

func buildFXApp() *fx.App {
	return fx.New(
		// register dependencies constructors
		fx.Provide(newApplicationContext),
		fx.Provide(config.New),
		fx.Provide(newGormConn),
		fx.Provide(newJetStreamClient),
		fx.Provide(newGoogleMapsClient),
		fx.Provide(newDocumentServiceClient),
		fx.Provide(event.NewCreatedChan, event.NewRemovedChan),
		fx.Provide(event.NewConsumer),
		fx.Provide(fx.Annotate(repositories.NewDocument, fx.As(new(usecases.DocumentRepository)))),
		fx.Provide(fx.Annotate(repositories.NewGeolocation, fx.As(new(usecases.GeolocationRepository)))),
		fx.Provide(fx.Annotate(repositories.NewReceipt, fx.As(new(usecases.ReceiptRepository)))),
		fx.Provide(fx.Annotate(repositories.NewStats, fx.As(new(usecases.StatsRepository)))),
		fx.Provide(usecases.NewProcessor),
		fx.Provide(fx.Annotate(usecases.NewStats, fx.As(new(rest.StatsUseCase)))),
		fx.Provide(rest.NewStats),

		// register decorators
		fx.Decorate(migrateDB),

		// register invokers
		fx.Invoke(runConsumers),
		fx.Invoke(func(ctx context.Context, processor *usecases.Processor) {
			processor.Run(ctx)
		}),
		fx.Invoke(runRESTServer),
	)
}
