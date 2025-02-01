package loan

import (
	billingScheduleModel "billing_engine/model/billing_schedule"
	borrowerModel "billing_engine/model/borrower"
	loanModel "billing_engine/model/loan"
	"fmt"
	"net/http"
	"time"
)

func (s *service) Create(reqBody *loanModel.CreateReqBody) (statusCode int, err error) {
	var billingSchedule []billingScheduleModel.BillingSchedule

	borrower, err := s.borrowerRepo.Take([]string{"id"}, &borrowerModel.Borrower{Id: reqBody.BorrowerID})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if borrower.Id == 0 {
		return http.StatusNotFound, fmt.Errorf("borrower not found")
	}

	// Validation for loan with borrower id
	loan, err := s.loanRepo.Take([]string{"borrower_id"}, &loanModel.Loan{BorrowerId: reqBody.BorrowerID})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if loan.BorrowerId == borrower.Id {
		return http.StatusBadRequest, fmt.Errorf("borrower id: %d already take loans", reqBody.BorrowerID)
	}
	now := time.Now().Format("02 Jan 2006")
	reqBody.TotalAmountDue = reqBody.PrincipalAmount + (reqBody.PrincipalAmount * (reqBody.InterestRate / 100))
	reqBody.OutstandingBalance = reqBody.TotalAmountDue
	reqBody.Status = "Active"
	reqBody.StartLoanDate = now

	startDate, err := time.Parse("02 Jan 2006", reqBody.StartLoanDate)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	for i := 0; i < reqBody.LoanDurationWeeks; i++ {
		dueDate := startDate.AddDate(0, 0, i*7)
		billingSchedule = append(billingSchedule, billingScheduleModel.BillingSchedule{
			DueDate:    dueDate,
			WeekNumber: i + 1,
			AmountDue:  reqBody.TotalAmountDue / float64(reqBody.LoanDurationWeeks),
			Status:     "Pending",
			CreatedAt:  time.Now(),
		})
	}
	return statusCode, s.loanRepo.Create(reqBody, billingSchedule)
}
