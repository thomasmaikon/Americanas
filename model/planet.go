package model

type PLanet struct {
	Id      int32
	Name    string   `json: "name" bson:"name"`
	Climate string   `json: "climate" bson: "climate"`
	Terrain string   `json: "terrain" bson: "terrain"`
	Films   []string `json: "films[]" bson: "films, omitempty"`
}
