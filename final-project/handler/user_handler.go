package handler

import (
	"final-project/dto"
	"final-project/models"
	"final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
	GetAllUser(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Me(c *gin.Context)
	GetBalance(c *gin.Context)
}

type userHandler struct {
	serv services.UserServiceInterface
}

func NewUserHandler(serv services.UserServiceInterface) UserHandler {
	return &userHandler{serv: serv}
}

// GetBalance implements UserHandler.
func (u *userHandler) GetBalance(c *gin.Context) {
	// Konversi ID dari string ke uint
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Memanggil service untuk mendapatkan saldo
	balance, err := u.serv.GetBalance(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan saldo terkini
	c.JSON(http.StatusOK, gin.H{"current_balance": balance})
}

// Update implements UserHandler.
func (u *userHandler) Update(c *gin.Context) {
	var updatedUser models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mohon input dengan benar"})
		return
	}

	err = u.serv.Update(uint(id), updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil di-update"})
}

// Delete implements UserHandler.
func (u *userHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = u.serv.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Menghapus user"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil di-hapus"})
}

// GetAllUser implements UserHandler.
func (u *userHandler) GetAllUser(c *gin.Context) {
	users, err := u.serv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respUser := dto.NewListUsers(users)
	c.JSON(http.StatusOK, respUser)
}

// GetProfile implements UserHandler.
func (u *userHandler) GetProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := u.serv.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	userResp := dto.NewUserDetail(user)
	c.JSON(http.StatusOK, userResp)
}

// Login implements UserHandler.
func (u *userHandler) Login(c *gin.Context) {
    var user dto.UserLogin
    // Ambil data JSON dari request body dan isi ke dalam struct user
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format request tidak valid"})
        return
    }

    // Buat objek models.User dari data DTO (Email dan Password)
    userModel := models.User{
        Email:    user.Email,
        Password: user.Password,
    }

    // Panggil service Login dengan satu parameter (userModel)
    token, err := u.serv.Login(userModel) // Pastikan Login di service menerima models.User
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login berhasil",
        "token":   token, // Kalau sukses dapat token
    })
}


// Me implements UserHandler.
func (u *userHandler) Me(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := u.serv.GetUserById(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userResp := dto.NewUserDetail(user)
	c.JSON(http.StatusOK, userResp)
}

// Register implements UserHandler.
func (u *userHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mohon input dengan benar"})
		return
	}

	_, err := u.serv.Register(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil"})
}
