package db

import (
	"context"
	"peter/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(email, password string) error {
	m := ConnectToDB()
	coll := m.Database("db").Collection("user")

	user := entities.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: password,
	}

	_, err := coll.InsertOne(context.Background(), user)

	defer DisconnectFromDB(m)

	return err
}

func ReferencePassword(email, website string, password entities.Password) error {
	m := ConnectToDB()
	coll := m.Database("db").Collection("user")

	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "encryptedPasswords." + website, Value: password}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)

	defer DisconnectFromDB(m)

	return err
}

func GetUsers(username string) ([]entities.User, error) {
	m := ConnectToDB()
	coll := m.Database("db").Collection("user")

	filter := bson.D{{Key: "email", Value: username}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return []entities.User{}, err
	}

	var users []entities.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return []entities.User{}, err
	}

	defer DisconnectFromDB(m)

	return users, nil
}
