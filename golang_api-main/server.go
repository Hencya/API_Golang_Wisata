package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/controller"
	"github.com/ydhnwb/golang_api/middleware"
	"github.com/ydhnwb/golang_api/repository"
	"github.com/ydhnwb/golang_api/service"
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

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
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
