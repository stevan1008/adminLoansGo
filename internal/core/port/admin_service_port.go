package port

import "github.com/stevan1008/adminLoansGo/internal/core/domain"

type AdminService interface {
    RegisterAdmin(fullName string, role domain.AdminRole, password string) (domain.Admin, error)
	LoginAdmin(loginReq domain.AdminLoginRequest) (domain.AdminLoginResponse, error)
    GetAdminByID(adminID string) (domain.Admin, error)
}