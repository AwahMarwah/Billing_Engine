package loan

import (
	borrowerRepo "billing_engine/repository/borrower"
	loanRepo "billing_engine/repository/loan"
	"billing_engine/service/loan"
	"gorm.io/gorm"
)

type controller struct {
	loanService loan.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{loanService: loan.NewService(loanRepo.NewRepo(db), borrowerRepo.NewRepo(db))}
}
