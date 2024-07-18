package repositories

import (
	db2 "Expense_Manager/commons/db"
	"Expense_Manager/pkg/user_manager_service/models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

const (
	collectionUsers = "users"
)

type UserManagerRepo interface {
	CreateNewUser(user models.User) error
	IsUsernameExist(username string) (bool, error)
	IsEmailExist(email string) (bool, error)
	UpdateUserPassword(username string, password []byte) error
}

type UserManagerRepoImpl struct {
	db *mongo.Database
}

func NewUserManagerRepository() (UserManagerRepo, error) {
	db, err := db2.NewMongoDatabase()
	if err != nil {
		return nil, err
	}

	return &UserManagerRepoImpl{db: db}, nil
}

func (u UserManagerRepoImpl) UpdateUserPassword(username string, password []byte) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"password": password}}

	_, err := u.db.Collection(collectionUsers).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (u UserManagerRepoImpl) IsUsernameExist(username string) (bool, error) {
	filter := bson.D{{"username", username}}

	var res models.User
	err := u.db.Collection(collectionUsers).FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, errors.Wrap(err, "could not perform find one")
	}

	return true, nil
}

func (u UserManagerRepoImpl) IsEmailExist(email string) (bool, error) {
	filter := bson.D{{"email", email}}

	var res models.User
	err := u.db.Collection(collectionUsers).FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, errors.Wrap(err, "could not perform find one")
	}

	return true, nil
}

func (u UserManagerRepoImpl) CreateNewUser(user models.User) error {
	one, err := u.db.Collection(collectionUsers).InsertOne(context.Background(), &user)
	if err != nil {
		return errors.Wrap(err, "could not create new user, username="+user.Username)
	}

	slog.Info("created a new user", "id", one.InsertedID)

	return nil
}
