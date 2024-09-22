package port

import "github.com/stevan1008/adminLoansGo/internal/core/domain"

type AdminService interface {
    RegisterAdmin(fullName string, role domain.AdminRole) (domain.Admin, error)
    GetAdminByID(adminID string) (domain.Admin, error)
}