package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stevan1008/adminLoansGo/internal/core/domain"
	"github.com/stevan1008/adminLoansGo/internal/core/port"
)

type AdminServiceImpl struct {
	admins    map[string]domain.Admin
	secretKey string
}

func NewAdminService(secretKey string) port.AdminService {
	return &AdminServiceImpl{
		admins:    make(map[string]domain.Admin),
		secretKey: secretKey,
	}
}

func (s *AdminServiceImpl) RegisterAdmin(fullName string, role domain.AdminRole, password string) (domain.Admin, error) {
	admin := domain.Admin{
		ID:       uuid.New().String(),
		FullName: fullName,
		Role:     role,
		Password: password,
	}

	s.admins[admin.ID] = admin

	return admin, nil
}

func (s *AdminServiceImpl) LoginAdmin(loginReq domain.AdminLoginRequest) (domain.AdminLoginResponse, error) {
	admin, exists := s.admins[loginReq.ID]
	if !exists || admin.Password != loginReq.Password {
		return domain.AdminLoginResponse{}, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  admin.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return domain.AdminLoginResponse{}, err
	}

	return domain.AdminLoginResponse{
		ID:    admin.ID,
		Token: tokenString,
	}, nil
}

func (s *AdminServiceImpl) GetAdminByID(adminID string) (domain.Admin, error) {
    admin, exists := s.admins[adminID]
    if !exists {
        return domain.Admin{}, errors.New("admin not found")
    }

    return admin, nil
}