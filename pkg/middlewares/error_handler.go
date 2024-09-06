package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		if errors.Is(err, gin.Error{}) {
			c.String(int(err.Type), err.Err.Error())
			return
		}
	}

	if len(c.Errors) > 0 {
		c.String(http.StatusInternalServerError, c.Errors[0].Error())
	}
}
