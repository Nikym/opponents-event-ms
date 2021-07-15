package event

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	minTitleLength = 3
	maxTitleLength = 50
)

var (
	ErrInvalidEvent  = errors.New("event: could not use invalid event")
	ErrTitleTooShort = fmt.Errorf("%w: min length for title is %d", ErrInvalidEvent, minTitleLength)
	ErrTitleTooLong  = fmt.Errorf("%w: max length for title is %d", ErrInvalidEvent, maxTitleLength)
)

type (
	Title       string
	Description string
)

// Event struct describes an event entity.
type Event struct {
	ID           uuid.UUID
	Title        Title
	Description  Description
	Host         User
	Participants []User
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// New created a new Event instance and returns a pointer to it.
func New(title, description string, host User) (*Event, error) {
	switch l := len(strings.TrimSpace(title)); {
	case l < minTitleLength:
		return &Event{}, ErrTitleTooLong
	case l > maxTitleLength:
		return &Event{}, ErrTitleTooShort
	}

	return &Event{
		ID:           uuid.New(),
		Title:        Title(title),
		Description:  Description(description),
		Host:         host,
		Participants: make([]User, 0),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

// AddParticipants adds User value types to the Event.
func (e *Event) AddParticipants(users ...User) {
	e.Participants = append(e.Participants, users...)
	e.UpdatedAt = time.Now()
}

// RemoveParticipant removes a User from the list of participants if given ID matches.
// Returns true if successful, false otherwise.
func (e *Event) RemoveParticipant(id uuid.UUID) bool {
	for i, u := range e.Participants {
		if u.ID == id {
			e.Participants[i] = e.Participants[len(e.Participants)-1]
			e.Participants = e.Participants[:len(e.Participants)-1]
			e.UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

// ChangeDescription changes the Event description to a given Description value.
func (e *Event) ChangeDescription(description Description) {
	e.Description = description
	e.UpdatedAt = time.Now()
}

// Rename changes the Event title to a given Title value.
func (e *Event) Rename(title Title) {
	e.Title = title
	e.UpdatedAt = time.Now()
}
