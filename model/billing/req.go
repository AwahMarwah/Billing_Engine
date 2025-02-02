package billing

type (
	CreateReqBody struct {
		LoanId     int     `binding:"required" json:"loan_id"`
		AmountPaid float64 `binding:"required" json:"amount_paid"`
		WeekNumber int     `binding:"required" json:"week_number"`
	}
	DetailReqPath struct {
		LoanId int `binding:"required" uri:"loan_id"`
	}
)
