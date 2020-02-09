package model

import (
	"time"
)

// Anime model for all animes
type Anime struct {
	ID int `bson:"id"`
	Name	string `bson:"name"`
	Current	int `bson:"current"`
	Total	int `bson:"total"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Animes type for export
type Animes []Anime