package billing

import (
	billingModel "billing_engine/model/billing"
	loanModel "billing_engine/model/loan"
	billingRepo "billing_engine/repository/billing"
	loanRepo "billing_engine/repository/loan"
)

type (
	IService interface {
		Create(reqBody *billingModel.CreateReqBody) (err error)
		Check() (resData []loanModel.Loan, err error)
		Detail(reqPath *billingModel.DetailReqPath) (resData billingModel.DetailResData, statusCode int, err error)
	}

	service struct {
		billingRepo billingRepo.IRepo
		loanRepo    loanRepo.IRepo
	}
)

func NewService(billingRepo billingRepo.IRepo, loanRepo loanRepo.IRepo) IService {
	return &service{
		billingRepo: billingRepo,
		loanRepo:    loanRepo,
	}
}
