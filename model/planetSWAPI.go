package model

type Swapi struct {
	Next    string   `json: "next"`
	Results []Planet `json: "results"`
}
