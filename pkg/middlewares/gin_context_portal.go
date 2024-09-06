package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/suziziziz/died"
	"github.com/suziziziz/gin-template/pkg/providers"
)

func GinContextPortal(died died.Died) gin.HandlerFunc {
	return func(c *gin.Context) {
		g := died.Require(func(*providers.GinContextPortal) {}).(*providers.GinContextPortal)

		g.SetContext(c)

		c.Next()
	}
}
