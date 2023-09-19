package db

import (
	"context"
	"peter/entities"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {
	const uri = "mongodb+srv://pkn3:ilpaf5341@cluster0.gw2u9ol.mongodb.net/"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

func DisconnectFromDB(m *mongo.Client) {
	if err := m.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

// TODO
func StartSession() (entities.Session, error) {
	return entities.Session{}, nil
}
