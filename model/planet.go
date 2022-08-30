package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `json: "name" bson:"name"`
	Climate string             `json: "climate" bson: "climate"`
	Terrain string             `json: "terrain" bson: "terrain"`
	Films   []string           `json: "films[]" bson: "films, omitempty"`
}
