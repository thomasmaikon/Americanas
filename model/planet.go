package model

type PLanet struct {
	Name    string   `json: "name"`
	Climate string   `json: "climate"`
	Terrain string   `json: "terrain"`
	Films   []string `json: "films[]"`
}
