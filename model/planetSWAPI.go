package model

type Swapi struct {
	Next    string   `json: "next"`
	Results []PLanet `json: "results"`
}
