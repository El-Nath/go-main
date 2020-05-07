package presenter

import (
	"../infra/config"
	"../usecase"
	"github.com/gin-gonic/gin"
)

func ViewUser(env *config.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usecase.GetAllUser(c, env.DB)
	}
	return fn
}

func AddUser(env *config.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usecase.InsertUser(c, env.DB)
	}
	return fn
}
