package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projeto/Americanas/model"
	"projeto/Americanas/utils"
)

const startUrl = "https://swapi.dev/api/planets/"

var respositoryDB = utils.GetFactoryDB().GetRedisDB()

func getData(url string) model.Swapi {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	var data model.Swapi
	json.Unmarshal(body, &data)
	return data
}

func addPlanets(swapi model.Swapi) {
	for _, planet := range swapi.Results {
		respositoryDB.Add(planet)
	}
}

// colocar gatilho para executar e atualizar os dados, talvez a api atualize
func SyncSwapi() {
	initialURL := startUrl
	for {

		swapi := getData(initialURL)
		addPlanets(swapi)
		if swapi.Next == "" {
			break
		}
		initialURL = swapi.Next

	}

}
