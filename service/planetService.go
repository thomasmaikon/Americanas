package service

import "projeto/Americanas/model"

var repository = FactoryDB{}.GetMongoDB()

type ServicePlanet struct {
}

func (s ServicePlanet) CreateNewPlanet(newPLanet model.PLanet) bool {
	return repository.Add(newPLanet)
}

func (s ServicePlanet) Find(name string, id int64) []model.PLanet {
	result := []model.PLanet{}
	if name != "" {
		return append(result, s.findByName(name))
	}
	if id != 0 {
		return append(result, s.findById(id))
	}
	return s.findAll()
}

func (s ServicePlanet) findByName(name string) model.PLanet {
	return repository.FindByName(name)
}

func (s ServicePlanet) findAll() []model.PLanet {
	return repository.FindAll()
}

func (s ServicePlanet) findById(id int64) model.PLanet {
	return repository.FindById(id)
}

func (s ServicePlanet) RemoveByName(name string) bool {
	return repository.RemoveByName(name)
}
