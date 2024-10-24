package models

import "time"

type DefaultDestination struct {
	ID uint      `gorm:"primaryKey"`
	UserID              uint      `gorm:"not null"`
	CityID              uint      `gorm:"not null"`
	CreatedAt           time.Time 
	DeletedAt           *time.Time

	User                Users      `gorm:"foreignKey:UserID;references:ID"`
	City                Cities     `gorm:"foreignKey:CityID;references:ID"`
}
