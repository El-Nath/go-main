package network

import (
	"../config"
	"github.com/gin-gonic/gin"
)

func InitRoutes(env *config.Env) *gin.Engine {

	r := gin.Default()
	vx := r.Group("api/vx")
	{
	}

	return r
}
