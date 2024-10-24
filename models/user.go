package models

import "time"

type Users struct {
	BaseEntity
	FirebaseId string
	Name       string
	Email      string
	CreatedAt  time.Time
	DeletedAt  *time.Time
	Cities     []Cities `gorm:"many2many:default_destinations;"`
}
