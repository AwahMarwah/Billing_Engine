package borrower

import (
	borrowerModel "billing_engine/model/borrower"
	borrowerRepo "billing_engine/repository/borrower"
)

type (
	IService interface {
		Create(req *borrowerModel.CreateReqBody) (err error)
		Detail(reqPath *borrowerModel.DetailReqPath) (resData borrowerModel.DetailResData, statusCode int, err error)
		Delete(reqPath *borrowerModel.DeleteReqPath) (err error)
		List(reqQuery *borrowerModel.ListReqQuery) (resData []borrowerModel.ListResData, count int64, err error)
	}

	service struct {
		borrowerRepo borrowerRepo.IRepo
	}
)

func NewService(borrowerRepo borrowerRepo.IRepo) *service {
	return &service{borrowerRepo: borrowerRepo}
}
