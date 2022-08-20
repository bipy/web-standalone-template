package models

import "time"

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name,omitempty"`
	CreateAt time.Time `json:"-"`
}
