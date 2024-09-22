package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/stevan1008/adminLoansGo/internal/core/port"
)

type APIHandler struct {
    externalAPIService port.ExternalAPIService
}

func NewAPIHandler(externalAPIService port.ExternalAPIService) *APIHandler {
    return &APIHandler{
        externalAPIService: externalAPIService,
    }
}

func (h *APIHandler) ValidateClientDocuments(c *fiber.Ctx) error {
    clientID := c.Params("id")

    isValid, err := h.externalAPIService.ValidateClientDocuments(clientID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "error validating documents",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "isValid": isValid,
    })
}

func (h *APIHandler) GetCreditScore(c *fiber.Ctx) error {
    clientID := c.Params("id")

    creditScore, err := h.externalAPIService.GetCreditScore(clientID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "error fetching credit score",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "creditScore": creditScore,
    })
}