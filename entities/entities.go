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
	User     User
	Website  string
	Username string
	Email    string
	Hash     string
}

type Session struct {
	User User
}
