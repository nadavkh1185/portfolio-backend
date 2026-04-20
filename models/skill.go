package models

type Skill struct {
	ID         uint   `gorm:"primaryKey"`
	ImageURL   string `json:"image_url"`
	ImageTitle string `json:"image_title"`
}