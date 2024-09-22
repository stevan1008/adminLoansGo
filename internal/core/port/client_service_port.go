package port

import "github.com/stevan1008/adminLoansGo/internal/core/domain"

type ClientService interface {
    RegisterClient(fullName string, email string) (domain.Client, error)
    GetClientByID(clientID string) (domain.Client, error)
    UpdateClientCreditScore(clientID string, creditScore int) error
}