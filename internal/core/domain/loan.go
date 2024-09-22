package domain

import "time"

type LoanStatus string

const (
    Pending    LoanStatus = "Pending"
    Approved   LoanStatus = "Approved"
    Rejected   LoanStatus = "Rejected"
    Paid       LoanStatus = "Paid"
    Delinquent LoanStatus = "Delinquent" // Préstamo cuando pasa a moroso
)

type Loan struct {
    ID             string
    ClientID       string
    Amount         float64
    InterestRate   float64
    TermInMonths   int
    Status         LoanStatus
    CreatedAt      time.Time
    ApprovedAt     *time.Time
    RejectedAt     *time.Time
    DueDate        time.Time
    RemainingAmount float64   // Monto que le falta al cliente para pagar el prestamo
    IsPaid         bool       // Si el cliente completa el pago del prestamo este será true, de resto false
}

type CreateLoanRequest struct {
    ClientID     string  `json:"clientId"`
    Amount       float64 `json:"amount"`
    TermInMonths int     `json:"termInMonths"`
}