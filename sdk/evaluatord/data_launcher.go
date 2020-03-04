package metathings_evaluatord_sdk

import (
	"context"
	"sync"
)

type DataLauncher interface {
	Launch(context.Context, Resource, Data) error
}

type DataLauncherFactory func(...interface{}) (DataLauncher, error)

var data_launcher_factories_once sync.Once
var data_launcher_factories map[string]DataLauncherFactory

func registry_data_launcher(name string, fty DataLauncherFactory) {
	data_launcher_factories_once.Do(func() {
		data_launcher_factories = make(map[string]DataLauncherFactory)
	})

	data_launcher_factories[name] = fty
}

func NewDataLauncher(name string, args ...interface{}) (dl DataLauncher, err error) {
	fty, ok := data_launcher_factories[name]
	if !ok {
		return nil, ErrUnsupportedDataLauncherFactory
	}

	return fty(args...)
}

func WithToken(ctx context.Context, tkn string) context.Context {
	return context.WithValue(ctx, "data-launcher-token", tkn)
}

func ExtractToken(ctx context.Context) string {
	return ctx.Value("data-launcher-token").(string)
}
