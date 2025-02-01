package borrower

import "time"

type Borrower struct {
	Id          uint32
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
