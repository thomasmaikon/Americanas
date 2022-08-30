package service

import (
	"context"
	"projeto/Americanas/db"
	"projeto/Americanas/utils"
)

type FactoryDB struct {
	redis db.GenericDB

	mongo db.GenericDB
}

func (f *FactoryDB) GetRedisDB() db.GenericDB {
	if f.redis == nil {
		f.redis = &db.Redis{Connection: utils.RedisConnection(), Ctx: context.Background()}
	}
	return f.redis
}

func (f *FactoryDB) GetMongoDB() db.GenericDB {
	if f.mongo == nil {
		connection := utils.GetConnectMongoDB()
		collection := connection.Database("americanas").Collection("planets")
		f.mongo = &db.Mongo{Collection: collection, Ctx: context.TODO()}
	}
	return f.mongo
}
