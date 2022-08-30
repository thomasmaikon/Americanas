package db

import "projeto/Americanas/model"

type GenericDB interface {
	Add(newPlanet model.Planet) bool
	FindAll() []model.Planet
	FindByName(name string) model.Planet
	FindById(id string) model.Planet
	RemoveByName(name string) bool
}
