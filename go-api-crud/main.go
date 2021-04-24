package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hencya/go-api-crud/config"
	"github.com/hencya/go-api-crud/controller"
	"github.com/hencya/go-api-crud/middleware"
	"github.com/hencya/go-api-crud/repository"
	"github.com/hencya/go-api-crud/service"
	"gorm.io/gorm"
)

var (
	db                    *gorm.DB                         = config.SetupDatabaseConnection()
	userRepository        repository.UserRepository        = repository.NewUserRepository(db)
	destinationRepository repository.DestinationRepository = repository.NewDestinationRepository(db)
	jwtService            service.JWTService               = service.NewJWTService()
	userService           service.UserService              = service.NewUserService(userRepository)
	destinationService    service.DestinationService       = service.NewDestinationService(destinationRepository)
	authService           service.AuthService              = service.NewAuthService(userRepository)
	authController        controller.AuthController        = controller.NewAuthController(authService, jwtService)
	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	destinationController controller.DestinationController = controller.NewDestinationController(destinationService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	autoRoutes := r.Group("api/auth")
	{
		autoRoutes.POST("/login", authController.Login)
		autoRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	destinationRoutes := r.Group("api/destinations", middleware.AuthorizeJWT(jwtService))
	{
		destinationRoutes.GET("/", destinationController.All)
		destinationRoutes.POST("/", destinationController.Insert)
		destinationRoutes.GET("/:id", destinationController.FindByID)
		destinationRoutes.PUT("/:id", destinationController.Update)
		destinationRoutes.DELETE("/:id", destinationController.Delete)
	}

	r.Run()
}
