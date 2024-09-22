package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stevan1008/adminLoansGo/internal/core/domain"
	"github.com/stevan1008/adminLoansGo/internal/core/port"
)

type ClientHandler struct {
    clientService port.ClientService
}

func NewClientHandler(clientService port.ClientService) *ClientHandler {
    return &ClientHandler{
        clientService: clientService,
    }
}

func (h *ClientHandler) RegisterClient(c *fiber.Ctx) error {
    var request domain.RegisterRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request",
        })
    }

    client, err := h.clientService.RegisterClient(request.FullName, request.Email)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(client)
}

func (h *ClientHandler) GetClientByID(c *fiber.Ctx) error {
    clientID := c.Params("id")

    client, err := h.clientService.GetClientByID(clientID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "client not found",
        })
    }

    return c.Status(fiber.StatusOK).JSON(client)
}

func (h *ClientHandler) UpdateCreditScore(c *fiber.Ctx) error {
    var request domain.UpdateCreditScoreRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request",
        })
    }

    clientID := c.Params("id")
    err := h.clientService.UpdateClientCreditScore(clientID, request.CreditScore)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusOK)
}