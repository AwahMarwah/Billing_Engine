package borrower

import (
	"billing_engine/common"
	borrowerModel "billing_engine/model/borrower"
	"errors"
	"gorm.io/gorm"
)

func (s *service) Delete(reqPath *borrowerModel.DeleteReqPath) (err error) {
	_, err = s.borrowerRepo.Take([]string{"id"}, &borrowerModel.Borrower{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(common.BorrowerNotFound)
		}
		return err
	}
	return s.borrowerRepo.Delete(reqPath)
}
