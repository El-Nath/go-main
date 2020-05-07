package network

import (
	"../../presenter"
	"../config"
	"github.com/gin-gonic/gin"
)

func InitRoutes(env *config.Env) *gin.Engine {

	r := gin.Default()
	vx := r.Group("api/vx")
	{
		vx.GET("/user", presenter.ViewUser(env))
		vx.POST("/user/add", presenter.AddUser(env))
	}

	return r
}
