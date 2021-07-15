package application

import (
	"errors"
	"fmt"

	"github.com/Nikym/opponents-event-ms/domain/event"
	"github.com/google/uuid"
)

// EventApp is a struct that holds all of the event.Event related application functionality.
type EventApp struct {
	repository event.Repository
}

// NewEventApp creates a new instance of EventApp and returns a pointer to it.
func NewEventApp(repository event.Repository) *EventApp {
	return &EventApp{repository}
}

func appendStringToError(base error, extra string) error {
	return fmt.Errorf("%w: %s", base, extra)
}

// CreateEvent creates a new event.Event instance and saves it to the EventApp repository.
func (ea *EventApp) CreateEvent(title, description string, user event.User) (*event.Event, error) {
	e, err := event.New(title, description, user)
	if err != nil {
		return &event.Event{}, err
	}

	if err := ea.repository.Add(e); err != nil {
		return &event.Event{}, appendStringToError(event.ErrRepositoryAdd, err.Error())
	}

	return e, nil
}

// RemoveEvent deletes the event.Event with the given ID from the EventApp repository.
func (ea *EventApp) RemoveEvent(id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := ea.repository.Remove(parsedID); err != nil {
		return appendStringToError(event.ErrRepositoryRemove, err.Error())
	}

	return nil
}

// UpdateEvent updates the event.Event stored in the EventApp repository.
func (ea *EventApp) UpdateEvent(e event.Event) error {
	if err := ea.repository.Update(&e); err != nil {
		return appendStringToError(event.ErrRepositoryUpdate, err.Error())
	}

	return nil
}

// AddParticipantsToEvent adds the event.User instances given to the event.Event with the given ID.
// If even does not exist or fault occurs, returns error.
func (ea *EventApp) AddParticipantsToEvent(eventID string, participants ...event.User) error {
	parsedID, err := uuid.Parse(eventID)
	if err != nil {
		return err
	}

	e, err := ea.repository.Get(parsedID)
	if err != nil {
		return appendStringToError(event.ErrRepositoryGet, err.Error())
	}

	e.AddParticipants(participants...)

	if err := ea.repository.Update(e); err != nil {
		return appendStringToError(event.ErrRepositoryUpdate, err.Error())
	}

	return nil
}

// RemoveParticipantFromEvent removes the participant with the given ID from the event with the given ID.
// If the event or participant is not found, or there is a fault, then an error is returned.
func (ea *EventApp) RemoveParticipantFromEvent(eventID, participantID string) error {
	eParsedID, err := uuid.Parse(eventID)
	if err != nil {
		return err
	}

	pParsedID, err := uuid.Parse(participantID)
	if err != nil {
		return err
	}

	e, err := ea.repository.Get(eParsedID)
	if err != nil {
		return appendStringToError(event.ErrRepositoryGet, err.Error())
	}

	if ok := e.RemoveParticipant(pParsedID); !ok {
		return errors.New("event: participant not found")
	}

	if err := ea.repository.Update(e); err != nil {
		return appendStringToError(event.ErrRepositoryUpdate, err.Error())
	}

	return nil
}
