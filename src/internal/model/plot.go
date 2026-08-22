package m

type Plot struct {
	MetricID  uint64 `db:"metric_id"`
	PlotStep  string `db:"step"`
	PlotValue string `db:"value"`
}
