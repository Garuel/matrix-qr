package matrix

import (
	"encoding/json"
	"log"
	matrizdomain "matriz-procesador-go-api/internal/domain/models"
	clients "matriz-procesador-go-api/internal/infrastructure/clients"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
    service    Service
    nodeClient clients.NodeClientInterface
}

func NewHandler(s Service, nc clients.NodeClientInterface) *Handler {
    return &Handler{service: s, nodeClient: nc}
}

func (h *Handler) FactorizeQR(c *fiber.Ctx) error {
    var req matrizdomain.MatrixRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Matriz inválida"})
    }

	log.Printf("Matriz recibida")
    qrResult, err := h.service.FactorizeQR(req.Matrix)
    if err != nil {
        log.Printf("Error factorizando QR: %v", err)
        return c.Status(500).JSON(fiber.Map{"error": "Factorización QR fallida"})
    }

    log.Printf("Factorización QR exitosa")

	log.Printf("Enviando matriz a la API de Node...")
    resp, err := h.nodeClient.SendToStats(qrResult)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "No se pudo conectar con la API de Node"})
    }
    
    defer resp.Body.Close()

    var nodeResponse interface{}
    if err := json.NewDecoder(resp.Body).Decode(&nodeResponse); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Error al decodificar respuesta de Node"})
    }

	log.Printf("Matriz enviada exitosamente")


    return c.Status(200).JSON(nodeResponse) 
}