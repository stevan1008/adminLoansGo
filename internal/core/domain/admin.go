package domain

type AdminRole string

// La idea es que dentro de un perfil de admin también se tenfga rangos
const (
	AdminGeneral  AdminRole = "General" 
	AdminManager  AdminRole = "Manager" 
	AdminDirector AdminRole = "Director"
)


type Admin struct {
	ID       string
	FullName string
	Role     AdminRole
}

type RegisterAdminRequest struct {
	FullName string    `json:"fullName"`
	Role     AdminRole `json:"role"`
}
