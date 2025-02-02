package loan

import (
	billingModel "billing_engine/model/billing_schedule"
	loanModel "billing_engine/model/loan"
)

func (r *repo) Delete(reqPath *loanModel.DeleteReqPath) (err error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err = tx.Where("loan_id = ?", reqPath.Id).Delete(&billingModel.BillingSchedule{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Model(loanModel.Loan{}).Delete(&loanModel.Loan{Id: reqPath.Id}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
