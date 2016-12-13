package models

import "encoding/json"

/*
Sponsor Name : Google
Sponsor Website: https://www.google.com
Sponsor Description : Google invented Go and is happy to sponsor Gophercon 2017
Sponsor Contact Name: Steve Francia
Sponsor Contact Phone: 555-1212
Sponsor Contact Email : spf@golang.org
Sponsor Logo Path(or URL?) : /assets/sponsors/platinum/google.png
*/

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
