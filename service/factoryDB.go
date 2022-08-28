package service

import (
	"projeto/Americanas/db"
)

func GetRedisDB() db.GenericDB {
	return &db.Redis{}
}
