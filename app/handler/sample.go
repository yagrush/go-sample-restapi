package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/usecase"
)

type SampleFugaHandler struct {
	U usecase.SampleFugaUsecase
}

func (h SampleFugaHandler) Serve(c *gin.Context) (int, gin.H) {
	return h.U.Serve(c)
}

type SampleCalcAddInt64Handler struct {
	U usecase.SampleCalcAddInt64Usecase
}

func (h SampleCalcAddInt64Handler) Serve(c *gin.Context) (int, gin.H) {
	return h.U.Serve(c)
}
