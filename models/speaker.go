package models

import (
	"encoding/json"
	"github.com/markbates/pop/nulls"
	"time"
)

type Speaker struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	FirstName string       `json:"first_name" db:"first_name"`
	LastName  string       `json:"last_name" db:"last_name"`
	Twitter   nulls.String `json:"twitter" db:"twitter"`
	Linkedin  nulls.String `json:"linkedin" db:"linkedin"`
	Facebook  nulls.String `json:"facebook" db:"facebook"`
	Imagepath nulls.String `json:"imagepath" db:"imagepath"`
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
