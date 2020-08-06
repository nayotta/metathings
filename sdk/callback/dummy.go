package metathings_callback_sdk

type DummyCallback struct{}

func (cb *DummyCallback) Emit(data interface{}, tags map[string]string) error {
	return nil
}

func NewDummyCallback(args ...interface{}) (Callback, error) {
	return &DummyCallback{}, nil
}

func init() {
	register_callback_factory("dummy", NewDummyCallback)
}
