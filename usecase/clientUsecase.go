package usecase

import (
	"strconv"

	"../repository"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func NewClient(c *gin.Context, cR *redis.Client) {

	cN := c.Params.ByName("client")
	if cN == "" {
		repository.BadResponse(c, "no client")
		return
	}

	tL := c.Params.ByName("ttl")
	if tL == "" {
		tL = "60"
	}
	
	tTL, _ := strconv.Atoi(tL)
	
	d, err := repository.GetData(cN, cR)
	if len(d) != 0 {
		repository.BadResponse(c, "Failed")
		return
	}

	dJ, err := repository.BindJsonRequest(c)
	err = repository.NewClientData(cR, dJ, cN, tTL)
	if err != nil {
		repository.BadResponse(c, err.Error())
		return
	}

	repository.SuccessResponse(c, "success adding data")
}

func GetRedisData(c *gin.Context, cR *redis.Client) {
	cN := c.Params.ByName("key")
	dR, err := repository.GetData(cN, cR)
	if err == redis.Nil {
		repository.BadResponse(c, cN+" not exist")
		return
	} else if err != nil {
		repository.BadResponse(c, err.Error())
		return
	}
	repository.SuccessGettingData(c, dR)
}
