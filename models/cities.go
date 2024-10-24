package models

type Cities struct {
	BaseEntity
	Name        string
	Description string
	Users       []Users `gorm:"many2many:default_destinations;"`

}
