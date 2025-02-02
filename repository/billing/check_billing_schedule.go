package billing

import (
	billingModel "billing_engine/model/billing"
	billingScheduleModel "billing_engine/model/billing_schedule"
)

func (r *repo) CheckBillingSchedule(reqBody *billingModel.CreateReqBody) (resData billingScheduleModel.BillingSchedule, err error) {
	return resData, r.db.Model(&billingScheduleModel.BillingSchedule{}).Where("loan_id = ? AND week_number = ?", reqBody.LoanId, reqBody.WeekNumber).First(&resData).Error
}
