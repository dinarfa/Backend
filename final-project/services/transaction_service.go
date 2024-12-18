package services

import (
	// "errors"

	"errors"
	"final-project/models"
	"final-project/repositories"
)

type TransactionServiceInterface interface {
	// Login(user models.User) (string, error)
	// Register(user models.User) (models.User, error)
	Create(transaction models.Transaction) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	GetTransactionById(id uint) (models.Transaction, error)
	Update(id uint, updatedData models.Transaction) error
	Delete(id uint) error
}

type transactionService struct {
	repo repositories.TransactionRepositoryInterface
}

func NewTransactionService(rp repositories.TransactionRepositoryInterface) TransactionServiceInterface {
	return &transactionService{
		repo: rp,
	}
}

func (u *transactionService) Create(transaction models.Transaction) (models.Transaction, error) {
	transaction, err := u.repo.CreateTransaction(transaction)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

// GetAll implements UserService.
// ambil semua user
func (u *transactionService) GetAll() ([]models.Transaction, error) {
	transactions, err := u.repo.GetAllTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetUserById implements UserService.
func (u *transactionService) GetTransactionById(id uint) (models.Transaction, error) {
	transaction, err := u.repo.GetTransactionById(id)
	if err != nil {
		return models.Transaction{}, err // Mengembalikan error dari repository
	}

	// Validasi jika data tidak ditemukan (dengan asumsi `ID` adalah field unik)
	if transaction.TransactionId == 0 {
		return models.Transaction{}, errors.New("data tidak ditemukan")
	}

	return transaction, nil
}

// Update implements UserService.
func (u *transactionService) Update(id uint, updatedData models.Transaction) error {
	err := u.repo.UpdateTransaction(id, updatedData)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements UserService.
func (u *transactionService) Delete(id uint) error {
	err := u.repo.DeleteTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

// func (u *userService) AddToWatchlist(userID uint, movieID uint) error {
// 	return u.repo.AddToWatchlist(userID, movieID)
// }

// func (u *userService) RemoveFromWatchlist(userID uint, movieID uint) error {
// 	return u.repo.RemoveFromWatchlist(userID, movieID)
// }
