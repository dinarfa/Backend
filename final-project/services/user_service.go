package services

import (
	"errors"
	"final-project/middleware"
	"final-project/models"
	"final-project/repositories"
)

type UserServiceInterface interface {
	Login(user models.User) (string, error)
	Register(user models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	Update(id uint, updatedData models.User) error
	Delete(id uint) error
	GetBalance(id uint) (int, error)
}

type userService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(rp repositories.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		repo: rp,
	}
}

func (u *userService) GetBalance(id uint) (int, error) {
	// Panggil repository untuk mendapatkan saldo terkini
	balance, err := u.repo.GetBalance(id)
	if err != nil {
		return 0, err
	}
	return balance, nil
}


// Register implements UserService.
func (u *userService) Register(user models.User) (models.User, error) {
	existingUser, err := u.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser.UserId != 0 {
		return models.User{}, errors.New("Email sudah terpakai")
	}

	hassPass, err := middleware.HassPass(user.Password)
	if err != nil {
		return models.User{}, err //
	}
	user.Password = hassPass
	if err := u.repo.CreateUser(user); err != nil {
		return models.User{}, err //
	}

	return user, nil
}

// Login implements UserService.
func (u *userService) Login(user models.User) (string, error) {
	existingUser, err := u.repo.GetUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	ok, err := middleware.ComparePass([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("Password salah")

	}
	if existingUser.UserId == 0 {
		return "", errors.New("Data tidak ditemukan")
	}
	token, err := middleware.CreateTokenJWT(int(user.UserId))
	if err != nil {
		return "", err
	}
	return token, nil

}


// GetAll implements UserService.
// ambil semua user
func (u *userService) GetAll() ([]models.User, error) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById implements UserService.
func (u *userService) GetUserById(id uint) (models.User, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Update implements UserService.
func (u *userService) Update(id uint, updatedData models.User) error {
	err := u.repo.UpdateUser(id, updatedData)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements UserService.
func (u *userService) Delete(id uint) error {
	err := u.repo.DeleteUser(id)
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