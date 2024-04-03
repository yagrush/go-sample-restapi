package usecase

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yagrush/go-sample-restapi/app/domain/repository"
)

type SampleCalcAddInt64Usecase struct {
	R repository.SampleRepository
}

func (u *SampleCalcAddInt64Usecase) Serve(c *gin.Context) (int, gin.H) {
	aStr := c.Request.FormValue("a")
	a, err := strconv.ParseInt(aStr, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"errorMessage": fmt.Errorf("invalid request parameter a=%s", aStr),
		}
	}

	bStr := c.Request.FormValue("b")
	b, err := strconv.ParseInt(bStr, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"errorMessage": fmt.Errorf("invalid request parameter b=%s", bStr),
		}
	}

	ret, err := u.R.CalcAddInt64(c, a, b)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"errorMessage": fmt.Errorf("error: CalcAddInt64"),
		}
	}

	// fmt.Fprint(c.Writer, ret)
	return http.StatusOK, gin.H{
		"message": ret,
	}
}
