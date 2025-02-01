package borrower

import borrowerModel "billing_engine/model/borrower"

func (r *repo) Delete(reqPath *borrowerModel.DeleteReqPath) (err error) {
	return r.db.Model(borrowerModel.Borrower{}).Delete(&borrowerModel.Borrower{Id: reqPath.Id}).Error
}
