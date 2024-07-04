package pkg

import (
	"github.com/gin-gonic/gin"
)

func DateTimeEndpointGin(datetime func() string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.String(200, datetime())
	}
	return gin.HandlerFunc(fn)
}
