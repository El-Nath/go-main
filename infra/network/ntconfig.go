package network

import "github.com/gin-gonic/gin"
import "../esconfig"

func InitRoutes(env *esconfig.Env) *gin.Engine {

	r := gin.Default()
	vx := r.Group("api/vx")
	{
		vx.POST("/premeditated_data/:fakedata")
		vx.GET("/search/:index")
	}

	return r
}
