package repositories

import (
	"errors"
	"final-project/models"

	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionById(id uint) (models.Transaction, error)
	DeleteTransaction(id uint) error
	UpdateTransaction(id uint, updatedData models.Transaction) error
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryInterface {
	return &transactionRepo{
		db: db,
	}
}

// CreateTransaction: Menambahkan transaksi baru ke database
func (u *transactionRepo) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := u.db.Create(&transaction).Error
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

// GetAllTransactions: Mengambil semua data transaksi
func (u *transactionRepo) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := u.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

// GetTransactionById: Mengambil data transaksi berdasarkan ID
func (u *transactionRepo) GetTransactionById(id uint) (models.Transaction, error) {
	var transaction models.Transaction

	// Cari transaksi berdasarkan kolom transaction_id
	result := u.db.Where("transaction_id = ?", id).First(&transaction)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Transaction{}, errors.New("data tidak ditemukan")
		}
		return models.Transaction{}, result.Error
	}
	return transaction, nil
}

// UpdateTransaction: Memperbarui data transaksi berdasarkan ID
func (u *transactionRepo) UpdateTransaction(id uint, updatedData models.Transaction) error {
	// Gunakan kolom transaction_id untuk update
	err := u.db.Model(&models.Transaction{}).Where("transaction_id = ?", id).Updates(updatedData).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteTransaction: Menghapus data transaksi berdasarkan ID
func (u *transactionRepo) DeleteTransaction(id uint) error {
	// Gunakan kolom transaction_id untuk delete
	err := u.db.Where("transaction_id = ?", id).Delete(&models.Transaction{}).Error
	if err != nil {
		return err
	}
	return nil
}
