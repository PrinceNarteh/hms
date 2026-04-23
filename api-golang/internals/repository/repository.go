// Package repository
package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type Reposotory struct {
	User UserRepository
}

func InitRespository(db *mongo.Database) *Reposotory {
	return &Reposotory{
		User: &userRepository{coll: db.Collection("users")},
	}
}
