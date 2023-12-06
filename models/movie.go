package models

import "time"

type Movie struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Genre     string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdate:false"`
	DeletedAt time.Time `gorm:"autoUpdate:false"`
}
