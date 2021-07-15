package event

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrRepositoryGet    = errors.New("event: could not get")
	ErrRepositoryAdd    = errors.New("event: could not add")
	ErrRepositoryRemove = errors.New("event: could not remove")
	ErrRepositoryUpdate = errors.New("event: could not update")
)

// Repository interface describing the functionality of an Event repo.
type Repository interface {
	// Get returns an Event or error if not found.
	Get(id uuid.UUID) (*Event, error)
	// GetAll returns a slice of all Event entities or error in case of failure.
	GetAll() ([]*Event, error)
	// Find returns an Event or nil if not found, and error in case of failure.
	Find(id uuid.UUID) (*Event, error)
	// Add inserts an Event into the repository.
	Add(e *Event) error
	// Remove deletes an Event with a given ID from the repository.
	Remove(id uuid.UUID) error
	// Update replaces the Event stored in the repository with the one specified.
	Update(e *Event) error
}
