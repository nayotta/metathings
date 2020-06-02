package metathings_evaluatord_timer

import (
	"context"
	"sync"
)

type TimerOption func(map[string]interface{})

func SetId(id string) func(map[string]interface{}) {
	return func(opt map[string]interface{}) {
		opt["id"] = id
	}
}

func SetSchedule(schedule string) func(map[string]interface{}) {
	return func(opt map[string]interface{}) {
		opt["schedule"] = schedule
	}
}

func SetTimezone(timezone string) func(map[string]interface{}) {
	return func(opt map[string]interface{}) {
		opt["timezone"] = timezone
	}
}

func SetEnabled(enabled bool) func(map[string]interface{}) {
	return func(opt map[string]interface{}) {
		opt["enabled"] = enabled
	}
}

type TimerApi interface {
	Id() string
	Schedule() string
	Timezone() string
	Enabled() bool

	Set(context.Context, ...TimerOption) error

	Delete(context.Context) error
}

type TimerBackend interface {
	Create(context.Context, ...TimerOption) (TimerApi, error)
	Get(context.Context, string) (TimerApi, error)
}

type NewTimerBackendFactory func(...interface{}) (TimerBackend, error)

var timer_backend_factories map[string]NewTimerBackendFactory
var timer_backend_factories_once sync.Once

func register_timer_backend_factory(name string, fty NewTimerBackendFactory) {
	timer_backend_factories_once.Do(func() {
		timer_backend_factories = map[string]NewTimerBackendFactory{}
	})
	timer_backend_factories[name] = fty
}

func NewTimerBackend(name string, args ...interface{}) (TimerBackend, error) {
	fty, ok := timer_backend_factories[name]
	if !ok {
		return nil, ErrUnknownTimerBackendDriver
	}

	return fty(args...)
}
