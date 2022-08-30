package service

import (
	"context"
	"projeto/Americanas/db"
	"projeto/Americanas/utils"
)

func GetRedisDB() db.GenericDB {
	return &db.Redis{Conexao: utils.RedisConnection(), Ctx: context.Background()}
}
