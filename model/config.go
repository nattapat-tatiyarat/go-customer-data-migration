package model

import "go.mongodb.org/mongo-driver/mongo"

type Argument struct {
	Uri        string
	Db         string
	Collection string
	Path       string
}

type MigrationConfig struct {
	Argument   Argument
	Collection *mongo.Collection
}
