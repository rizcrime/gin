package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizcrime/gin_trial/config"
	"github.com/rizcrime/gin_trial/controller"
	"github.com/rizcrime/gin_trial/repository"
	"github.com/rizcrime/gin_trial/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService service.JWTService = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
