package models

import (
	"encoding/json"
	"time"
)

// Sponsor represents a conference sponsor
type Sponsor struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Name        string    `json:"name" db:"name"`
	Website     string    `json:"website" db:"website"`
	Description string    `json:"description" db:"description"`
	Logo        string    `json:"logo" db:"logo"`
	ContactID   int       `json:"contact_id" db:"contact_id"`
}

// String is not required by pop and may be deleted
func (s Sponsor) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// Sponsors is not required by pop and may be deleted
type Sponsors []Sponsor

// String is not required by pop and may be deleted
func (s Sponsors) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}
