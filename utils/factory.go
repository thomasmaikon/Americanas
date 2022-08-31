package utils

var factory *factoryDB

func GetFactoryDB() *factoryDB {
	if factory == nil {
		factory = new(factoryDB)
	}
	return factory
}
