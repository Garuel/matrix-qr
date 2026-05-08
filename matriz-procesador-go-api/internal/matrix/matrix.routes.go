package matrix

import "github.com/gofiber/fiber/v2"

func MapRoutes(router fiber.Router, handler *Handler) {
    router.Post("/process", handler.FactorizeQR)
}