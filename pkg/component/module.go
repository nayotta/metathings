package metathings_component

type Module interface {
	Start() error
	Stop() error
	Wait() chan bool
	Err() error
}
