package network

import "github.com/gin-gonic/gin"
import "../esconfig"
import "../../presenter"

func InitRoutes(env *esconfig.Env) *gin.Engine {

	r := gin.Default()
	vx := r.Group("api/vx")
	{
		vx.POST("/premeditated_data/:fakedata", presenter.InsertData(env))
		vx.GET("/search/:index", presenter.ResultData(env))
	}

	return r
}
