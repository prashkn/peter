package resolver

import (
	"errors"
	"peter/db"
	"peter/entities"
)

func Login(username, password string) (entities.User, error) {
	found, err := db.GetUsers(username)
	if err != nil {
		return entities.User{}, err
	}

	if len(found) == 0 {
		return entities.User{}, errors.New("found no users with that username. Please sign up")
	}

	return found[0], err
}

func Signup(email, password string) error {
	found, err := db.GetUsers(email)
	if err != nil {
		return err
	}

	if len(found) > 0 {
		return errors.New("found a user with the same email. Please use a different email address")
	}

	err = db.InsertUser(email, password)
	return err
}
