package qm

const InsertParametersQuery = `
INSERT INTO parameters (run_id, name, value) 
VALUES (:run_id, :name, :value);
`

const GetParametersQuery = `
SELECT * FROM parameters 
WHERE run_id = ?;
`

const GetParameterIDsQuery = `
SELECT parameter_id FROM parameters 
WHERE run_id = ?;
`
