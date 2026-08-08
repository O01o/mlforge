package qm

const ExperimentInsertQuery = `
INSERT INTO experiments (name, description) VALUES (?, ?);
`

const ExperimentGetListQuery = `
SELECT * FROM experiments;
`

const ExperimentGetByIDQuery = `
SELECT * FROM experiments 
WHERE experiment_id = ?;
`

const ExperimentDeleteByIDQuery = `
DELETE FROM experiments 
WHERE experiment_id = ?;
`
