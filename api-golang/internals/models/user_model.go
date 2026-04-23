package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&u.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
