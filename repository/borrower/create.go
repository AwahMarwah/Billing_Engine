package borrower

import borrowerModel "billing_engine/model/borrower"

func (r *repo) Create(borrower *borrowerModel.Borrower) (err error) {
	return r.db.Create(&borrower).Error
}
