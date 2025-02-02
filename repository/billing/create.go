package billing

import billingModel "billing_engine/model/billing"

func (r *repo) Create(payment *billingModel.Repayment) (err error) {
	return r.db.Create(&payment).Error
}
