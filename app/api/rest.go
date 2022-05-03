package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/podlevskikh/statham_quotes_client/app/api/controllers/quotes"
	"github.com/rs/zerolog"
)

type RestAPI struct {
	logger *zerolog.Logger
}

func NewRestAPI(logger *zerolog.Logger) *RestAPI {
	return &RestAPI{logger: logger}
}

func (a *RestAPI) Run(ctx context.Context, port string) error {
	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.Writer.Header().Set("Content-Type", "application/json")
	}))

	a.quoteHandlers(r)

	return r.Run(port)
}

func (a *RestAPI) quoteHandlers(r *gin.Engine) {
	getCurrents := quotes.NewGetQuote(a.logger)
	r.GET("/api/quotes/random", getCurrents.HTTPHandler)
}