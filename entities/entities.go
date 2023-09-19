package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                 primitive.ObjectID  `bson:"_id"`
	EncryptedPasswords map[string]Password `bson:"encryptedPasswords"`
	Email              string              `bson:"email"`
	Validated          bool                `bson:"validated"`
	Password           string              `bson:"password"`
}

type Password struct {
	ID       primitive.ObjectID `bson:"_id"`
	User     User               `bson:"user"`
	Website  string             `bson:"website"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Hash     string             `bson:"hash"`
}

type Session struct {
	User User
}
