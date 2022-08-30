package db

import "projeto/Americanas/model"

type GenericDB interface {
	Add(newPlanet model.PLanet) bool
	FindAll() []model.PLanet
	FindByName(name string) model.PLanet
	FindById(id int64) model.PLanet
	RemoveAll()
	RemoveByName(name string) bool
	//RemoveById(id int64)
}
