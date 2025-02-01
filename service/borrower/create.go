package borrower

import borrowerModel "billing_engine/model/borrower"

func (s *service) Create(req *borrowerModel.CreateReqBody) (err error) {
	return s.borrowerRepo.Create(&borrowerModel.Borrower{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	})
}
