package billing

import (
	"billing_engine/common"
	billingModel "billing_engine/model/billing"
	loanModel "billing_engine/model/loan"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func (s *service) Detail(reqPath *billingModel.DetailReqPath) (resData billingModel.DetailResData, statusCode int, err error) {
	loan, err := s.loanRepo.Take([]string{"id", "outstanding_balance", "status"}, &loanModel.Loan{Id: reqPath.LoanId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusNotFound, errors.New(common.LoanNotFound)
		}
		return resData, http.StatusInternalServerError, err
	}
	resData.LoanId = loan.Id
	resData.OutstandingBalance = loan.OutstandingBalance
	resData.Status = loan.Status
	return
}
