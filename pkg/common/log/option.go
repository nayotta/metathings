package log_helper

import (
	"github.com/PeerXu/option-go"
	"github.com/sirupsen/logrus"
)

const (
	OPTION_LOGGER = "logger"
)

var (
	WithLogger, GetLogger = option.New[logrus.FieldLogger](OPTION_LOGGER)
)
