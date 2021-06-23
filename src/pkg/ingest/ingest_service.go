package ingest

type Ingest interface {
	Initialise() error
	Create() (*string, error)
}
