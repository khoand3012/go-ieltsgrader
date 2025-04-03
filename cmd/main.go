package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khoand3012/go-ieltsgrader/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	gin := gin.Default()
	gin.Run(env.ServerAddress)
}
