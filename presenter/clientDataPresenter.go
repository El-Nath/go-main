package presenter

import (
	"../infra/redconfig"
	"../usecase"
	"github.com/gin-gonic/gin"
)

func RedisData(env *redconfig.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usecase.GetRedisData(c, env.RC)
	}
	return fn
}
