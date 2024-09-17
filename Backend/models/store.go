package models

type Store struct {
	ID       uint       `gorm:"primaryKey;autoIncrement"`
	Name     string     `gorm:"size:255;not null"`
	Location string     `gorm:"size:255"`
	Products []*Product `gorm:"foreignKey:StoreID"`
}