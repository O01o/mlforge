package sem

import (
	"context"
	m "mlforge/internal/model"
	rm "mlforge/internal/repository/mysql"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"

	"github.com/jmoiron/sqlx"
)

type runService struct {
	db *sqlx.DB
}

func NewRunService(db *sqlx.DB) sei.RunService {
	return &runService{db: db}
}

func (s *runService) CreateRun(req *sc.CreateRunRequest) (uint64, []uint64, error) {
	ctx := context.Background()
	tx, err := s.db.Beginx()
	if err != nil {
		return 0, nil, err
	}
	defer tx.Rollback()

	repoRun := rm.NewRunRepository(ctx, tx)
	runID, err := repoRun.CreateRun(m.CreateRun{
		ExperimentID:   req.ExperimentID,
		RunName:        req.RunName,
		RunDescription: req.RunDescription,
		RunStatus:      0,
	})
	if err != nil {
		return 0, nil, err
	}

	var parameters []m.CreateParameter
	for _, param := range req.Parameters {
		parameters = append(parameters, m.CreateParameter{
			RunID:          runID,
			ParameterName:  param.ParameterName,
			ParameterValue: param.ParameterValue,
		})
	}
	repoParameter := rm.NewParameterRepository(ctx, tx)
	err = repoParameter.CreateParameters(parameters)
	if err != nil {
		return 0, nil, err
	}

	var metrics []m.CreateMetric
	for _, metric := range req.Metrics {
		metrics = append(metrics, m.CreateMetric{
			RunID:      runID,
			MetricName: metric.MetricName,
		})
	}
	repoMetric := rm.NewMetricRepository(ctx, tx)
	metricIDs, err := repoMetric.CreateMetrics(metrics)
	if err != nil {
		return 0, nil, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, nil, err
	}

	return runID, metricIDs, nil
}

func (s *runService) FinishRunByID(id uint64) error {
	ctx := context.Background()
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repo := rm.NewRunRepository(ctx, tx)
	err = repo.FinishRunByID(id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *runService) FailRunByID(id uint64) error {
	ctx := context.Background()
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	repo := rm.NewRunRepository(ctx, tx)
	err = repo.FailRunByID(id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *runService) DeleteRunByID(id uint64) error {
	ctx := context.Background()
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repo := rm.NewRunRepository(ctx, tx)
	err = repo.DeleteRunByID(id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
