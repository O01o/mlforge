package qm

const InsertPlotsQuery = `
INSERT INTO plots (metric_id, step, value) 
VALUES (:metric_id, :step, :value);
`

const GetPlotsByRangeOfStepQuery = `
SELECT plot_id FROM plots 
WHERE metric_id = ? AND step BETWEEN ? AND ?;
`
