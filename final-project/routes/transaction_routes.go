package routes

import (
	"final-project/handler"
	"final-project/middleware"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRoutes(group *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTransactionRepository(db)
	svc := services.NewTransactionService(repo)
	handler := handler.NewTransactionHandler(svc)

	// Buat grup 'auth' sebagai RouterGroup terlebih dahulu
	auth := group.Group("/")
	auth.Use(middleware.JWTMiddleware())

	// auth.GET("/:id", handler.GetProfile)
	auth.GET("/:id", handler.GetById)
	auth.GET("", handler.GetAll)
	auth.POST("", handler.Create)
	auth.DELETE("/:id", handler.Delete)
	// auth.GET("/me", handler.Me)
	auth.PUT("/:id", handler.Update)
}
