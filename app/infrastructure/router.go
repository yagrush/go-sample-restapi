package infrastructure

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/adapter/api"
	"github.com/yagrush/go-sample-restapi/app/handler"
	"github.com/yagrush/go-sample-restapi/app/infrastructure/config"
	"github.com/yagrush/go-sample-restapi/app/infrastructure/persistence"
	"github.com/yagrush/go-sample-restapi/app/schemas"
)

func NewEngine(c config.Config) (*gin.Engine, error) {
	f, _ := os.Create(c.ServerLogFile)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	engine := gin.Default()

	schemas.RegisterHandlersWithOptions(engine, &handler.Handler{
		R: persistence.SamplePersistence{},
	}, schemas.GinServerOptions{
		Middlewares: []schemas.MiddlewareFunc{
			api.Middleware(),
		},
	})

	return engine, nil
}
