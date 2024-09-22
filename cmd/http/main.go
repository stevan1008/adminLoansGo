package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/stevan1008/adminLoansGo/internal/adapter/handler"
    "github.com/stevan1008/adminLoansGo/internal/adapter/router"
    "github.com/stevan1008/adminLoansGo/internal/core/service"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())

    clientService := service.NewClientService()
    loanService := service.NewLoanService(clientService)
	adminService := service.NewAdminService()
    externalAPIService := service.NewExternalAPIService()

    clientHandler := handler.NewClientHandler(clientService)
    loanHandler := handler.NewLoanHandler(loanService)
    adminHandler := handler.NewAdminHandler(adminService)
    apiHandler := handler.NewAPIHandler(externalAPIService)

    router.SetupRouter(app, clientHandler, loanHandler, adminHandler, apiHandler)

    log.Fatal(app.Listen(":9002"))
}