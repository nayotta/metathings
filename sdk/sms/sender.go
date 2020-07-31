package metathings_sms_sdk

import (
	"context"
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type SmsSender interface {
	SendSms(ctx context.Context, id string, numbers []string, arguments map[string]string) error
}

type SmsSenderFactory func(...interface{}) (SmsSender, error)

var sms_sender_factories map[string]SmsSenderFactory
var sms_sender_factories_once sync.Once

func register_sms_sender_factory(name string, fty SmsSenderFactory) {
	sms_sender_factories_once.Do(func() {
		sms_sender_factories = make(map[string]SmsSenderFactory)
	})
	sms_sender_factories[name] = fty
}

func NewSmsSender(name string, args ...interface{}) (SmsSender, error) {
	fty, ok := sms_sender_factories[name]
	if !ok {
		return nil, ErrUnsupportedSmsSenderDriver
	}

	return fty(args...)
}

func ToSmsSender(ss *SmsSender) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		if *ss, ok = val.(SmsSender); !ok {
			return opt_helper.InvalidArgument(key)
		}
		return nil
	}
}
