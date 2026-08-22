package qm

const InsertExperimentQuery = `
INSERT INTO experiments (name, description) 
VALUES (:name, :description);
`

const GetExperimentsQuery = `
SELECT * FROM experiments;
`

const GetExperimentByIDQuery = `
SELECT * FROM experiments 
WHERE experiment_id = ?;
`

const DeleteExperimentByIDQuery = `
DELETE FROM experiments 
WHERE experiment_id = ?;
`
