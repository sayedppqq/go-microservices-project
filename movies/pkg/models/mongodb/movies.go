package mongodb

import (
	"context"
	"errors"
	"github.com/sayedppqq/go-microservices-project/movies/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieModel struct {
	C *mongo.Collection
}

func (m *MovieModel) GetAllMovies() ([]models.Movie, error) {
	movies := []models.Movie{}
	cur, err := m.C.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &movies)
	if err != nil {
		return nil, err
	}

	return movies, err
}

func (m *MovieModel) GetMovieByID(id string) (*models.Movie, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	movie := models.Movie{}
	filter := bson.D{primitive.E{Key: "_id", Value: p}}

	err = m.C.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document find by this id")
		}
		return nil, err
	}
	return &movie, nil
}

func (m *MovieModel) InsertNewMovie(movie models.Movie) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), movie)
}

func (m *MovieModel) DeleteMovieByID(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
