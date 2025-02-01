package loan

type (
	CreateReqBody struct {
		BorrowerID         int     `json:"borrower_id"`
		PrincipalAmount    float64 `json:"principal_amount"`
		InterestRate       float64 `json:"interest_rate"`
		LoanDurationWeeks  int     `json:"loan_duration_weeks"`
		TotalAmountDue     float64 `json:"total_amount_due"`
		OutstandingBalance float64 `json:"outstanding_balance"`
		Status             string  `json:"status"`
		StartLoanDate      string  `json:"start_loan_date"`
	}

	DeleteReqPath struct {
		Id int `binding:"required" uri:"id"`
	}

	DetailReqPath struct {
		Id int `binding:"required" uri:"id"`
	}

	ListReqQuery struct {
		Limit  int `form:"limit"`
		Offset int
		Page   int `form:"page"`
	}
)
