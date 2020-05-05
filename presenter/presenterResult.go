package presenter

import (
	"../infra/esconfig"
	"../usecase"
	"github.com/gin-gonic/gin"
)

func ResultData(env *esconfig.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usecase.ResultSearch(c, env.ES)
	}
	return fn
}
