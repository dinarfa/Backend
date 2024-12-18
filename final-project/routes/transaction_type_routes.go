package routes

import (
	"final-project/handler"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionTypeRoutes(group *gin.RouterGroup, db *gorm.DB) {
	// Inisialisasi repository, service, dan handler
	repo := repositories.NewTransactionTypeRepository(db)
	svc := services.NewTransactionTypeService(repo)
	handler := handler.NewTransactionTypeHandler(svc)

	// Daftar route untuk transaction type
	group.POST("/", handler.CreateTransactionType) // Create transaction type
	group.GET("/", handler.GetAllTransactionTypes) // Get all transaction types
}
