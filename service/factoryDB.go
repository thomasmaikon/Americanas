package service

import (
	"context"
	"projeto/Americanas/db"
	"projeto/Americanas/utils"
)

type FactoryDB struct {
}

func (f FactoryDB) GetRedisDB() db.GenericDB {
	return &db.Redis{Connection: utils.RedisConnection(), Ctx: context.Background()}
}

func (f FactoryDB) GetMongoDB() db.GenericDB {
	connection := utils.GetConnectMongoDB()
	collection := connection.Database("americanas").Collection("planets")
	return &db.Mongo{Collection: collection, Ctx: context.TODO()}
}
