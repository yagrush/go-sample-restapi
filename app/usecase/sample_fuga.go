package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/domain/repository"
)

type SampleFugaUsecase struct {
	R repository.SampleRepository
}

func (u *SampleFugaUsecase) Serve(c *gin.Context) {
	m, err := u.R.GetFuga(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": m.GetMessage(),
	})
}
