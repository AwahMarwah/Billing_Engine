package loan

import (
	loanModel "billing_engine/model/loan"
)

func (r *repo) List(reqQuery *loanModel.ListReqQuery) (resData []loanModel.ListResData, count int64, err error) {
	resData = make([]loanModel.ListResData, 0)
	return resData, count, r.db.Select("created_at, id, borrower_id, principal_amount, interest_rate, loan_duration_weeks, total_amount_due, outstanding_balance, status, is_delinquent, start_loan_date, updated_at").Model(loanModel.Loan{}).Count(&count).Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}
