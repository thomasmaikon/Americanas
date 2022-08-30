package integrationredis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"projeto/Americanas/model"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/ory/dockertest"
)

var db *redis.Client
var ctx = context.Background()

func TestMain(m *testing.M) {

	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("redis", "7", nil)
	resource.Expire(30 * uint(time.Minute))

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = pool.Retry(func() error {
		db = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})
		return db.Ping(ctx).Err()
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

func TestAddData(t *testing.T) {
	test := model.Planet{Name: "exemplo"}
	data, _ := json.Marshal(test)
	db.Set(ctx, "exemplo", data, 2*time.Hour)

	value, _ := db.Get(ctx, "exemplo").Result()
	var newData model.Planet
	json.Unmarshal([]byte(value), &newData)

	if reflect.DeepEqual(newData, test) == false {
		t.Fail()
	}
}

func TestRemoveData(t *testing.T) {
	test := model.Planet{Name: "exemplo"}
	data, _ := json.Marshal(test)
	db.Set(ctx, "exemplo", data, 2*time.Hour)

	value, _ := db.Get(ctx, "exemplo").Result()
	var newData model.Planet
	json.Unmarshal([]byte(value), &newData)

	if reflect.DeepEqual(newData, test) == false {
		t.Fail()
	}
}
