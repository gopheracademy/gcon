package models

import (
	"encoding/json"
	"time"
)

// Presentation represents a presentation/session
type Presentation struct {
	ID               int       `json:"id" db:"id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	Title            string    `json:"title" db:"title"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	LongDescription  string    `json:"long_description" db:"long_description"`
	SpeakerIDs       []int     `json:"speaker_ids"`
}

// String is not required by pop and may be deleted
func (p Presentation) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// Presentations is not required by pop and may be deleted
type Presentations []Presentation

// String is not required by pop and may be deleted
func (p Presentations) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
