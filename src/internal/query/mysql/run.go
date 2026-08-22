package qm

const InsertRunQuery = `
INSERT INTO runs (experiment_id, name, description, status) 
VALUES (:experiment_id, :name, :description, 0);
`

const GetRunsByExperimentIDQuery = `
SELECT * FROM runs 
WHERE experiment_id = ?;
`

const FinishRunByIDQuery = `
UPDATE runs 
SET status = 1 
WHERE experiment_id = ? AND run_id = ?;
`

const FailRunByIDQuery = `
UPDATE runs 
SET status = 2 
WHERE experiment_id = ? AND run_id = ?;
`

const DeleteRunByIDQuery = `
DELETE FROM runs 
WHERE experiment_id = ? AND run_id = ?;
`
