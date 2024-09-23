package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/stevan1008/adminLoansGo/internal/core/domain"
    "github.com/stevan1008/adminLoansGo/internal/core/port"
)

type LoanHandler struct {
    loanService port.LoanService
}

func NewLoanHandler(loanService port.LoanService) *LoanHandler {
    return &LoanHandler{
        loanService: loanService,
    }
}

func (h *LoanHandler) CreateLoan(c *fiber.Ctx) error {
    var request domain.CreateLoanRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request",
        })
    }

    loan := domain.Loan{
        ClientID:     request.ClientID,
        Amount:       request.Amount,
        TermInMonths: request.TermInMonths,
        Status:       domain.Pending,
    }

    createdLoan, err := h.loanService.CreateLoan(loan)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdLoan)
}

func (h *LoanHandler) ApproveLoan(c *fiber.Ctx) error {
    loanID := c.Params("id")
    adminID := c.Query("adminId")

    if err := h.loanService.ApproveLoan(loanID, adminID); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusOK)
}

func (h *LoanHandler) RejectLoan(c *fiber.Ctx) error {
    loanID := c.Params("id")
    adminID := c.Query("adminId")

    if err := h.loanService.RejectLoan(loanID, adminID); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusOK)
}

func (h *LoanHandler) ListLoansHistory(c *fiber.Ctx) error {
    loans, err := h.loanService.ListLoansHistory()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(loans)
}

func (h *LoanHandler) RegisterPayment(c *fiber.Ctx) error {
    var request domain.PaymentRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request",
        })
    }

    payment, err := h.loanService.RegisterPayment(request.LoanID, request.Amount)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(payment)
}

func (h *LoanHandler) MarkLoanAsDelinquent(c *fiber.Ctx) error {
    loanID := c.Params("id")

    if err := h.loanService.MarkLoanAsDelinquent(loanID); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusOK)
}

func (h *LoanHandler) MarkAllLoansAsDelinquent(c *fiber.Ctx) error {
    if err := h.loanService.MarkAllLoansAsDelinquent(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusOK)
}

func (h *LoanHandler) GetActiveLoan(c *fiber.Ctx) error {
    clientID := c.Query("clientId")
    if clientID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "clientId is required",
        })
    }

    loan, err := h.loanService.GetActiveLoan(clientID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(loan)
}