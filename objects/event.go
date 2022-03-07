package objects

import "time"

// EventStatus defines the status of the event.
type EventStatus string

//Default event status
const (
	Original    EventStatus = "original"
	Cancelled   EventStatus = "cancelled"
	Rescheduled EventStatus = "rescheduled"
)

// TimeSlot for Event
type TimeSlot struct {
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

//Event Object for the API
type Event struct {
	//ID of the event
	ID string `gorm:"primary_key json:"id,omitempty"`

	//General Details about the event
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Website     string `json:"website,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`

	//Event Slot duration
	Slot *TimeSlot `gorm:"embedded" json:"slot,omitempty"`

	//Change status of the event
	Status EventStatus `json:"status,omitempty"`

	//Metadata about the event
	CreatedOn     time.Time `json:"created_on,omitempty"`
	UpdatedOn     time.Time `json:"updated_on,omitempty"`
	CancelledOn   time.Time `json:"cancelled_on,omit"`
	RescheduledOn time.Time `json:"rescheduled_on,omitempty"`
}
