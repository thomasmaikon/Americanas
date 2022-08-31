package db

import (
	"encoding/json"
	"projeto/Americanas/model"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	Connection *redis.Client
}

func (r *Redis) Add(newPlanet model.Planet) bool {

	data, _ := json.Marshal(newPlanet)
	err := r.Connection.Set(newPlanet.Name, data, 24*time.Hour).Err()

	if err != nil {
		panic(err)
	}

	return true
}

func (r *Redis) FindAll() []model.Planet {
	return []model.Planet{}
}

func (r *Redis) FindByName(name string) model.Planet {
	value, err := r.Connection.Get(name).Result()
	if err != nil {
		return model.Planet{}
	}
	var planet model.Planet
	json.Unmarshal([]byte(value), &planet)
	return planet
}

func (r *Redis) FindById(id string) model.Planet {
	return model.Planet{}
}

func (r *Redis) RemoveByName(name string) bool {
	return false
}
