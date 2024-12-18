package repositories

import (
	// "errors"
	"final-project/models"

	"gorm.io/gorm"
)

type TransactionTypeRepositoryInterface interface {
	CreateTransactionType(transactionType models.TransactionType) (models.TransactionType, error)
	GetAllTransactionTypes() ([]models.TransactionType, error)
	// GetUserById(id string) (models.User, error)
	// GetTransactionTypeById(email string) (models.TransactionType, error)
	// DeleteTransactionType(id string) error
	// UpdateTransactionType(id string, updatedData models.TransactionType) error
}

type transactionTypeRepo struct {
	db *gorm.DB
}

func NewTransactionTypeRepository(db *gorm.DB) TransactionTypeRepositoryInterface {
	return &transactionTypeRepo{
		db: db,
	}
}

// CreateUser implements UserRepositoryInterface.
func (u *transactionTypeRepo) CreateTransactionType(transactionType models.TransactionType) (models.TransactionType, error) {
	err := u.db.Create(&transactionType).Error
	if err != nil {
		return models.TransactionType{}, err
	}
	return transactionType, nil

}

// GetAlltransactionTypes implements transactionTypeRepositoryInterface
func (u *transactionTypeRepo) GetAllTransactionTypes() ([]models.TransactionType, error) {
	var transactionTypes []models.TransactionType
	result := u.db.Find(&transactionTypes)
	return transactionTypes, result.Error
}

// GettransactionTypeByEmail implements transactionTypeRepositoryInterface.
// func (u *transactionTypeRepo) GetTransactionTypeByEmail(email string) (models.TransactionType, error) {
// 	var transactionType models.TransactionType

// 	result := u.db.Debug().Where("email = ?", email).First(&transactionType)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return transactionType, nil
// 		}
// 		return transactionType, result.Error
// 	}
// 	return transactionType, nil
// }

// GettransactionTypeById implements transactionTypeRepositoryInterface.
// func (u *transactionTypeRepo) GetTransactionTypeById(id string) (models.TransactionType, error) {
// 	var transactionType models.TransactionType

// 	result := u.db.First(&transactionType, id)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return transactionType, nil
// 		}
// 		return transactionType, result.Error
// 	}
// 	return transactionType, nil
// }

// // Update implements transactionTypeRepositoryInterface.
// func (u *transactionTypeRepo) UpdateTransactionType(id string, updatedData models.TransactionType) error {
// 	err := u.db.Model(&models.TransactionType{}).Where("id = ?", id).Updates(updatedData).Error
// 	if err != nil { //jika ada eror
// 		return err
// 	}
// 	return nil

// }

// // Delete implements transactionTypeRepositoryInterface.
// func (u *transactionTypeRepo) DeleteTransactionType(id string) error {
// 	err := u.db.Delete(&models.TransactionType{}, id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // func (u *transactionTypeRepo) AddToWatchlist(transactionTypeID uint, movieID uint) error {
// // 	var transactionType models.transactionType
// // 	var movie models.Movie

// // 	if err := u.db.First(&transactionType, transactionTypeID).Error; err != nil {
// // 		return err
// // 	}
// // 	if err := u.db.First(&movie, movieID).Error; err != nil {
// // 		return err
// // 	}

// // 	return u.db.Model(&transactionType).Association("Watchlist").Append(&movie)
// // }

// // func (u *transactionTypeRepo) RemoveFromWatchlist(transactionTypeID uint, movieID uint) error {
// // 	var transactionType models.transactionType
// // 	var movie models.Movie

// // 	if err := u.db.First(&transactionType, transactionTypeID).Error; err != nil {
// // 		return err
// // 	}
// // 	if err := u.db.First(&movie, movieID).Error; err != nil {
// // 		return err
// // 	}

// // 	return u.db.Model(&transactionType).Association("Watchlist").Delete(&movie)
// // }
