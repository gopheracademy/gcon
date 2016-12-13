package models

import (
	"encoding/json"
	"time"
)

// Contact represents an entity used for communications with a number of components
type Contact struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
}

// String is not required by pop and may be deleted
func (c Contact) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// Contacts is not required by pop and may be deleted
type Contacts []Contact

// String is not required by pop and may be deleted
func (c Contacts) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
