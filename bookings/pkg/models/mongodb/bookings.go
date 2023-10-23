package mongodb

import (
	"context"
	"errors"
	"github.com/sayedppqq/go-microservices-project/bookings/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingModel struct {
	C *mongo.Collection
}

func (m *BookingModel) GetAllBookings() ([]models.Booking, error) {
	ctx := context.TODO()
	b := []models.Booking{}

	bookingCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = bookingCursor.All(ctx, &b)
	if err != nil {
		return nil, err
	}

	return b, err
}

func (m *BookingModel) GetBookingsByID(id string) (*models.Booking, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var booking = models.Booking{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &booking, nil
}

func (m *BookingModel) InsertNewBookings(booking models.Booking) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), booking)
}

func (m *BookingModel) DeleteBookingsByID(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
