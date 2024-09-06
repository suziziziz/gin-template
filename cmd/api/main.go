package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suziziziz/died"
	"github.com/suziziziz/gin-template/configs"
	"github.com/suziziziz/gin-template/pkg/di"
	"github.com/suziziziz/gin-template/pkg/middlewares"
	"github.com/suziziziz/gin-template/pkg/utils"
)

func main() {
	configs.Env()

	events := utils.NewEventEmitter()

	db := configs.DB(os.Getenv("DB_DSN"))

	died := died.New()

	r := gin.Default()
	r.Use(middlewares.GinContextPortal(died))
	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.CORS("*"))

	di.DI(r.Group("v1"), db, events, died)

	r.Run()
}
