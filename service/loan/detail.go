package loan

import (
	"billing_engine/common"
	loanModel "billing_engine/model/loan"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func (s *service) Detail(reqPath *loanModel.DetailReqPath) (resData loanModel.DetailResData, statusCode int, err error) {
	loan, err := s.loanRepo.Take([]string{"id", "borrower_id", "principal_amount", "interest_rate", "loan_duration_weeks", "total_amount_due", "outstanding_balance", "status", "is_delinquent", "start_loan_date", "created_at"}, &loanModel.Loan{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusNotFound, errors.New(common.LoanNotFound)
		}
		return resData, http.StatusInternalServerError, err
	}
	resData.ID = loan.Id
	resData.BorrowerID = loan.BorrowerId
	resData.PrincipalAmount = loan.PrincipalAmount
	resData.InterestRate = loan.InterestRate
	resData.LoanDurationWeeks = loan.LoanDurationWeeks
	resData.TotalAmountDue = loan.TotalAmountDue
	resData.OutstandingBalance = loan.OutstandingBalance
	resData.Status = loan.Status
	resData.IsDelinquent = loan.IsDelinquent
	resData.StartLoanDate = loan.StartLoanDate.Format("02 Jan 2006")
	resData.CreatedAt = loan.CreatedAt.Format("02 Jan 2006")
	return
}
