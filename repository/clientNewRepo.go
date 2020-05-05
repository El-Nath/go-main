package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"../domain/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func NewClientData(c *redis.Client, d []model.User, cN string, tL int) error {

	dB, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.Set(cN, dB, time.Duration(tL)*time.Minute).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func BindJsonRequest(c *gin.Context) ([]model.User, error) {
	var d []model.User
	if err := c.BindJSON(&d); err != nil {
		return d, err
	}
	return d, nil
}
