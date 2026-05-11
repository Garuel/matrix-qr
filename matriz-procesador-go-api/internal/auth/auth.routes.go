package auth

import "github.com/gofiber/fiber/v2"

func MapRoutes(router fiber.Router, handler AuthHandler) {
    router.Post("/login", handler.Login)
}