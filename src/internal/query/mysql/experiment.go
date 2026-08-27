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

const UpdateExperimentByIDQuery = `
UPDATE experiments 
SET name = :name, description = :description 
WHERE experiment_id = :id;
`

const DeleteExperimentByIDQuery = `
DELETE FROM experiments 
WHERE experiment_id = ?;
`
