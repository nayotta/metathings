package metathings_evaluatord_sdk

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/cast"
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
	return cast.ToString(ctx.Value("data-launcher-token"))
}

func WithDevice(ctx context.Context, dev string) context.Context {
	return context.WithValue(ctx, "data-launcher-device", dev)
}

func ExtractDevice(ctx context.Context) string {
	return cast.ToString(ctx.Value("data-launcher-device"))
}

func WithTags(ctx context.Context, tags map[string]string) context.Context {
	return context.WithValue(ctx, "data-launcher-tags", tags)
}

func ExtractTags(ctx context.Context) map[string]string {
	return cast.ToStringMapString(ctx.Value("data-launcher-tags"))
}

func WithTimestamp(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, "data-launcher-timestamp", ts)
}

func ExtractTimestamp(ctx context.Context) time.Time {
	ts, err := cast.ToTimeE(ctx.Value("data-launcher-timestamp"))
	if err != nil {
		ts = time.Now()
	}

	return ts
}
