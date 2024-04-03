package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/adapter/api"
	"github.com/yagrush/go-sample-restapi/app/handler"
	"github.com/yagrush/go-sample-restapi/app/infrastructure/persistence"
	"github.com/yagrush/go-sample-restapi/app/usecase"
)

func NewEngine() (*gin.Engine, error) {
	engine := gin.Default()
	engine.Use(api.Middleware())

	engine.GET("/api/v1/sample/fuga", func(c *gin.Context) {
		c.JSON(handler.SampleFugaHandler{
			U: usecase.SampleFugaUsecase{
				R: persistence.NewSamplePersistence(),
			},
		}.Serve(c))
	})

	engine.GET("/api/v1/sample/calcAddInt64", func(c *gin.Context) {
		c.JSON(handler.SampleCalcAddInt64Handler{
			U: usecase.SampleCalcAddInt64Usecase{
				R: persistence.NewSamplePersistence(),
			},
		}.Serve(c))
	})

	return engine, nil
}
