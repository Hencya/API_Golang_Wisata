package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hencya/go-api-crud/config"
	"github.com/hencya/go-api-crud/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	autoRoutes := r.Group("api/auth")
	{
		autoRoutes.POST("/login", authController.Login)
		autoRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
