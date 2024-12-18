package routes

import (


	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func Routes(r *gin.Engine, apiKey string, db *gorm.DB) {
func Routes(r *gin.Engine, db *gorm.DB) {
	user := r.Group("/users")
	transaction := r.Group("/transactions")
	transaction_type := r.Group("/transactions_types")
	// admin := r.Group("/admins")
	// movie := r.Group("/movies")
	UserRoutes(user, db)
	TransactionRoutes(transaction, db)
	TransactionTypeRoutes(transaction_type, db)
	// MovieRoutes(movie, apiKey, db)

}
