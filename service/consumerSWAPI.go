package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projeto/Americanas/model"
)

const startUrl = "https://swapi.dev/api/planets/"

var respositoryDB = factory.GetRedisDB()

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
