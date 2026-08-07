package rm

import (
	"mlforge/internal/assets"
	ri "mlforge/internal/repository/interface"

	"github.com/jmoiron/sqlx"
)

type setupRepository struct {
	db *sqlx.DB
}

func NewSetupRepository(db *sqlx.DB) ri.SetupRepository {
	return &setupRepository{db: db}
}

func (r *setupRepository) Setup() error {
	ddl, err := assets.Files.ReadFile("db/ddl.sql")
	if err != nil {
		return err
	}
	_, err = r.db.Exec(string(ddl))
	if err != nil {
		return err
	}
	return nil
}
