package rsi

type ArtifactRepository interface {
	Save(artifact *Artifact) error
	FindByID(id string) (*Artifact, error)
	FindByRunID(runID string) ([]*Artifact, error)
	DeleteByID(id string) error
	DeleteByRunID(runID string) error
}
