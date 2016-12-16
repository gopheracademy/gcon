package models

import (
	"encoding/json"
	"time"
)

// PresentationSpeaker relates speakers to presentations
type PresentationSpeaker struct {
	ID             int       `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	PresentationID int       `json:"presentation_id" db:"presentation_id"`
	SpeakerID      int       `json:"speaker_id" db:"speaker_id"`
}

// String is not required by pop and may be deleted
func (p PresentationSpeaker) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// PresentationSpeakers is not required by pop and may be deleted
type PresentationSpeakers []PresentationSpeaker

// String is not required by pop and may be deleted
func (p PresentationSpeakers) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
