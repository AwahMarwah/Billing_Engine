package billing

import "time"

type (
	Repayment struct {
		ID            int
		LoanID        int
		PaymentDate   time.Time
		AmountPaid    float64
		PaymentStatus string
		WeekNumber    int
	}
)
