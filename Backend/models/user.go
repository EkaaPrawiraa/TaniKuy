package models

import "time"

type User struct {
    UserID       uint      `gorm:"primaryKey;autoIncrement"`
    PhotoProfile string    // Firebase saving image
    Name         string    `gorm:"size:255;not null"`
    Username     string    `gorm:"size:20;unique"`
    Email        string    `gorm:"size:255;not null;unique"`
    Password     string    `gorm:"size:255;not null"`
    Phone        string    `gorm:"size:15"`
    Gender       string    `gorm:"size:6"` // "male" or "female"
    Birth        time.Time
    CreatedDate  time.Time `gorm:"autoCreateTime"`
    Addresses    []*Address
    BankAccounts []*Bank
    UserType     string    `gorm:"size:5;not null"` // "user" or "admin"
    StoreID      uint
    Store        *Store    `gorm:"foreignKey:StoreID"`
}