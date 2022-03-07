package objects

import (
	"encoding/json"
	"net/http"
)

// MaxListlmit maximum number of listings
const MaxListlmit = 200

// GetRequest for single Event object
type GetRequest struct {
	ID string `json:"id"`
}

// ListRequest for list of event objects
type ListRequest struct {
	Limit int    `json:"limit"`
	After string `json:"after"`
	//Optional name matching
	Name string `json:"name"`
}

// CreateRequest for new Event object
type CreateRequest struct {
	Event *Event `json:"event"`
}

// UpdateDetailsRequest  to update details of an Event object
type UpdateDetailsRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone"`
}

// CancelRequest to cancel an Event object
type CancelRequest struct {
	ID string `json:"id"`
}

// RescheduleRequest to reschedule an Event object
type RescheduleRequest struct {
	ID      string    `json:"id"`
	NewSlot *TimeSlot `json:"new_slot"`
}

// DeleteRequest to delete an Event object
type DeleteRequest struct {
	ID string `json:"id"`
}

// EventResponseWrapper response to an Event request
type EventResponseWrapper struct {
	Event  *Event   `json:"event,omitempty"`
	Events []*Event `json:"events,omitempty"`
	Code   int      `json:"-"`
}

// JSON convert EventResponseWrapper in JSON
func (e *EventResponseWrapper) JSON() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

// StatusCode returns status code of the event
func (e *EventResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOk
	}
	return e.Code
}
