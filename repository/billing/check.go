package billing

import (
	delinquencyModel "billing_engine/model/delinquency"
	loanModel "billing_engine/model/loan"
	"time"
)

func (r *repo) Check(twoWeekAgo time.Time) (resData []loanModel.Loan, err error) {
	var overdueLoans []loanModel.Loan
	err = r.db.Model(&loanModel.Loan{}).
		Joins("JOIN billing_schedules b ON b.loan_id = loans.id").
		Where("b.due_date <= ? AND b.status = ? AND loans.is_delinquent = ?", twoWeekAgo, "Pending", false).
		Distinct().
		Find(&overdueLoans).Error
	if err != nil {
		return nil, err
	}

	if len(overdueLoans) == 0 {
		return overdueLoans, nil
	}

	for _, loan := range overdueLoans {
		if err := r.db.Model(&loan).Update("is_delinquent", true).Error; err != nil {
			return nil, err
		}

		delinquency := delinquencyModel.Deliquency{
			LoanId:      loan.Id,
			Delinquent:  true,
			DateChecked: time.Now(),
		}
		if err := r.db.Create(&delinquency).Error; err != nil {
			return nil, err
		}
	}

	return overdueLoans, nil
}
