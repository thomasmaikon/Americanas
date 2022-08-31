package service

import (
	"projeto/Americanas/db"
	"projeto/Americanas/model"
)

type servicePlanet struct {
	Repository      db.GenericDB
	RepositoryRedis db.GenericDB
}

func (s servicePlanet) CreateNewPlanet(newPLanet model.Planet) bool {
	newPLanet.Films = s.RepositoryRedis.FindByName(newPLanet.Name).Films
	return s.Repository.Add(newPLanet)
}

func (s servicePlanet) Find(name string, id string) []model.Planet {
	result := []model.Planet{}
	if name != "" {
		return append(result, s.findByName(name))
	}
	if id != "" {
		return append(result, s.findById(id))
	}
	return s.findAll()
}

func (s servicePlanet) findByName(name string) model.Planet {
	return s.Repository.FindByName(name)
}

func (s servicePlanet) findAll() []model.Planet {
	return s.Repository.FindAll()
}

func (s servicePlanet) findById(id string) model.Planet {
	return s.Repository.FindById(id)
}

func (s servicePlanet) RemoveByName(name string) bool {
	return s.Repository.RemoveByName(name)
}
