package domain

type Client struct {
	ID          string
	FullName    string
	Email       string
	CreditScore int
}

type RegisterRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type UpdateCreditScoreRequest struct {
	CreditScore int `json:"creditScore"`
}