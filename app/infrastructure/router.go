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
	"github.com/yagrush/go-sample-restapi/app/usecase"
)

func NewEngine(c config.Config) (*gin.Engine, error) {
	f, _ := os.Create(c.ServerLogFile)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	engine := gin.Default()

	samplePersistence := persistence.NewSamplePersistence()
	schemas.RegisterHandlersWithOptions(engine, &handler.Handler{
		SampleFugaUsecase: usecase.SampleFugaUsecase{
			R: samplePersistence,
		},
		SampleCalcAddInt64Usecase: usecase.SampleCalcAddInt64Usecase{
			R: samplePersistence,
		},
	}, schemas.GinServerOptions{
		Middlewares: []schemas.MiddlewareFunc{
			api.Middleware(),
		},
	})

	return engine, nil
}
