package loan

import loanModel "billing_engine/model/loan"

func (s *service) List(reqQuery *loanModel.ListReqQuery) (resData []loanModel.ListResData, count int64, err error) {
	return s.loanRepo.List(reqQuery)
}
