package models

type Bank struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	UserID      uint   // Foreign key to User
	User        *User  `gorm:"foreignKey:UserID"`
	BankName    string `gorm:"size:255"`
	AccountNo   string `gorm:"size:255"`
	AccountType string `gorm:"size:255"`
}