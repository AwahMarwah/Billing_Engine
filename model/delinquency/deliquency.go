package delinquency

import "time"

type Deliquency struct {
	DelinquencyId int
	LoanId        int
	Delinquent    bool
	DateChecked   time.Time
}
