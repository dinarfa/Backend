package repositories

import (
	"database/sql"
	"errors"
	"final-project/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, updatedData models.User) error
	GetBalance(id uint) (int, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepo{
		db: db,
	}
}

// GetBalance implements UserRepositoryInterface.
func (u *userRepo) GetBalance(id uint) (int, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("user dengan id %d tidak ditemukan", id)
		}
		return 0, err
	}

	// Ambil total pemasukan
	var income sql.NullInt64
	if err := u.db.Model(&models.Transaction{}).
		Where("user_id = ? AND transaction_type_id = ?", id, 1).
		Select("SUM(amount)").
		Scan(&income).Error; err != nil {
		return 0, err
	}
	// Periksa apakah income adalah NULL
	if !income.Valid {
		income.Int64 = 0
	}

	// Ambil total pengeluaran
	var expense sql.NullInt64
	if err := u.db.Model(&models.Transaction{}).
		Where("user_id = ? AND transaction_type_id = ?", id, 2).
		Select("SUM(amount)").
		Scan(&expense).Error; err != nil {
		return 0, err
	}
	// Periksa apakah expense adalah NULL
	if !expense.Valid {
		expense.Int64 = 0
	}

	// Hitung saldo terkini
	currentBalance := int(user.Balance) + int(income.Int64) - int(expense.Int64)
	return currentBalance, nil
}
// CreateUser implements UserRepositoryInterface.
func (u *userRepo) CreateUser(user models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return fmt.Errorf("gagal membuat user: %w", err)
	}
	return nil
}

// GetAllUsers implements UserRepositoryInterface
func (u *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("gagal mendapatkan daftar user: %w", err)
	}
	return users, nil
}

// GetUserByEmail implements UserRepositoryInterface.
func (u *userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user dengan email %s tidak ditemukan", email)
		}
		return user, err
	}
	return user, nil
}

// GetUserById implements UserRepositoryInterface.
func (u *userRepo) GetUserById(id uint) (models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user dengan id %d tidak ditemukan", id)
		}
		return user, err
	}
	return user, nil
}

// UpdateUser implements UserRepositoryInterface.
func (u *userRepo) UpdateUser(id uint, updatedData models.User) error {
	if err := u.db.Model(&models.User{}).Where("user_id = ?", id).Updates(updatedData).Error; err != nil {
		return fmt.Errorf("gagal memperbarui data user dengan id %d: %w", id, err)
	}
	return nil
}

// DeleteUser implements UserRepositoryInterface.
func (u *userRepo) DeleteUser(id uint) error {
	if err := u.db.Delete(&models.User{}, id).Error; err != nil {
		return fmt.Errorf("gagal menghapus user dengan id %d: %w", id, err)
	}
	return nil
}
