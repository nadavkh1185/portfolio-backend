package models

type Experience struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Paragraph string
}