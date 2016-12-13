package models

import (
	"encoding/json"
	"time"
)

// Hotel represents hotel data for the conference
type Hotel struct {
	ID               int       `json:"id" db:"id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	PropertyName     string    `json:"property_name" db:"property_name"`
	Address          string    `json:"address" db:"address"`
	City             string    `json:"city" db:"city"`
	State            string    `json:"state" db:"state"`
	Zip              string    `json:"zip" db:"zip"`
	ReservationPhone string    `json:"reservation_phone" db:"reservation_phone"`
	BlockCode        string    `json:"block_code" db:"block_code"`
	RoomRate         string    `json:"room_rate" db:"room_rate"`
	PhotoURL         string    `json:"photo_url" db:"photo_url"`
}

// String is not required by pop and may be deleted
func (h Hotel) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}

// Hotels is not required by pop and may be deleted
type Hotels []Hotel

// String is not required by pop and may be deleted
func (h Hotels) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}
