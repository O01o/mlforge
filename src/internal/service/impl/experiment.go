package sem

import (
	"context"
	m "mlforge/internal/model"
	rm "mlforge/internal/repository/mysql"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"

	"github.com/jmoiron/sqlx"
)

type experimentService struct {
	db *sqlx.DB
}

func NewExperimentService(db *sqlx.DB) sei.ExperimentService {
	return &experimentService{db: db}
}

func (s *experimentService) CreateExperiment(req *sc.CreateExperimentRequest) (uint64, error) {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	repo := rm.NewExperimentRepository(ctx, tx)
	id, err := repo.CreateExperiment(&m.CreateExperiment{
		ExperimentName:        req.ExperimentName,
		ExperimentDescription: req.ExperimentDescription,
	})
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (s *experimentService) GetExperimentByID(id uint64) (*sc.GetExperimentResponse, error) {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repoExperiment := rm.NewExperimentRepository(ctx, tx)
	experiment, err := repoExperiment.GetExperimentByID(id)
	if err != nil {
		return nil, err
	}

	repoRun := rm.NewRunRepository(ctx, tx)
	runs, err := repoRun.GetRuns(id)
	if err != nil {
		return nil, err
	}

	var runds []sc.RunDetails
	for _, run := range runs {
		runs := sc.RunSummary{
			RunID:            run.RunID,
			RunName:          run.RunName,
			RunDescription:   run.RunDescription,
			RunFailureReason: run.RunFailureReason,
			RunStatus:        run.RunStatus,
			StartedAt:        run.StartedAt,
			EndedAt:          run.EndedAt,
		}

		repoParameter := rm.NewParameterRepository(ctx, tx)
		parameters, err := repoParameter.GetParameters(run.RunID)
		if err != nil {
			return nil, err
		}
		var parameterS []sc.Parameter
		for _, p := range parameters {
			parameterS = append(parameterS, sc.Parameter{
				ParameterID:    p.ParameterID,
				ParameterName:  p.ParameterName,
				ParameterValue: p.ParameterValue,
			})
		}

		repoMetric := rm.NewMetricRepository(ctx, tx)
		metrics, err := repoMetric.GetMetrics(run.RunID)
		if err != nil {
			return nil, err
		}
		var metricS []sc.Metric
		for _, m := range metrics {
			metricS = append(metricS, sc.Metric{
				MetricID:   m.MetricID,
				MetricName: m.MetricName,
			})
		}

		runds = append(runds, sc.RunDetails{
			RunSummary: runs,
			Parameters: parameterS,
			Metrics:    metricS,
		})
	}

	tx.Commit()
	return &sc.GetExperimentResponse{
		ExperimentDetails: &sc.ExperimentDetails{
			ExperimentSummary: sc.ExperimentSummary{
				ExperimentID:          experiment.ExperimentID,
				ExperimentName:        experiment.ExperimentName,
				ExperimentDescription: experiment.ExperimentDescription,
				CreatedAt:             experiment.CreatedAt,
				UpdatedAt:             experiment.UpdatedAt,
			},
			RunDetails: runds,
		},
	}, nil
}

func (s *experimentService) GetExperimentList() (*sc.GetExperimentSummariesResponse, error) {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	repo := rm.NewExperimentRepository(ctx, tx)
	experiments, err := repo.GetExperimentList()
	if err != nil {
		return nil, err
	}
	tx.Commit()
	var experimentS []*sc.ExperimentSummary
	for _, experiment := range experiments {
		experimentS = append(
			experimentS,
			&sc.ExperimentSummary{
				ExperimentID:          experiment.ExperimentID,
				ExperimentName:        experiment.ExperimentName,
				ExperimentDescription: experiment.ExperimentDescription,
				CreatedAt:             experiment.CreatedAt,
				UpdatedAt:             experiment.UpdatedAt,
			},
		)
	}
	return &sc.GetExperimentSummariesResponse{
		ExperimentSummaries: experimentS,
	}, nil
}

func (s *experimentService) UpdateExperimentByID(id uint64, req *sc.CreateExperimentRequest) error {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	repo := rm.NewExperimentRepository(ctx, tx)
	err = repo.UpdateExperimentByID(id, &m.UpdateExperiment{
		ExperimentID:          id,
		ExperimentName:        req.ExperimentName,
		ExperimentDescription: req.ExperimentDescription,
	})
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (s *experimentService) DeleteExperimentByID(id uint64) error {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	repo := rm.NewExperimentRepository(ctx, tx)
	err = repo.DeleteExperimentByID(id)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
