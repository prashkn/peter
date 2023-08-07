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

func Signup(username, password string) error {
	found, err := db.GetUsers(username)
	if err != nil {
		return err
	}

	if len(found) > 0 {
		return errors.New("found a user with the same username. Please use a different username")
	}

	err = db.InsertUser(username, password)
	return err
}
