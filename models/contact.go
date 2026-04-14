package models

type Contact struct {
	ID            uint `gorm:"primaryKey"`
	LinkedinLink  string
	InstagramLink string
}