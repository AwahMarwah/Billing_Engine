package billing

import (
	billingRepo "billing_engine/repository/billing"
	loanRepo "billing_engine/repository/loan"
	"billing_engine/service/billing"
	"gorm.io/gorm"
)

type controller struct {
	billingService billing.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{billingService: billing.NewService(billingRepo.NewRepo(db), loanRepo.NewRepo(db))}
}
