package handler

import (
	// "encoding/json"
	"final-project/dto"
	"final-project/models"
	"strconv"

	// "strconv"

	// "final-project/repositories"
	"final-project/services"
	// "log"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	// GetProfile(c *gin.Context)
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type transactionHandler struct{
	serv services.TransactionServiceInterface
}

func NewTransactionHandler(serv services.TransactionServiceInterface) TransactionHandler {
	return &transactionHandler{serv: serv}
}

func (u *transactionHandler) Create(c *gin.Context){
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mohon input dengan benar"}) // klo kosong ya eror
		return
	}

	_, err := u.serv.Create(transaction)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Mohon Input dengan benar"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaksi berhasil ditambahkan",
	})
}

func (u *transactionHandler) GetById(c *gin.Context) {
    transactionIDStr := c.Param("id") // Mengambil ID sebagai string

    // Konversi dari string ke uint
    transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "ID harus berupa angka positif",
        })
        return
    }

    // Panggil service dengan ID yang sudah dikonversi
    transaction, err := u.serv.GetTransactionById(uint(transactionID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
        return
    }

    // Gunakan DTO atau langsung respons
    transactionResp := dto.NewTransactionDetail(transaction)
    c.JSON(http.StatusOK, transactionResp)
}


func (u *transactionHandler) Update(c *gin.Context) {
	var updatedTransaction models.Transaction

	// Ambil parameter id dari URL
	idStr := c.Param("id")

	// Konversi dari string ke uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka positif",
		})
		return
	}

	// Bind data JSON ke updatedTransaction
	if err := c.ShouldBindJSON(&updatedTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Mohon Input dengan benar",
		})
		return
	}

	// Panggil service untuk melakukan update
	err = u.serv.Update(uint(id), updatedTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaksi Berhasil di-update",
	})
}


// Delete implements UserHandler.
// func (u *transactionHandler) Delete(c *gin.Context) {
// 	// Ambil ID dari parameter URL
// 	idStr := c.Param("id")

// 	// Konversi dari string ke uint
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "ID harus berupa angka positif",
// 		})
// 		return
// 	}

// 	// Panggil service untuk menghapus transaksi
// 	err = u.serv.Delete(uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Gagal menghapus Transaksi. Silakan coba lagi.",
// 		})
// 		return
// 	}

// 	// Response sukses
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Transaksi berhasil dihapus",
// 	})
// }

func (u *transactionHandler) Delete(c *gin.Context) {
	// Ambil ID dari parameter URL
	idStr := c.Param("id")

	// Konversi dari string ke uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka positif",
		})
		return
	}

	// Periksa apakah transaksi dengan ID tersebut ada
	transaction, err := u.serv.GetTransactionById(uint(id))
	if err != nil {
		// Jika transaksi tidak ditemukan, kirimkan pesan error dengan status Not Found
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Transaksi dengan ID tersebut tidak ditemukan",
		})
		return
	}

	// Panggil service untuk menghapus transaksi
	err = u.serv.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus Transaksi. Silakan coba lagi.",
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaksi berhasil dihapus",
		"transaction_id": transaction.TransactionId,
	})
}


// GetAlltransaction implements transactionHandler.
func (u *transactionHandler) GetAll(c *gin.Context) {
	transactions, err := u.serv.GetAll()
	// klo ada eror, kirim pesan baru stop
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	respTransaction := dto.NewListTransactions(transactions)

	c.JSON(http.StatusOK, respTransaction)
}
