package usecase

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/domain/repository"
	"github.com/yagrush/go-sample-restapi/app/schemas"
)

type SampleCalcAddInt64Usecase struct{}

func (u SampleCalcAddInt64Usecase) Serve(c *gin.Context, r repository.SampleRepository, params schemas.SampleCalcAddInt64Params) {
	aStr := c.Request.FormValue("a")
	a, err := strconv.ParseInt(aStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": fmt.Errorf("invalid request parameter a=%s", aStr),
		})
		return
	}

	bStr := c.Request.FormValue("b")
	b, err := strconv.ParseInt(bStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": fmt.Errorf("invalid request parameter b=%s", bStr),
		})
		return
	}

	ret, err := r.CalcAddInt64(c, a, b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": fmt.Errorf("error: CalcAddInt64"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": ret,
	})
	return
}
