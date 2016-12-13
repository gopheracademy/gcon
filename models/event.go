package models

import (
	"encoding/json"
	"time"
)

// Event represents an occurrance of note during the time of the conference
type Event struct {
	ID               int       `json:"id" db:"id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	Name             string    `json:"name" db:"name"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	LongDescription  string    `json:"long_description" db:"long_description"`
	Date             time.Time `json:"date" db:"date"`
	StartTime        time.Time `json:"start_time" db:"start_time"`
	EndTime          time.Time `json:"end_time" db:"end_time"`
	LocationID       int       `json:"location_id" db:"location_id"`
	GoogleMapsLink   string    `json:"google_maps_link" db:"google_maps_link"`
}

// String is not required by pop and may be deleted
func (e Event) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
