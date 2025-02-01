package borrower

import borrowerModel "billing_engine/model/borrower"

func (s *service) List(reqQuery *borrowerModel.ListReqQuery) (resData []borrowerModel.ListResData, count int64, err error) {
	return s.borrowerRepo.List(reqQuery)
}
