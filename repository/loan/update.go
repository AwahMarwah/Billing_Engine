package loan

import loanModel "billing_engine/model/loan"

func (r *repo) Update(id *int, values *map[string]any) (err error) {
	return r.db.Model(loanModel.Loan{}).Where("id = ?", id).Updates(values).Error
}
