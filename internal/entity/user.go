package entity

import "time"

type User struct {
	ID        string    `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Address   string    `bson:"address" json:"address"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
