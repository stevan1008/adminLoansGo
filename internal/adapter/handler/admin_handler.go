package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stevan1008/adminLoansGo/internal/core/domain"
	"github.com/stevan1008/adminLoansGo/internal/core/port"
)

type AdminHandler struct {
	adminService port.AdminService
}

func NewAdminHandler(adminService port.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) RegisterAdmin(c *fiber.Ctx) error {
	var request domain.RegisterAdminRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	admin, err := h.adminService.RegisterAdmin(request.FullName, request.Role, request.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(admin)
}

func (h *AdminHandler) LoginAdmin(c *fiber.Ctx) error {
	var loginReq domain.AdminLoginRequest

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	loginResponse, err := h.adminService.LoginAdmin(loginReq)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(loginResponse)
}

func (h *AdminHandler) GetAdminByID(c *fiber.Ctx) error {
	adminID := c.Params("id")

	admin, err := h.adminService.GetAdminByID(adminID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "admin not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(admin)
}
