package loan

import (
	"billing_engine/common"
	loanModel "billing_engine/model/loan"
	"errors"
	"gorm.io/gorm"
)

func (s *service) Delete(reqPath *loanModel.DeleteReqPath) (err error) {
	_, err = s.loanRepo.Take([]string{"id"}, &loanModel.Loan{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(common.LoanNotFound)
		}
		return err
	}
	return s.loanRepo.Delete(reqPath)
}
