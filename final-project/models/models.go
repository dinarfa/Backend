package models

import (
	"time"
)

type User struct {
	UserId   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Balance  int    `gorm:"not null" json:"balance"`

	// Transactions []Transaction `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"transactions"`
}

type TransactionType struct {
	TransactionTypeId uint   `gorm:"primaryKey;autoIncrement" json:"transaction_type_id"`
	Name              string `gorm:"not null" json:"name"`
	// Transactions []		Transaction		`gorm:"foreignKey:TransactionTypeId;constraint:OnDelete:CASCADE;" json:"transactions"`

}

type Transaction struct {
	TransactionId     uint            `gorm:"primaryKey;autoIncrement" json:"transaction_id"`
	UserId            uint            `gorm:"not null" json:"user_id"`
	User              User            `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"user"` // biar bisa ambil data user yg lengkaap
	TransactionTypeId uint            `gorm:"not null" json:"transaction_type_id"`
	TransactionType   TransactionType `gorm:"foreignKey:TransactionTypeId;constraint:OnDelete:CASCADE;" json:"transaction_type"`
	Description       string          `gorm:"not null" json:"description"`
	Amount            int             `gorm:"not null" json:"amount"`
	Date              time.Time       `gorm:"not null" json:"date"`
}
