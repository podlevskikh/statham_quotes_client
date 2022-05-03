package app

import (
	"context"
	"github.com/podlevskikh/statham_quotes_client/app/api"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type App struct {
	logger *zerolog.Logger
}

func (a *App) Start(ctx context.Context, logger *zerolog.Logger) error {
	a.logger = logger

	restAPI := api.NewRestAPI(a.logger)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return restAPI.Run(ctx, ":"+specs.HTTPPost)
	})
	return g.Wait()
}
