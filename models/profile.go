package models

type Profile struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Career   string
	Headline string
	Line     string
}