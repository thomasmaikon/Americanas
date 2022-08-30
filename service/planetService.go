package service

import "projeto/Americanas/model"

var repository = FactoryDB{}.GetMongoDB()
var repositoryRedis = FactoryDB{}.GetRedisDB()

type ServicePlanet struct {
}

func (s ServicePlanet) CreateNewPlanet(newPLanet model.Planet) bool {
	newPLanet.Films = repositoryRedis.FindByName(newPLanet.Name).Films
	return repository.Add(newPLanet)
}

func (s ServicePlanet) Find(name string, id string) []model.Planet {
	result := []model.Planet{}
	if name != "" {
		return append(result, s.findByName(name))
	}
	if id != "" {
		return append(result, s.findById(id))
	}
	return s.findAll()
}

func (s ServicePlanet) findByName(name string) model.Planet {
	return repository.FindByName(name)
}

func (s ServicePlanet) findAll() []model.Planet {
	return repository.FindAll()
}

func (s ServicePlanet) findById(id string) model.Planet {
	return repository.FindById(id)
}

func (s ServicePlanet) RemoveByName(name string) bool {
	return repository.RemoveByName(name)
}
