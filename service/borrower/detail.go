package borrower

import (
	"billing_engine/common"
	borrowerModel "billing_engine/model/borrower"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func (s *service) Detail(reqPath *borrowerModel.DetailReqPath) (resData borrowerModel.DetailResData, statusCode int, err error) {
	borrower, err := s.borrowerRepo.Take([]string{"created_at", "address", "name", "phone_number", "email"}, &borrowerModel.Borrower{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusNotFound, errors.New(common.BorrowerNotFound)
		}
		return resData, http.StatusInternalServerError, err
	}
	resData.CreatedAt = borrower.CreatedAt.Format("02 Jan 2006")
	resData.PhoneNumber = borrower.PhoneNumber
	resData.Email = borrower.Email
	resData.Address = borrower.Address
	resData.Name = borrower.Name
	return
}
