package models

type Skill struct {
	ID         uint `gorm:"primaryKey"`
	ImageURL   string
	ImageTitle string
}