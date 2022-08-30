package db

import (
	"context"
	"fmt"
	"log"
	"projeto/Americanas/model"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func (m *Mongo) Add(newPlanet model.Planet) bool {

	result := m.FindByName(newPlanet.Name)
	if reflect.DeepEqual(result, model.Planet{}) == false {
		return false
	}

	doc := bson.D{{"name", newPlanet.Name}, {"climate", newPlanet.Climate}, {"terrain", newPlanet.Terrain}, {"films", newPlanet.Films}}
	_, err := m.Collection.InsertOne(m.Ctx, doc)

	if err != nil {
		panic(err)
	}

	return true
}

func (m *Mongo) FindAll() []model.Planet {
	listPlanets := []model.Planet{}

	cursor, err := m.Collection.Find(m.Ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var result bson.D
		var planet model.Planet
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		bytes, _ := bson.Marshal(result)
		bson.Unmarshal(bytes, &planet)
		listPlanets = append(listPlanets, planet)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return listPlanets
}

func (m *Mongo) parser(data bson.M) model.Planet {
	name := data["name"]
	terrain := data["terrain"]
	climate := data["climate"]
	films := data["films"]

	fmt.Println(name, terrain, climate, films)
	return model.Planet{Name: name.(string), Terrain: terrain.(string), Climate: climate.(string), Films: films.([]string)}
}

func (m *Mongo) FindByName(name string) model.Planet {

	var planet model.Planet

	filter := bson.D{{"name", name}}
	err := m.Collection.FindOne(m.Ctx, filter).Decode(&planet)

	if err != nil {
		return model.Planet{}
	}

	return planet
}

func (m *Mongo) FindById(id string) model.Planet {
	var planet model.Planet
	objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", objId}}
	err := m.Collection.FindOne(m.Ctx, filter).Decode(&planet)

	if err != nil {
		return model.Planet{}
	}

	return planet
}

func (m *Mongo) RemoveByName(name string) bool {
	removed, err := m.Collection.DeleteOne(m.Ctx, bson.M{"name": name})
	if err != nil || removed.DeletedCount == 0 {
		return false
	}
	return true
}
