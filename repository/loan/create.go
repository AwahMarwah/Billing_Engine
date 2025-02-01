package loan

import (
	billingScheduleModel "billing_engine/model/billing_schedule"
	loanModel "billing_engine/model/loan"
	"log"
	"time"
)

func (r *repo) Create(reqBody *loanModel.CreateReqBody, billingSchedule []billingScheduleModel.BillingSchedule) (err error) {
	tx := r.db.Begin()

	loan := loanModel.Loan{
		BorrowerId:         reqBody.BorrowerID,
		PrincipalAmount:    reqBody.PrincipalAmount,
		InterestRate:       reqBody.InterestRate,
		LoanDurationWeeks:  reqBody.LoanDurationWeeks,
		TotalAmountDue:     reqBody.TotalAmountDue,
		OutstandingBalance: reqBody.OutstandingBalance,
		Status:             reqBody.Status,
		StartLoanDate:      time.Now(),
	}

	result := tx.Table("loans").Create(&loan)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	loanID := loan.Id

	for i := range billingSchedule {
		billingSchedule[i].LoanID = loanID
	}

	for _, billing := range billingSchedule {
		if err := tx.Table("billing_schedules").Create(&billing).Error; err != nil {
			tx.Rollback()
			log.Println("Error creating billing schedule:", err)
			return err
		}
	}

	if commitError := tx.Commit().Error; commitError != nil {
		log.Println("Error committing transaction:", commitError)
		return commitError
	}

	return nil
}
