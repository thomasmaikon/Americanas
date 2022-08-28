package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projeto/Americanas/model"
)

const startUrl = "https://swapi.dev/api/planets/"

func getData(url string) model.Swapi {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	var data model.Swapi
	json.Unmarshal(body, &data)
	return data
}

func SyncSwapi() {
	for {

		initialURL := startUrl

		swapi := getData(initialURL)
		if swapi.Next == "" {
			break
		}

		initialURL = swapi.Next
	}

}
