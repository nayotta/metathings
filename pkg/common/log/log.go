package log_helper

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

func NewLogger(service_name string, level string) (log.FieldLogger, error) {
	logger := log.New()
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(lvl)
	return logger.WithFields(log.Fields{
		"#service": service_name,
	}), nil
}

type GetLoggerer struct {
	logger log.FieldLogger
}

func (l *GetLoggerer) GetLogger() log.FieldLogger {
	pc := make([]uintptr, 2)
	runtime.Callers(1, pc)
	return l.logger.WithField("#caller", runtime.FuncForPC(pc[1]).Name())
}

func NewGetLoggerer(logger log.FieldLogger) *GetLoggerer {
	return &GetLoggerer{logger: logger}
}
