package metathings_sms_sdk

import (
	"context"

	"github.com/nayotta/pomelo-sdk-go"
	"github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type PomeloSmsSenderOption struct {
	BaseURL  string
	Insecure bool
}

func NewPomeloSmsSenderOption() *PomeloSmsSenderOption {
	return &PomeloSmsSenderOption{}
}

type PomeloSmsSender struct {
	sdk    pomelo.PomeloSDK
	logger logrus.FieldLogger
}

func (ss *PomeloSmsSender) SendSms(ctx context.Context, id string, numbers []string, arguments map[string]string) error {
	logger := ss.logger.WithFields(logrus.Fields{
		"sms":     id,
		"numbers": numbers,
	})

	err := ss.sdk.SendSMS(ctx, id, numbers, arguments)
	if err != nil {
		logger.WithError(err).Debugf("failed to send sms")
		return err
	}

	logger.Debugf("send sms")

	return nil
}

func NewPomeloSmsSender(args ...interface{}) (SmsSender, error) {
	var logger logrus.FieldLogger

	opt := NewPomeloSmsSenderOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"base_url": opt_helper.ToString(&opt.BaseURL),
		"insecure": opt_helper.ToBool(&opt.Insecure),
		"logger":   opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	sdk, err := pomelo.NewPomeloSDK(&pomelo.PomeloSDKOption{
		BaseURL:            opt.BaseURL,
		InsecureSkipVerify: opt.Insecure,
	})
	if err != nil {
		return nil, err
	}

	return &PomeloSmsSender{
		sdk:    sdk,
		logger: logger,
	}, nil
}

func init() {
	register_sms_sender_factory("pomelo", NewPomeloSmsSender)
	register_sms_sender_factory("default", NewPomeloSmsSender)
}
