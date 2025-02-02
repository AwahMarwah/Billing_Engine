package billing

import (
	"billing_engine/common"
	billingModel "billing_engine/model/billing"
	loanModel "billing_engine/model/loan"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (s *service) Create(reqBody *billingModel.CreateReqBody) (err error) {
	loan, err := s.loanRepo.Take([]string{"id", "outstanding_balance"}, &loanModel.Loan{Id: reqBody.LoanId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(common.LoanNotFound)
		}
		return err
	}
	schedule, err := s.billingRepo.CheckBillingSchedule(reqBody)
	if err != nil {
		return err
	}

	paymentStatus := "Paid"
	if reqBody.AmountPaid < schedule.AmountDue {
		paymentStatus = "Missed"
	}

	payment := billingModel.Repayment{
		LoanID:        reqBody.LoanId,
		PaymentDate:   time.Now(),
		AmountPaid:    reqBody.AmountPaid,
		PaymentStatus: paymentStatus,
		WeekNumber:    reqBody.WeekNumber,
	}

	if err := s.billingRepo.Create(&payment); err != nil {
		return err
	}

	updateValues := map[string]any{
		"status":     paymentStatus,
		"updated_at": time.Now(),
	}

	if err := s.billingRepo.UpdateBillingSchedule(&schedule.ID, &updateValues); err != nil {
		return err
	}

	// Update outstanding balance
	newOutstandingBalance := loan.OutstandingBalance - reqBody.AmountPaid
	if newOutstandingBalance <= 0 {
		updateLoanValue := map[string]any{
			"outstanding_balance": 0,
			"status":              "Closed",
			"updated_at":          time.Now(),
		}
		if err := s.loanRepo.Update(&loan.Id, &updateLoanValue); err != nil {
			return err
		}
	} else {
		loanValues := map[string]any{
			"outstanding_balance": newOutstandingBalance,
			"updated_at":          time.Now(),
		}
		if err := s.loanRepo.Update(&loan.Id, &loanValues); err != nil {
			return err
		}
	}
	return err
}
