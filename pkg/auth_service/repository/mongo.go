package repository

import (
	db2 "Expense_Manager/commons/db"
	models2 "Expense_Manager/pkg/user_manager_service/models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionUsers = "users"
)

type AuthServiceMongoRepo interface {
	GetHashedPasswordOfUser(username string) ([]byte, error)
}

type AuthServiceMongoRepoImpl struct {
	db *mongo.Database
}

func NewAuthServiceMongoRepository() (AuthServiceMongoRepo, error) {
	db, err := db2.NewMongoDatabase()
	if err != nil {
		return nil, err
	}

	return &AuthServiceMongoRepoImpl{db: db}, nil
}

func (a *AuthServiceMongoRepoImpl) GetHashedPasswordOfUser(username string) ([]byte, error) {
	query := bson.D{
		{"username", username},
	}

	var res models2.User
	err := a.db.Collection(collectionUsers).FindOne(context.Background(), query).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "could not perform finding document")
	}

	return res.HashedPassword, nil
}
