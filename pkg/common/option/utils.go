package option_helper

import (
	optiongo "github.com/PeerXu/option-go"
	"github.com/sirupsen/logrus"
)

const (
	OPTION_LOGGER = "logger"
)

var (
	WithLogger, GetLogger = optiongo.New[*logrus.Entry](OPTION_LOGGER)
)
