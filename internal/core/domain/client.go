package domain

type Client struct {
	ID          string
	FullName    string
	Email       string
	Password    string
	CreditScore int
}

type RegisterRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type UpdateCreditScoreRequest struct {
	CreditScore int `json:"creditScore"`
}
