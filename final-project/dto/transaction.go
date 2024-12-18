package dto

import (
	"final-project/models"
	"time"
)

// type UserLogin struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type TransactionDetail struct {
// 	TransactionId     uint      `json:"id"`
// 	UserId            uint      `json:"user_id"`
// 	TransactionTypeId uint      `json:"transaction_type_id"`
// 	Description       string    `json:"description"`
// 	Amount            int       `json:"amount"`
// 	Date              time.Time `json:"date"`
// }

type TransactionDetail struct {
	TransactionId      uint      `json:"id"`
	UserId             uint      `json:"user_id"`
	// UserName           string    `json:"user_name"` // Nama user dari tabel User
	TransactionTypeId  uint      `json:"transaction_type_id"`
	// TransactionType    string    `json:"transaction_type"` // Nama tipe transaksi
	Description        string    `json:"description"`
	Amount             int       `json:"amount"`
	Date               time.Time `json:"date"`
}


// func NewTransactionDetail(user models.User, transaction models.Transaction) TransactionDetail {
// 	return TransactionDetail{
// 		TransactionId: transaction.TransactionId,
// 		UserId:        user.UserId,
// 		TransactionTypeId: transaction.TransactionTypeId,
// 		Description: transaction.Description,
// 		Amount: transaction.Amount,
// 		Date: transaction.Date,
// 	}
// }

func NewTransactionDetail(transaction models.Transaction) TransactionDetail {
	return TransactionDetail{
		TransactionId:     transaction.TransactionId,
		UserId:            transaction.UserId,
		TransactionTypeId: transaction.TransactionTypeId,
		Description:       transaction.Description,
		Amount:            transaction.Amount,
		Date:              transaction.Date,
	}
}


func NewListTransactions(transactions []models.Transaction) []TransactionDetail {
	listTransaction := []TransactionDetail{}
	for _, transaction := range transactions {
		transactionDetail := TransactionDetail{
			TransactionId:     transaction.TransactionId,
			UserId:            transaction.UserId,
			TransactionTypeId: transaction.TransactionTypeId,
			Description:       transaction.Description,
			Amount:            transaction.Amount,
			Date:              transaction.Date,
		}
		listTransaction = append(listTransaction, transactionDetail)
	}
	return listTransaction
}
