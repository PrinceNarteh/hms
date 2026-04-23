package repository

import (
	"context"

	"hms-backend/internals/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ UserRepository = (*userRepository)(nil)

type UserRepository interface {
	FindByID(context.Context, string) (*models.User, error)
}

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) FindAll(ctx context.Context) (*[]models.User, error) {
	cursor, err := r.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cursor.Close(ctx); err != nil {
		return nil, err
	}

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
