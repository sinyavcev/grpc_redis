package models

import (
	"time"
)

type User struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Phone     string    `json:"phone,omitempty" bson:"phone,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
