package billing

import (
	billingScheduleModel "billing_engine/model/billing_schedule"
)

func (r *repo) UpdateBillingSchedule(id *int, values *map[string]any) (err error) {
	return r.db.Model(billingScheduleModel.BillingSchedule{}).Where("id = ?", id).Updates(values).Error
}
