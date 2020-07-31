package metathings_sms_sdk

import (
	"context"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	"github.com/sirupsen/logrus"
)

type DummySmsSender struct {
	logger logrus.FieldLogger
}

func (ss *DummySmsSender) SendSms(ctx context.Context, id string, numbers []string, arguments map[string]string) error {
	ss.logger.WithFields(logrus.Fields{
		"id":        id,
		"numbers":   numbers,
		"arguments": arguments,
	}).Debugf("send sms")

	return nil
}

func NewDummySmsSender(args ...interface{}) (SmsSender, error) {
	var logger logrus.FieldLogger

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	return &DummySmsSender{
		logger: logger,
	}, nil
}

func init() {
	register_sms_sender_factory("dummy", NewDummySmsSender)
}
