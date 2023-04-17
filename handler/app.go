package handler

import (
	"go-jwt/database"
	"go-jwt/repository/product_repository/product_pg"
	"go-jwt/repository/user_repository/user_pg"
	"go-jwt/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	productRepo := product_pg.NewProductPG(db)

	productService := service.NewProductService(productRepo)

	productHandler := NewProductHandler(productService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, productRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	productRoute := route.Group("/products")
	{
		productRoute.GET("/", authService.Authentication(), productHandler.GetProduct)
		productRoute.GET("/:productId", authService.Authentication(), authService.Authorization(), productHandler.GetProductById)
		productRoute.POST("/", authService.Authentication(), productHandler.CreateProduct)

		productRoute.PUT("/:productId", authService.Authentication(), authService.Authorization(), productHandler.UpdateProductById)

		productRoute.DELETE("/:productId", authService.Authentication(), authService.Authorization(), productHandler.deleteProduct)

	}

	route.Run(":" + port)
}
