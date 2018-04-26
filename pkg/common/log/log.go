package log_helper

import (
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
