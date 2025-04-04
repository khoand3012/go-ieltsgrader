package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/go-ieltsgrader/bootstrap"
	"github.com/khoand3012/go-ieltsgrader/db"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db db.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewLoginRouter(env, timeout, db, publicRouter)
}
