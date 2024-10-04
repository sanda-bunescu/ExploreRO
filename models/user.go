package models

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	FirebaseId string
	Name       string
	Email      string
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
