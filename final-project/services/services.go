package services

import (
	"final-project/models"
	"final-project/repositories"
)

type TransactionTypeServicesInterface interface {
	CreateTransactionType(transactionType models.TransactionType) (models.TransactionType, error)
	GetAllTransactionTypes() ([]models.TransactionType, error)
}

type transactionTypeService struct {
	repo repositories.TransactionTypeRepositoryInterface
}

func NewTransactionTypeService(repo repositories.TransactionTypeRepositoryInterface) TransactionTypeServicesInterface {
	return &transactionTypeService{repo: repo}
}

// CreateTransactionType implementasi logika untuk membuat tipe transaksi
func (s *transactionTypeService) CreateTransactionType(transactionType models.TransactionType) (models.TransactionType, error) {
	return s.repo.CreateTransactionType(transactionType)
}

// GetAllTransactionTypes implementasi logika untuk mendapatkan semua tipe transaksi
func (s *transactionTypeService) GetAllTransactionTypes() ([]models.TransactionType, error) {
	return s.repo.GetAllTransactionTypes()
}
