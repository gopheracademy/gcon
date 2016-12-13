package models

import (
	"encoding/json"
	"time"
)

// Speaker represents a person speaking at the conference
type Speaker struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Bio       string    `json:"bio" db:"bio"`
	Website   string    `json:"website" db:"website"`
	Twitter   string    `json:"twitter" db:"twitter"`
	Linkedin  string    `json:"linkedin" db:"linkedin"`
	Facebook  string    `json:"facebook" db:"facebook"`
	ContactID int       `json:"contact_id" db:"contact_id"` // link to contact table
	PhotoURL  string    `json:"photo_url" db:"photo_url"`
}

// String is not required by pop and may be deleted
func (s Speaker) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// Speakers is not required by pop and may be deleted
type Speakers []Speaker

// String is not required by pop and may be deleted
func (s Speakers) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}
