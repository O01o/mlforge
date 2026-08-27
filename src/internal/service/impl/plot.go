package sem

import (
	"context"
	m "mlforge/internal/model"
	rm "mlforge/internal/repository/mysql"
	sc "mlforge/internal/schema"
	sei "mlforge/internal/service/interface"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type plotService struct {
	db *sqlx.DB
}

func NewPlotService(db *sqlx.DB) sei.PlotService {
	return &plotService{db: db}
}

func (s *plotService) CreatePlots(metricId uint64, req *sc.CreatePlotsRequest) error {
	var plots []m.Plot
	for stepStr, value := range req.Plots {
		step, err := strconv.ParseUint(stepStr, 10, 64)
		if err != nil {
			return err
		}
		plots = append(plots, m.Plot{
			MetricID:  metricId,
			PlotStep:  step,
			PlotValue: value,
		})
	}

	ctx := context.Background()
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repoPlot := rm.NewPlotRepository(ctx, tx)
	err = repoPlot.CreatePlots(plots)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *plotService) GetPlots(req *sc.GetPlotsRequest) (*sc.GetPlotsResponse, error) {
	ctx := context.Background()
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repoMetric := rm.NewMetricRepository(ctx, tx)
	repoPlot := rm.NewPlotRepository(ctx, tx)
	var runMetricPlots []sc.RunMetricPlots
	for _, runID := range req.RunIDs {
		metricIDs, err := repoMetric.GetMetricIDs(runID)
		if err != nil {
			return nil, err
		}
		var metricPlots []sc.MetricPlots
		for _, metricID := range metricIDs {
			plots, err := repoPlot.GetPlots(
				metricID,
				req.PlotRange.StartStep,
				req.PlotRange.EndStep,
			)
			if err != nil {
				return nil, err
			}
			plotMap := make(map[string]float64)
			for _, plot := range plots {
				plotMap[strconv.FormatUint(plot.PlotStep, 10)] = plot.PlotValue
			}
			metricPlots = append(
				metricPlots, sc.MetricPlots{
					MetricID: metricID,
					Plots:    plotMap,
				},
			)
		}
		runMetricPlots = append(
			runMetricPlots,
			sc.RunMetricPlots{
				RunID:       runID,
				MetricPlots: metricPlots,
			},
		)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &sc.GetPlotsResponse{RunMetricPlots: runMetricPlots}, nil
}
