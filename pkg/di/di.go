package di

import (
	"github.com/gin-gonic/gin"
	"github.com/suziziziz/died"
	"github.com/suziziziz/gin-template/pkg/providers"
	"github.com/suziziziz/gin-template/pkg/utils"
	"gorm.io/gorm"
)

func DI(
	g *gin.RouterGroup,
	db *gorm.DB,
	events utils.EventEmitter,
	d died.Died,
) {
	d.Inject(providers.NewGinContextPortalProvider)

	d.Inject(func() *gin.RouterGroup { return g })
	d.Inject(func() *gorm.DB { return db })
	d.Inject(func() utils.EventEmitter { return events })

	d.Invoke()
}
