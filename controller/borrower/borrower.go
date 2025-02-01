package borrower

import (
	borrowerRepo "billing_engine/repository/borrower"
	"billing_engine/service/borrower"
	"gorm.io/gorm"
)

type controller struct {
	borrowerService borrower.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{borrowerService: borrower.NewService(borrowerRepo.NewRepo(db))}
}
