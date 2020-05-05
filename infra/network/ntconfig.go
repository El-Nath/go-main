package network

import (
	"../../presenter"
	"../redconfig"
	"github.com/gin-gonic/gin"
)

func InitRoutes(env *redconfig.Env) *gin.Engine {

	r := gin.Default()
	vx := r.Group("api/vx")
	{
		vx.GET("/list_data")
		vx.POST("/data/:client/:ttl", presenter.InsertData(env))
		vx.GET("/data/:key", presenter.RedisData(env))
	}

	return r
}
