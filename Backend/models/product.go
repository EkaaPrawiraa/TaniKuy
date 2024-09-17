package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:500"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	CategoryID  uint
	Category    *Category `gorm:"foreignKey:CategoryID"`
	StoreID     uint
	Store       *Store    `gorm:"foreignKey:StoreID"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}