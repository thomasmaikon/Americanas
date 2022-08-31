package utils

import (
	"context"
	"projeto/Americanas/db"
)

type factoryDB struct {
	redis db.GenericDB

	mongo db.GenericDB
}

func (f *factoryDB) GetRedisDB() db.GenericDB {
	if f.redis == nil {
		f.redis = &db.Redis{Connection: getRedisConnection()}
	}
	return f.redis
}

func (f *factoryDB) GetMongoDB() db.GenericDB {
	if f.mongo == nil {
		connection := getConnectMongoDB()
		collection := connection.Database("americanas").Collection("planets")
		f.mongo = &db.Mongo{Collection: collection, Ctx: context.TODO()}
	}
	return f.mongo
}
