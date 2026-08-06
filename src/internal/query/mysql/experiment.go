package qm

const ExperimentInsertQuery = `
INSERT INTO experiments (name, description) VALUES (?, ?);
`

const ExperimentGetListQuery = `
SELECT experiment_id, experiment_name, experiment_description
FROM experiments;
`

const ExperimentGetByIDQuery = `
SELECT experiment_id, experiment_name, experiment_description
FROM experiments 
WHERE experiment_id = ?;
`

const ExperimentDeleteByIDQuery = `
DELETE FROM experiments 
WHERE experiment_id = ?;
`
