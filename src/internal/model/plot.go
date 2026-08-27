package m

type Plot struct {
	MetricID  uint64  `db:"metric_id"`
	PlotStep  uint64  `db:"step"`
	PlotValue float64 `db:"value"`
}
