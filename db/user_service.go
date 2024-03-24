package db

import (
	"context"
	"errors"

	"com.arturs.shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson"
)

const UsersCollection = "users"
const EmailField = "email"

var (
	ErrUserExists = errors.New("User already exists")
)

type UserService struct {
	Database *Database
}

func (service *UserService) createUser(user *models.User) error {
	exists, err := service.ContainsUser(user.Email)
	if exists == false && err != nil {
		_, err := service.Database.Collection(UsersCollection).InsertOne(context.Background(), user)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) ContainsUser(email string) (bool, error) {
	count, err := service.Database.Collection(UsersCollection).CountDocuments(context.Background(), bson.M{EmailField: email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
