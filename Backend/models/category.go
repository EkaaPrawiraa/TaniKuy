package models

type Category struct {
	ID       uint       `gorm:"primaryKey;autoIncrement"`
	Name     string     `gorm:"size:255;not null;unique"`
	Products []*Product `gorm:"foreignKey:CategoryID"`
}