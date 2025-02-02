package billing

import (
	billingScheduleModel "billing_engine/model/billing_schedule"
)

func (r *repo) ListBillingSchedule(loanId *int) (resData []billingScheduleModel.BillingSchedule, err error) {
	resData = make([]billingScheduleModel.BillingSchedule, 0)
	return resData, r.db.Select("id, loan_id, week_number, due_date, amount_due, status, created_at, updated_at").Model(&billingScheduleModel.BillingSchedule{}).Order("id").Scan(&resData).Error
}
