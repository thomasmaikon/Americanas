package db

import (
	"projeto/Americanas/model"

	"github.com/go-redis/redis/v9"
)

var generic map[string]*model.PLanet = make(map[string]*model.PLanet)

type Redis struct {
	Conexao *redis.Client
}

func (r *Redis) Create(newPlanet model.PLanet) bool {

	if generic[newPlanet.Name] == nil {
		generic[newPlanet.Name] = &newPlanet
		return true
	}

	return false
}

func (r *Redis) FindAll() []model.PLanet {
	list := []model.PLanet{}
	for _, planet := range generic {
		list = append(list, *planet)
	}
	return list
}

func (r *Redis) FindByName(name string) model.PLanet {
	return *generic[name]
}

func (r *Redis) FindById(id int64) model.PLanet {
	return model.PLanet{}
}

func (r *Redis) RemoveAll() {
	for k := range generic {
		delete(generic, k)
	}
}

func (r *Redis) RemoveByName(name string) bool {
	alredyExist := generic[name]

	if alredyExist != nil {
		alredyExist = nil
		return true
	}
	return false
}
