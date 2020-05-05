package presenter

import (
	"../infra/esconfig"
	"../usecase"
	"github.com/gin-gonic/gin"
)

func InsertData(env *esconfig.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usecase.NewData(c, env.ES)
	}
	return fn
}
