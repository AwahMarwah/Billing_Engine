package billing

import (
	loanModel "billing_engine/model/loan"
	"time"
)

func (s *service) Check() (resData []loanModel.Loan, err error) {
	twoWeeksAgo := time.Now().AddDate(0, 0, -14)
	return s.billingRepo.Check(twoWeeksAgo)
}
