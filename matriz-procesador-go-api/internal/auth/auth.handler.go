package auth

import (
	auth_structs "matriz-procesador-go-api/internal/auth/structs"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service IAuthService
}

func NewAuthHandler(s IAuthService) *AuthHandler {
	return &AuthHandler{service: s}
}


func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req auth_structs.LoginRequest


	if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Request inválida"})
    }
	

	token, err := h.service.Login(req.Usuario, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}