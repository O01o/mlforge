package rm

import (
	"github.com/jmoiron/sqlx"
)

type plotRepository struct {
	db *sqlx.DB
}

func NewPlotRepository(db *sqlx.DB) *plotRepository {
	return &plotRepository{db: db}
}
