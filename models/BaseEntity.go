package models

import "time"

type BaseEntity struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt *time.Time
}
