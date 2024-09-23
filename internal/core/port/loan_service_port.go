package port

import "github.com/stevan1008/adminLoansGo/internal/core/domain"

type LoanService interface {
    CreateLoan(loan domain.Loan) (domain.Loan, error)
    ApproveLoan(loanID string, adminID string) error
    RejectLoan(loanID string, adminID string) error
    GetLoanByID(loanID string) (domain.Loan, error)
    ListLoansByClientID(clientID string) ([]domain.Loan, error)
    ListLoansHistory() ([]domain.Loan, error)
    RegisterPayment(loanID string, amount float64) (domain.Payment, error)
    MarkLoanAsDelinquent(loanID string) error
    MarkAllLoansAsDelinquent() error
	GetActiveLoan(clientID string) (domain.Loan, error)
}
