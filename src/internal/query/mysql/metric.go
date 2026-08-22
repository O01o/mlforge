package qm

const InsertMetricsQuery = `
INSERT INTO metrics (run_id, name) 
VALUES (:run_id, :name);
`

const GetMetricsQuery = `
SELECT * FROM metrics 
WHERE run_id = ?;
`

const GetMetricIDsQuery = `
SELECT metric_id FROM metrics 
WHERE run_id = ?;
`
