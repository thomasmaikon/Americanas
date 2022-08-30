package db

import (
	"context"
	"encoding/json"
	"projeto/Americanas/model"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	Connection *redis.Client
	Ctx        context.Context
}

func (r *Redis) Add(newPlanet model.PLanet) bool {

	data, _ := json.Marshal(newPlanet)
	err := r.Connection.Set(r.Ctx, newPlanet.Name, data, 0).Err()

	if err != nil {
		panic(err)
	}

	return true
}

func (r *Redis) FindAll() []model.PLanet {
	return []model.PLanet{}
}

func (r *Redis) FindByName(name string) model.PLanet {
	value, err := r.Connection.Get(r.Ctx, name).Result()
	if err != nil {
		panic(err)
	}
	var planet model.PLanet
	json.Unmarshal([]byte(value), &planet)
	return planet
}

func (r *Redis) FindById(id int64) model.PLanet {
	return model.PLanet{}
}

func (r *Redis) RemoveByName(name string) bool {
	return false
}
