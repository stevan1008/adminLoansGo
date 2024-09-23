package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stevan1008/adminLoansGo/internal/core/domain"
	"github.com/stevan1008/adminLoansGo/internal/core/port"
	"golang.org/x/crypto/bcrypt"
)

type ClientServiceImpl struct {
	clients    map[string]domain.Client
	secretKey  string
}

func NewClientService(secretKey string) port.ClientService {
	return &ClientServiceImpl{
		clients:   make(map[string]domain.Client),
		secretKey: secretKey,
	}
}

func (s *ClientServiceImpl) RegisterClient(fullName, email, password string) (domain.Client, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Client{}, err
	}

	client := domain.Client{
		ID:          uuid.New().String(),
		FullName:    fullName,
		Email:       email,
		Password:    string(hashedPassword),
		CreditScore: 0,
	}

	s.clients[client.ID] = client

	return client, nil
}

func (s *ClientServiceImpl) ValidateClientCredentials(email, password string) (domain.Client, error) {
	for _, client := range s.clients {
		if client.Email == email {
			err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
			if err != nil {
				return domain.Client{}, errors.New("invalid credentials")
			}
			return client, nil
		}
	}
	return domain.Client{}, errors.New("client not found")
}

func (s *ClientServiceImpl) LoginClient(loginReq domain.LoginRequest) (domain.LoginResponse, error) {
	client, exists := s.clients[loginReq.ID]
	if !exists {
		return domain.LoginResponse{}, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(loginReq.Password))
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  client.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		ID:    client.ID,
		Token: tokenString,
	}, nil
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

func (s *ClientServiceImpl) GetAllClients() ([]domain.Client, error) {
    var clientList []domain.Client

    for _, client := range s.clients {
        clientList = append(clientList, client)
    }

    return clientList, nil
}