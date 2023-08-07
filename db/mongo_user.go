package db

import (
	"context"
	"fmt"
	"peter/entities"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(username, password string) error {
	m := ConnectToDB()
	coll := m.Database("db").Collection("user")

	user := entities.User{
		Name:     username,
		Password: password,
	}

	_, err := coll.InsertOne(context.Background(), user)

	defer DisconnectToDB(m)

	return err
}

func GetUsers(username string) ([]entities.User, error) {
	m := ConnectToDB()
	coll := m.Database("db").Collection("user")

	filter := bson.D{{Key: "name", Value: username}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return []entities.User{}, err
	}

	var users []entities.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return []entities.User{}, err
	}

	fmt.Printf("%+v", users)

	return users, nil
}
