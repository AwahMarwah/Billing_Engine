package health

import healthRepo "billing_engine/repository/health"

type (
	IService interface {
		Check() (err error)
	}

	service struct {
		healthRepo healthRepo.IRepo
	}
)

func NewService(healthRepo healthRepo.IRepo) IService {
	return &service{healthRepo: healthRepo}
}
