package models

type About struct {
	ID         uint   `gorm:"primaryKey"`
	ImageURL   string `json:"image_url"`
	Subtitle   string `json:"subtitle"`
	Paragraph1 string `json:"paragraph1"`
	Paragraph2 string `json:"paragraph2"`
	Paragraph3 string `json:"paragraph3"`
}