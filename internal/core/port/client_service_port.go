package port

import "github.com/stevan1008/adminLoansGo/internal/core/domain"

type ClientService interface {
    RegisterClient(fullName, email, password string) (domain.Client, error)
	ValidateClientCredentials(email, password string) (domain.Client, error)
	LoginClient(loginReq domain.LoginRequest) (domain.LoginResponse, error)
    GetClientByID(clientID string) (domain.Client, error)
    UpdateClientCreditScore(clientID string, creditScore int) error
	GetAllClients() ([]domain.Client, error) 
}