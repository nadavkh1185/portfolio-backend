package models

type Project struct {
	ID       uint `gorm:"primaryKey"`
	ImageURL string
	Subtitle string
}