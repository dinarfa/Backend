
	package routes

	import (
		"final-project/handler"
		"final-project/middleware"
		"final-project/repositories"
		"final-project/services"

		"github.com/gin-gonic/gin"
		"gorm.io/gorm"
	)

func UserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	svc := services.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	group.POST("/register", handler.Register)
	group.POST("/login", handler.Login)

	// Buat grup 'auth' sebagai RouterGroup terlebih dahulu
	auth := group.Group("/")
	auth.Use(middleware.JWTMiddleware())

	auth.GET("/:id", handler.GetProfile)
	auth.GET("", handler.GetAllUser)
	auth.DELETE("/:id", handler.Delete)
	auth.GET("/:id/balance", handler.GetBalance)
	auth.PUT("/:id", handler.Update)
}



