package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/go-ieltsgrader/bootstrap"
	"github.com/khoand3012/go-ieltsgrader/mongo"
)

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

func LoginHandler(c *gin.Context) {
	lr := LoginResponse{
		AccessToken:  "some-access-token",
		RefreshToken: "my-refresh-token",
	}
	c.JSON(http.StatusOK, lr)
}

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {

	group.POST("/login", LoginHandler)
}
