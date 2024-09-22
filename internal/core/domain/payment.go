package domain

import "time"

type Payment struct {
    ID         string
    LoanID     string
    Amount     float64
    Date       time.Time
    Status     string
}

type PaymentRequest struct {
    LoanID string  `json:"loan_id"`
    Amount float64 `json:"amount"`
}