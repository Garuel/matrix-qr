package bootstrap

import (
	Config "matriz-procesador-go-api/config"
	"matriz-procesador-go-api/internal/auth"
	auth_middlewares "matriz-procesador-go-api/internal/auth/middleware"
	"matriz-procesador-go-api/internal/infrastructure/clients"
	"matriz-procesador-go-api/internal/matrix"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initialize(cfg *Config.ConfigStruct) *fiber.App {
    app := fiber.New()

    // infrastructure y modulos
    nodeClient := clients.NewNodeClient(cfg.NODE_API_URL)

    matrixService := matrix.NewService()
    authService := auth.NewAuthService(cfg.JWT_SECRET)


    authHandler := auth.NewAuthHandler(authService)
    matrixHandler := matrix.NewHandler(matrixService, nodeClient)



    app.Use(cors.New(cors.Config{
    AllowOrigins: "*", // Para conectar al frontend, en produccion seria el dominio del frontend
    AllowHeaders: "Origin, Content-Type, Accept",
}))

    api := app.Group("/api")

    authGroup := api.Group("/auth")
    auth.MapRoutes(authGroup, *authHandler)


    matrixGroup := api.Group("/matrix", auth_middlewares.JWTMiddleware(cfg.JWT_SECRET))
    matrix.MapRoutes(matrixGroup, matrixHandler)

    return app
}