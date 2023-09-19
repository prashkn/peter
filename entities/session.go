package entities

import (
	"errors"
	"time"
)

func (s *Session) isExpired() bool {
	return s.ExpirationTime.Before(time.Now())
}

func (s *Session) GetCurrentUser() (User, error) {
	if !s.isExpired() {
		return s.User, nil
	}

	return User{}, errors.New("current session expired")
}
