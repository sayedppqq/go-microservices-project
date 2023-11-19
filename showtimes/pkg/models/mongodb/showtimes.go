package mongodb

import (
	"context"
	"errors"
	"github.com/sayedppqq/go-microservices-project/showtimes/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ShowTimeModel represent a mgo database session with a showtime data model.
type ShowTimeModel struct {
	C *mongo.Collection
}

func (m *ShowTimeModel) GetAllShowTimes() ([]models.ShowTime, error) {
	st := []models.ShowTime{}
	cur, err := m.C.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &st)
	if err != nil {
		return nil, err
	}

	return st, err
}

func (m *ShowTimeModel) GetShowTimeByID(id string) (*models.ShowTime, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	st := models.ShowTime{}
	filter := bson.D{primitive.E{Key: "_id", Value: p}}

	err = m.C.FindOne(context.TODO(), filter).Decode(&st)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document find by this id")
		}
		return nil, err
	}
	return &st, nil
}

func (m *ShowTimeModel) GetShowTimeByDate(date string) (*models.ShowTime, error) {
	// Find showtime by date
	var showtime = models.ShowTime{}
	err := m.C.FindOne(context.TODO(), bson.M{"date": date}).Decode(&showtime)
	if err != nil {
		// Checks if the showtime was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &showtime, nil
}

func (m *ShowTimeModel) InsertNewShowTime(showtime models.ShowTime) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), showtime)
}

func (m *ShowTimeModel) DeleteShowTimeByID(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
