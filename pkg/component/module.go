package metathings_component

type Module interface {
	Start() error
	Stop() error
	Err() error
}
