package borrower

import borrowerModel "billing_engine/model/borrower"

func (r *repo) List(reqQuery *borrowerModel.ListReqQuery) (resData []borrowerModel.ListResData, count int64, err error) {
	resData = make([]borrowerModel.ListResData, 0)
	return resData, count, r.db.Select("TO_CHAR(created_at, 'DD Mon YYYY') AS created_at, id, name, email, phone_number, address").Model(borrowerModel.Borrower{}).Count(&count).Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}
