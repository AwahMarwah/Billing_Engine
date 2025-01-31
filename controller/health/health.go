package health

import (
	healthRepo "billing_engine/repository/health"
	"billing_engine/service/health"
	"database/sql"
)

type controller struct {
	healthService health.IService
}

func NewController(db *sql.DB) *controller {
	return &controller{healthService: health.NewService(healthRepo.NewRepo(db))}
}
