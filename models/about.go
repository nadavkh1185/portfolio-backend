package models

type About struct {
	ID         uint `gorm:"primaryKey"`
	ImageURL   string
	Subtitle   string
	Paragraph1 string
	Paragraph2 string
	Paragraph3 string
}