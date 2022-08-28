package model

type PlanetSWAPI struct {
	Name    string   `json: "name"`
	Climate string   `json: "climate"`
	Terrain string   `json: "terrain"`
	Films   []string `json: "films[]"`
}

type Swapi struct {
	Next    string        `json: "next"`
	Results []PlanetSWAPI `json: "results"`
}
