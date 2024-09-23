package domain

type AdminRole string

const (
	AdminGeneral  AdminRole = "General"
	AdminManager  AdminRole = "Manager"
	AdminDirector AdminRole = "Director"
)

type Admin struct {
	ID       string
	FullName string
	Role     AdminRole
	Password string
}

type AdminLoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type RegisterAdminRequest struct {
	FullName string    `json:"fullName"`
	Role     AdminRole `json:"role"`
	Password string    `json:"password"`
}