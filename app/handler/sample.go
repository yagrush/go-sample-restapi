package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/domain/repository"
	"github.com/yagrush/go-sample-restapi/app/schemas"
	"github.com/yagrush/go-sample-restapi/app/usecase"
)

type Handler struct {
	schemas.ServerInterface
	SampleRepository repository.SampleRepository
}

func (h Handler) SampleFuga(c *gin.Context) {
	new(usecase.SampleFugaUsecase).Serve(c, h.SampleRepository)
}

func (h Handler) SampleCalcAddInt64(c *gin.Context, params schemas.SampleCalcAddInt64Params) {
	new(usecase.SampleCalcAddInt64Usecase).Serve(c, h.SampleRepository, params)
}
