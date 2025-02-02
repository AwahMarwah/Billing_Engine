package loan

import (
	loanModel "billing_engine/model/loan"
	billingRepo "billing_engine/repository/billing"
	borrowerRepo "billing_engine/repository/borrower"
	"billing_engine/repository/loan"
)

type (
	IService interface {
		Create(reqBody *loanModel.CreateReqBody) (statusCode int, err error)
		Detail(reqPath *loanModel.DetailReqPath) (resData loanModel.DetailResData, statusCode int, err error)
		Delete(reqPath *loanModel.DeleteReqPath) (err error)
		List(reqQuery *loanModel.ListReqQuery) (resData []loanModel.ListResData, count int64, err error)
	}

	service struct {
		loanRepo     loan.IRepo
		borrowerRepo borrowerRepo.IRepo
		billingRepo  billingRepo.IRepo
	}
)

func NewService(loanRepo loan.IRepo, borrowerRepo borrowerRepo.IRepo, billingRepo billingRepo.IRepo) IService {
	return &service{
		loanRepo:     loanRepo,
		borrowerRepo: borrowerRepo,
		billingRepo:  billingRepo,
	}
}
