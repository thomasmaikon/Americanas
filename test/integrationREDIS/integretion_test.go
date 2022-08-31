package integrationredis

import (
	"fmt"
	"log"
	"os"
	"projeto/Americanas/model"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis"

	dataBase "projeto/Americanas/db"

	"github.com/ory/dockertest"
)

var db *redis.Client

func TestMain(m *testing.M) {

	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("redis", "3.2", nil)
	resource.Expire(30 * uint(time.Minute))

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = pool.Retry(func() error {
		db = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})
		return db.Ping().Err()
	})

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

// nao e possivel testar com redis 7
func TestAddAndFindData(t *testing.T) {
	var banco dataBase.GenericDB
	banco = &dataBase.Redis{Connection: db}

	planet := model.Planet{Name: "exemplo"}

	banco.Add(planet)
	result := banco.FindByName(planet.Name)

	if reflect.DeepEqual(planet, result) == false {
		t.Fatal("planets are not equals")
	}
}

func TestFindDataThatNotExist(t *testing.T) {
	var banco dataBase.GenericDB

	banco = &dataBase.Redis{Connection: db}

	result := banco.FindByName("NotExist")

	if reflect.DeepEqual(result, model.Planet{}) == false {
		t.Fatal("Planet are not equals")
	}
}
