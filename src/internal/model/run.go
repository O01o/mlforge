package m

import "time"

type Run struct {
	RunID            uint64    `db:"run_id"`
	ExperimentID     uint64    `db:"experiment_id"`
	RunName          string    `db:"name"`
	RunDescription   string    `db:"description"`
	RunFailureReason string    `db:"failure_reason"`
	RunStatus        uint16    `db:"status"`
	StartedAt        time.Time `db:"started_at"`
	EndedAt          time.Time `db:"ended_at"`
}

type CreateRun struct {
	ExperimentID   uint64 `db:"experiment_id"`
	RunName        string `db:"name"`
	RunDescription string `db:"description"`
	RunStatus      uint16 `db:"status"`
}
