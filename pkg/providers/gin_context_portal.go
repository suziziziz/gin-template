package providers

import "github.com/gin-gonic/gin"

type GinContextPortal struct {
	ginC *gin.Context
}

func NewGinContextPortalProvider() *GinContextPortal {
	return &GinContextPortal{}
}

func (g *GinContextPortal) SetContext(ginC *gin.Context) {
	g.ginC = ginC
}

func (g *GinContextPortal) Context() *gin.Context {
	return g.ginC
}
