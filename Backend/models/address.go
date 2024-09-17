package models

type Address struct {
    ID       uint   `gorm:"primaryKey;autoIncrement"`
    UserID   uint   `gorm:"index"` // Create an index on UserID for faster queries
    User     *User  `gorm:"foreignKey:UserID"`
    Street   string `gorm:"size:255"`
    City     string `gorm:"size:255"`
    ZipCode  string `gorm:"size:20"` 
    Country  string `gorm:"size:255"`
}