package models

import (
	"encoding/json"
	"time"
)

// Hotel represents hotel data for the conference
type Hotel struct {
	ID         int       `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	BlockCode  string    `json:"block_code" db:"block_code"`
	RoomRate   int       `json:"room_rate" db:"room_rate"`
	PhotoURL   string    `json:"photo_url" db:"photo_url"`
	LocationID int       `json:"location_id" db:"location_id"`
	ContactID  int       `json:"contact_id" db:"contact_id"`
	OnSale     bool      `json:"on_sale" db:"on_sale"`
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
