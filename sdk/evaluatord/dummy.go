package metathings_evaluatord_sdk

import "context"

type DummyDataLauncher struct{}

func (ddl *DummyDataLauncher) Launch(ctx context.Context, src Resource, dat Data) error {
	return nil
}

func NewDummyDataLauncher(args ...interface{}) (DataLauncher, error) {
	return &DummyDataLauncher{}, nil
}

func init() {
	registry_data_launcher("dummy", NewDummyDataLauncher)
}
