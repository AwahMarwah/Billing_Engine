package billing

type (
	DetailResData struct {
		LoanId             int     `json:"loan_id"`
		OutstandingBalance float64 `json:"outstanding_balance"`
		Status             string  `json:"status"`
	}
)
