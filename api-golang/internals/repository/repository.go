// Package repository
package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrNotFound = errors.New("user not found")

type Repository struct {
	User UserRepository
}

func InitRespository(db *mongo.Database) *Repository {
	return &Repository{
		User: &userRepository{coll: db.Collection("users")},
	}
}
