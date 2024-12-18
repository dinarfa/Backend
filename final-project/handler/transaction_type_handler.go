package handler

import (
	"final-project/models"
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionTypeHandler struct {
	service services.TransactionTypeServicesInterface
}

func NewTransactionTypeHandler(service services.TransactionTypeServicesInterface) *TransactionTypeHandler {
	return &TransactionTypeHandler{service: service}
}

// CreateTransactionType handler untuk membuat tipe transaksi
func (h *TransactionTypeHandler) CreateTransactionType(c *gin.Context) {
	var transactionType models.TransactionType
	if err := c.ShouldBindJSON(&transactionType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTransactionType, err := h.service.CreateTransactionType(transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTransactionType)
}

// GetAllTransactionTypes handler untuk mendapatkan semua tipe transaksi
func (h *TransactionTypeHandler) GetAllTransactionTypes(c *gin.Context) {
	transactionTypes, err := h.service.GetAllTransactionTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactionTypes)
}
