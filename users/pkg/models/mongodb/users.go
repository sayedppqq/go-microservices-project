package mongodb

import (
	"context"
	"errors"
	"github.com/sayedppqq/go-microservices-project/users/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	C *mongo.Collection
}

func (m *UserModel) GetAllUsers() ([]models.User, error) {
	// Define variables
	ctx := context.TODO()
	uu := []models.User{}

	// Find all users
	userCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = userCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}

	return uu, err
}

func (m *UserModel) GetUserByID(id string) (*models.User, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find user by id
	var user = models.User{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&user)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) InsertNewUser(user models.User) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

func (m *UserModel) DeleteUserById(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
