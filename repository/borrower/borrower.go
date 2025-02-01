package borrower

import (
	borrowerModel "billing_engine/model/borrower"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(reqBody *borrowerModel.Borrower) (err error)
		Delete(reqPath *borrowerModel.DeleteReqPath) (err error)
		List(reqQuery *borrowerModel.ListReqQuery) (resData []borrowerModel.ListResData, count int64, err error)
		Take(selectParams []string, conditions *borrowerModel.Borrower) (borrower borrowerModel.Borrower, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
