package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/stevan1008/adminLoansGo/internal/adapter/handler"
)

func SetupRouter(app *fiber.App, clientHandler *handler.ClientHandler, loanHandler *handler.LoanHandler, adminHandler *handler.AdminHandler, apiHandler *handler.APIHandler) {
    // Rutas para clientes
    app.Post("/clients", clientHandler.RegisterClient)
	app.Get("/clients", clientHandler.GetAllClients)
	app.Post("/clients/login", clientHandler.LoginClient)
    app.Get("/clients/:id", clientHandler.GetClientByID)
	app.Patch("/clients/:id/credit-score", clientHandler.UpdateCreditScore)

    // Rutas para préstamos | loans
    app.Post("/loans", loanHandler.CreateLoan)
	app.Get("/loans/history", loanHandler.ListLoansHistory)
    app.Patch("/loans/:id/approve", loanHandler.ApproveLoan)
    app.Patch("/loans/:id/reject", loanHandler.RejectLoan)
	app.Post("/loans/payment", loanHandler.RegisterPayment)
	app.Patch("/loans/:id/delinquent", loanHandler.MarkLoanAsDelinquent)
	app.Patch("/loans/delinquent/all", loanHandler.MarkAllLoansAsDelinquent)
	app.Get("/loans/active", loanHandler.GetActiveLoan)

    // Rutas para administradores
	app.Post("/admins", adminHandler.RegisterAdmin)
	app.Post("/admins/login", adminHandler.LoginAdmin)
	app.Get("/admins/:id", adminHandler.GetAdminByID)

    // Rutas para APIs de proveedores externos
    app.Get("/api/clients/:id/validate-documents", apiHandler.ValidateClientDocuments)
    app.Get("/api/clients/:id/credit-score", apiHandler.GetCreditScore)
}