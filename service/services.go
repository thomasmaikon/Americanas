package service

import "projeto/Americanas/utils"

func GetServicePlanet() servicePlanet {
	factoryDB := utils.GetFactoryDB()
	return servicePlanet{Repository: factoryDB.GetMongoDB(), RepositoryRedis: factoryDB.GetRedisDB()}
}
