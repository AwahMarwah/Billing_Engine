package loan

import (
	loanModel "billing_engine/model/loan"
)

func (r *repo) Take(selectParams []string, conditions *loanModel.Loan) (loan loanModel.Loan, err error) {
	return loan, r.db.Select(selectParams).Take(&loan, conditions).Error
}
