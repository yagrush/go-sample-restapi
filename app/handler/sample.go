package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/schemas"
	"github.com/yagrush/go-sample-restapi/app/usecase"
)

type Handler struct {
	schemas.ServerInterface
	SampleFugaUsecase         usecase.SampleFugaUsecase
	SampleCalcAddInt64Usecase usecase.SampleCalcAddInt64Usecase
}

func (h Handler) SampleFuga(c *gin.Context) {
	h.SampleFugaUsecase.Serve(c)
}

func (h Handler) SampleCalcAddInt64(c *gin.Context, params schemas.SampleCalcAddInt64Params) {
	h.SampleCalcAddInt64Usecase.Serve(c, params)
}
