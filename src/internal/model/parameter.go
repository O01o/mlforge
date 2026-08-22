package m

type Parameter struct {
	ParameterID    uint64 `db:"parameter_id"`
	RunID          uint64 `db:"run_id"`
	ParameterName  string `db:"name"`
	ParameterValue string `db:"value"`
}

type CreateParameter struct {
	RunID          uint64 `db:"run_id"`
	ParameterName  string `db:"name"`
	ParameterValue string `db:"value"`
}
