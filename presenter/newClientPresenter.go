package presenter

import (
	"../infra/redconfig"
	"../usecase"
	"github.com/gin-gonic/gin"
)

func InsertData(env *redconfig.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usecase.NewClient(c, env.RC)
	}
	return fn
}
