package event

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	minUsernameLength = 3
)

var (
	ErrInvalidUser      = errors.New("user: could not use invalid user")
	ErrUsernameTooShort = fmt.Errorf("%w: min length for username allowed is %d", ErrInvalidUser, minUsernameLength)
	ErrInvalidUUID      = fmt.Errorf("%w: ID is an invalid UUID", ErrInvalidUser)
)

type (
	Username string
)

// User struct describes a user registered to the event.
type User struct {
	ID       uuid.UUID
	Username Username
}

// NewUser creates a new User instance and returns it.
func NewUser(id string, username string) (User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return User{}, ErrInvalidUUID
	}
	if len(strings.TrimSpace(username)) < minUsernameLength {
		return User{}, ErrUsernameTooShort
	}

	return User{
		ID:       parsedID,
		Username: Username(username),
	}, nil
}

// String returns a string representation of the User (i.e. their username).
func (u User) String() string {
	return string(u.Username)
}

// Equals compares two User instances to see if their ID matches.
func (u User) Equals(u2 User) bool {
	return u.ID.String() == u2.ID.String()
}
