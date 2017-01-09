package models

import (
	"encoding/json"
	"time"
)

// Location represents a place for events or other types
type Location struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	City      string    `json:"city" db:"city"`
	State     string    `json:"state" db:"state"`
	Zip       string    `json:"zip" db:"zip"`
}

// String is not required by pop and may be deleted
func (l Location) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// Locations is not required by pop and may be deleted
type Locations []Location

// String is not required by pop and may be deleted
func (l Locations) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}
