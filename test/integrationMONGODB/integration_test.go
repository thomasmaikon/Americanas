package integrationmongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"projeto/Americanas/db"
	"projeto/Americanas/model"
	"reflect"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client
var ctx context.Context

func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "5.0",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=americanas",
			"MONGO_INITDB_ROOT_PASSWORD=americanas",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = pool.Retry(func() error {
		var err error
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://americanas:americanas@localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(ctx, nil)
	})

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	if err = dbClient.Disconnect(ctx); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestAddPlanet(t *testing.T) {
	collection := dbClient.Database("americanas").Collection("planets")
	var banco db.GenericDB
	banco = &db.Mongo{Collection: collection, Ctx: ctx}

	newPlanet := model.Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert", Films: []string{}}

	result := banco.Add(newPlanet)
	if !result {
		t.Fail()
	}
}

func TestFindPlanet(t *testing.T) {
	collection := dbClient.Database("americanas").Collection("planets")
	var banco db.GenericDB
	banco = &db.Mongo{Collection: collection, Ctx: ctx}

	newPlanet := model.Planet{Name: "Test", Climate: "arid", Terrain: "desert"}

	banco.Add(newPlanet)
	planet := banco.FindByName("Tattoine")

	if reflect.DeepEqual(newPlanet, planet) == false {
		t.Fatalf("Os objetos eram para ser iguais")
	}
}

func TestRemovePlanet(t *testing.T) {
	collection := dbClient.Database("americanas").Collection("planets")
	var banco db.GenericDB
	banco = &db.Mongo{Collection: collection, Ctx: context.TODO()}

	newPlanet := model.Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert"}

	banco.Add(newPlanet)
	result := banco.RemoveByName("Tattoine")

	if !result {
		t.Fatalf("O planeta nao era para existir")
	}
}
