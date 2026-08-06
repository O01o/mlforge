package sem

import (
	rm "mlforge/internal/repository/db/mysql"
	sei "mlforge/internal/service/interface"

	"github.com/jmoiron/sqlx"
)

type setup struct {
	db *sqlx.DB
}

func NewSetupService(db *sqlx.DB) sei.SetupService {
	return &setup{db: db}
}

func (s *setup) Setup() error {
	repo := rm.NewSetupRepository(s.db)
	err := repo.Setup()
	if err != nil {
		return err
	}
	return nil
}
