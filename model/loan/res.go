package loan

import (
	billingScheduleModel "billing_engine/model/billing_schedule"
	"time"
)

type (
	DetailResData struct {
		ID                 int                                    `json:"id"`
		BorrowerID         int                                    `json:"borrower_id"`
		PrincipalAmount    float64                                `json:"principal_amount"`
		InterestRate       float64                                `json:"interest_rate"`
		LoanDurationWeeks  int                                    `json:"loan_duration_weeks"`
		TotalAmountDue     float64                                `json:"total_amount_due"`
		OutstandingBalance float64                                `json:"outstanding_balance"`
		Status             string                                 `json:"status"`
		IsDelinquent       bool                                   `json:"is_delinquent"`
		StartLoanDate      string                                 `json:"start_loan_date"`
		CreatedAt          string                                 `json:"created_at"`
		Schedule           []billingScheduleModel.BillingSchedule `json:"schedule"`
	}
	ListResData struct {
		ID                 int       `json:"id"`
		BorrowerID         int       `json:"borrower_id"`
		PrincipalAmount    float64   `json:"principal_amount"`
		InterestRate       float64   `json:"interest_rate"`
		LoanDurationWeeks  int       `json:"loan_duration_weeks"`
		TotalAmountDue     float64   `json:"total_amount_due"`
		OutstandingBalance float64   `json:"outstanding_balance"`
		Status             string    `json:"status"`
		IsDelinquent       bool      `json:"is_delinquent"`
		StartLoanDate      time.Time `json:"start_loan_date"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	}
)
