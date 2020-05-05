package repository

import (
	"encoding/json"

	"../domain/model"
	"github.com/go-redis/redis"
)

func GetData(k string, cR *redis.Client) ([]model.User, error) {

	rP := []model.User{}
	rD, err := cR.Get(k).Result()
	if err != nil {
		return rP, err
	}
	json.Unmarshal([]byte(rD), &rP)
	return rP, nil
}
