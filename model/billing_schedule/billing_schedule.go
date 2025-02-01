package billing_schedule

import "time"

type BillingSchedule struct {
	ID         int       `json:"id"`
	LoanID     int       `json:"loan_id"`
	WeekNumber int       `json:"week_number"`
	DueDate    time.Time `json:"due_date"`
	AmountDue  float64   `json:"amount_due"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
