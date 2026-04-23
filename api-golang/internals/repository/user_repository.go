package repository

import (
	"context"
	"errors"
	"time"

	"hms-backend/internals/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ UserRepository = (*userRepository)(nil)

type UserRepository interface {
	Create(context.Context, *models.User) error
	FindAll(context.Context) (*[]models.User, error)
	FindOneByID(context.Context, string) (*models.User, error)
	FindOneByEmail(context.Context, string) (*models.User, error)
}

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) Create(ctx context.Context, data *models.User) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	_, err := r.coll.InsertOne(ctx, data)
	return err
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

func (r *userRepository) findOneBy(ctx context.Context, filter bson.M) (*models.User, error) {
	user := new(models.User)
	if err := r.coll.FindOne(ctx, filter).Decode(user); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound
		}
	}
	return user, nil
}

func (r *userRepository) FindOneByID(ctx context.Context, id string) (*models.User, error) {
	return r.findOneBy(ctx, bson.M{"_id": id})
}

func (r *userRepository) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.findOneBy(ctx, bson.M{"email": email})
}
