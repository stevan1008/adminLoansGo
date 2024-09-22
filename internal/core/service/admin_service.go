package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/stevan1008/adminLoansGo/internal/core/domain"
	"github.com/stevan1008/adminLoansGo/internal/core/port"
)

type AdminServiceImpl struct {
    admins map[string]domain.Admin
}

func NewAdminService() port.AdminService {
    return &AdminServiceImpl{
        admins: make(map[string]domain.Admin),
    }
}

func (s *AdminServiceImpl) RegisterAdmin(fullName string, role domain.AdminRole) (domain.Admin, error) {
    admin := domain.Admin{
        ID:       uuid.New().String(),
        FullName: fullName,
        Role:     role,
    }

    s.admins[admin.ID] = admin

    return admin, nil
}

func (s *AdminServiceImpl) GetAdminByID(adminID string) (domain.Admin, error) {
    admin, exists := s.admins[adminID]
    if !exists {
        return domain.Admin{}, errors.New("admin not found")
    }

    return admin, nil
}