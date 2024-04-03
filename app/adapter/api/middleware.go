package api

import (
	"github.com/gin-gonic/gin"
)

type Serve func(c *gin.Context) error

func Middleware() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}
