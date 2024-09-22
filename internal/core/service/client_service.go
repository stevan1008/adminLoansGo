package service

import (
    "errors"
    "github.com/google/uuid"
    "github.com/stevan1008/adminLoansGo/internal/core/domain"
    "github.com/stevan1008/adminLoansGo/internal/core/port"
)

type ClientServiceImpl struct {
    clients map[string]domain.Client
}

func NewClientService() port.ClientService {
    return &ClientServiceImpl{
        clients: make(map[string]domain.Client),
    }
}

func (s *ClientServiceImpl) RegisterClient(fullName string, email string) (domain.Client, error) {
    client := domain.Client{
        ID:          uuid.New().String(),
        FullName:    fullName,
        Email:       email,
        CreditScore: 0, // Inicia con un puntaje de crédito de 0
    }

    s.clients[client.ID] = client

    return client, nil
}

func (s *ClientServiceImpl) GetClientByID(clientID string) (domain.Client, error) {
    client, exists := s.clients[clientID]
    if !exists {
        return domain.Client{}, errors.New("client not found")
    }

    return client, nil
}

func (s *ClientServiceImpl) UpdateClientCreditScore(clientID string, creditScore int) error {
    client, exists := s.clients[clientID]
    if !exists {
        return errors.New("client not found")
    }

    client.CreditScore = creditScore
    s.clients[clientID] = client

    return nil
}