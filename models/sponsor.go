package models

import "encoding/json"

// Sponsor
type Sponsor struct {
	Name         string `json:"name" db:"name"`
	Website      string `json:"website" db:"website"`
	Description  string `json:"description" db:"description"`
	ContactName  string `json:"contact_name" db:"contact_name"`
	ContactPhone string `json:"contact_phone" db:"contact_phone"`
	ContactEmail string `json:"contact_email" db:"contact_email"`
	Logo         string `json:"contact_logo" db:"contact_logo"`
}

// String is not required by pop and may be deleted
func (s Sponsor) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// Sponsors
type Sponsors []Sponsor

// String
func (s Sponsors) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}
