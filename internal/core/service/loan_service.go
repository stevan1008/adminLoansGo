package service

import (
    "errors"
    "time"

    "github.com/google/uuid"
    "github.com/stevan1008/adminLoansGo/internal/core/domain"
    "github.com/stevan1008/adminLoansGo/internal/core/port"
)

type LoanServiceImpl struct {
    clientService port.ClientService
    loans         map[string]domain.Loan
}

func NewLoanService(clientService port.ClientService) port.LoanService {
    return &LoanServiceImpl{
        clientService: clientService,
        loans:         make(map[string]domain.Loan),
    }
}

func (s *LoanServiceImpl) CreateLoan(loan domain.Loan) (domain.Loan, error) {
    client, err := s.clientService.GetClientByID(loan.ClientID)
    if err != nil {
        return domain.Loan{}, err
    }

    // Validar el crédito del cliente, default en 600 para pruebas
    if client.CreditScore < 600 {
        return domain.Loan{}, errors.New("client has insufficient credit score")
    }

    loan.InterestRate = s.calculateInterestRate(client.CreditScore)
    loan.ID = uuid.New().String()
    loan.CreatedAt = time.Now()
    loan.DueDate = loan.CreatedAt.AddDate(0, loan.TermInMonths, 0)
    loan.Status = domain.Pending // Cuando se crea un prestamo la idea es que el estado inicial sea pending para que desde el frontend se apruebe o rechace 
    loan.RemainingAmount = loan.Amount

    s.loans[loan.ID] = loan

    return loan, nil
}

func (s *LoanServiceImpl) ApproveLoan(loanID string, adminID string) error {
    loan, exists := s.loans[loanID]
    if !exists {
        return errors.New("loan not found")
    }

    if loan.Status != domain.Pending {
        return errors.New("loan is not pending and cannot be approved")
    }

    now := time.Now()
    loan.Status = domain.Approved
    loan.ApprovedAt = &now

    s.loans[loan.ID] = loan

    return nil
}

func (s *LoanServiceImpl) RejectLoan(loanID string, adminID string) error {
    loan, exists := s.loans[loanID]
    if !exists {
        return errors.New("loan not found")
    }

    if loan.Status != domain.Pending {
        return errors.New("loan is not pending and cannot be rejected")
    }

    now := time.Now()
    loan.Status = domain.Rejected
    loan.RejectedAt = &now

    s.loans[loan.ID] = loan

    return nil
}

func (s *LoanServiceImpl) GetLoanByID(loanID string) (domain.Loan, error) {
    loan, exists := s.loans[loanID]
    if !exists {
        return domain.Loan{}, errors.New("loan not found")
    }
    return loan, nil
}

func (s *LoanServiceImpl) ListLoansByClientID(clientID string) ([]domain.Loan, error) {
    var clientLoans []domain.Loan
    for _, loan := range s.loans {
        if loan.ClientID == clientID {
            clientLoans = append(clientLoans, loan)
        }
    }
    if len(clientLoans) == 0 {
        return nil, errors.New("no loans found for this client")
    }
    return clientLoans, nil
}

func (s *LoanServiceImpl) ListLoansHistory() ([]domain.Loan, error) {
    var allLoans []domain.Loan
    for _, loan := range s.loans {
        allLoans = append(allLoans, loan)
    }
    if len(allLoans) == 0 {
        return nil, errors.New("no loans found")
    }
    return allLoans, nil
}

func (s *LoanServiceImpl) calculateInterestRate(creditScore int) float64 {
    if creditScore >= 750 {
        return 3.5
    } else if creditScore >= 650 {
        return 5.0
    }
    return 7.5
}

func (s *LoanServiceImpl) RegisterPayment(loanID string, amount float64) (domain.Payment, error) {
    loan, exists := s.loans[loanID]
    if !exists {
        return domain.Payment{}, errors.New("loan not found")
    }

    // Valida que el monto no exceda el valor que falta por pagar del orestamo
    if amount > loan.RemainingAmount {
        return domain.Payment{}, errors.New("payment exceeds loan balance")
    }

    loan.RemainingAmount -= amount

    payment := domain.Payment{
        ID:     uuid.New().String(),
        LoanID: loanID,
        Amount: amount,
        Date:   time.Now(),
        Status: "completed",
    }

    if loan.RemainingAmount == 0 {
        loan.Status = domain.Paid
        loan.IsPaid = true
    } else {
        loan.IsPaid = false
    }

    s.loans[loan.ID] = loan

    return payment, nil
}

func (s *LoanServiceImpl) MarkLoanAsDelinquent(loanID string) error {
    loan, exists := s.loans[loanID]
    if !exists {
        return errors.New("loan not found")
    }

    if loan.Status == domain.Pending && time.Now().After(loan.DueDate) {
        loan.Status = domain.Delinquent
        s.loans[loanID] = loan
        return nil
    }

    return errors.New("loan is either not pending or has not passed the due date")
}

func (s *LoanServiceImpl) MarkAllLoansAsDelinquent() error {
    for _, loan := range s.loans {
        if loan.Status == domain.Pending && time.Now().After(loan.DueDate) {
            loan.Status = domain.Delinquent
            s.loans[loan.ID] = loan
        }
    }
    return nil
}

func (s *LoanServiceImpl) GetActiveLoan(clientID string) (domain.Loan, error) {
    for _, loan := range s.loans {
        if loan.ClientID == clientID && (loan.Status == domain.Pending || loan.Status == domain.Approved) {
            return loan, nil
        }
    }
    return domain.Loan{}, errors.New("no active loan found for this client")
}