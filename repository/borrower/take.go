package borrower

import borrowerModel "billing_engine/model/borrower"

func (r *repo) Take(selectParams []string, conditions *borrowerModel.Borrower) (borrower borrowerModel.Borrower, err error) {
	return borrower, r.db.Select(selectParams).Take(&borrower, conditions).Error
}
