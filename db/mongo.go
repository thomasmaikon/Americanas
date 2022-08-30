package db

import (
	"context"
	"projeto/Americanas/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func (m *Mongo) Add(newPlanet model.PLanet) bool {
	doc := bson.D{{"name", newPlanet.Name}, {"climate", newPlanet.Climate}, {"terrain", newPlanet.Terrain}, {"films", newPlanet.Films}}
	_, err := m.Collection.InsertOne(m.Ctx, doc)

	if err != nil {
		panic(err)
	}
	return true
}

func (m *Mongo) FindAll() []model.PLanet {
	list := []model.PLanet{}

	return list
}

func (m *Mongo) FindByName(name string) model.PLanet {

	var planet model.PLanet

	filter := bson.D{{"name", name}}
	err := m.Collection.FindOne(m.Ctx, filter).Decode(&planet)

	if err != nil {
		panic(err)
	}

	return planet
}

func (m *Mongo) FindById(id int64) model.PLanet {
	return model.PLanet{}
}

func (m *Mongo) RemoveByName(name string) bool {
	_, err := m.Collection.DeleteOne(m.Ctx, bson.M{"name": name})
	if err != nil {
		panic(err)
	}
	return true
}
