package loan

import "time"

type Loan struct {
	Id                 int
	BorrowerId         int
	PrincipalAmount    float64
	InterestRate       float64
	LoanDurationWeeks  int
	TotalAmountDue     float64
	OutstandingBalance float64
	Status             string
	IsDelinquent       bool
	StartLoanDate      time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
