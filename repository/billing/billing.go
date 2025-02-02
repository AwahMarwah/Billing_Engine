package billing

import (
	billingModel "billing_engine/model/billing"
	billingScheduleModel "billing_engine/model/billing_schedule"
	loanModel "billing_engine/model/loan"
	"gorm.io/gorm"
	"time"
)

type (
	IRepo interface {
		Check(twoWeekAgo time.Time) (resData []loanModel.Loan, err error)
		CheckBillingSchedule(reqBody *billingModel.CreateReqBody) (resData billingScheduleModel.BillingSchedule, err error)
		Create(Payment *billingModel.Repayment) (err error)
		ListBillingSchedule(loanId *int) (resData []billingScheduleModel.BillingSchedule, err error)
		UpdateBillingSchedule(id *int, values *map[string]any) (err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
