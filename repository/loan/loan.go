package loan

import (
	billingScheduleModel "billing_engine/model/billing_schedule"
	loanModel "billing_engine/model/loan"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(reqBody *loanModel.CreateReqBody, billingSchedule []billingScheduleModel.BillingSchedule) (err error)
		Delete(reqPath *loanModel.DeleteReqPath) (err error)
		List(reqQuery *loanModel.ListReqQuery) (resData []loanModel.ListResData, count int64, err error)
		Take(selectParams []string, conditions *loanModel.Loan) (loan loanModel.Loan, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
