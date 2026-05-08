package bootstrap

import (
	Config "matriz-procesador-go-api/config"
	"matriz-procesador-go-api/internal/infrastructure/clients"
	"matriz-procesador-go-api/internal/matrix"

	"github.com/gofiber/fiber/v2"
)

func Initialize(cfg *Config.ConfigStruct) *fiber.App {
    app := fiber.New()

    // infrastructure y modulos
    nodeClient := clients.NewNodeClient(cfg.NODE_API_URL)
    matrixService := matrix.NewService()
    matrixHandler := matrix.NewHandler(matrixService, nodeClient)

    // rutas
    api := app.Group("/api/matrix")
    matrix.MapRoutes(api, matrixHandler)

    return app
}