package option_helper

import (
	"net/http"

	optiongo "github.com/PeerXu/option-go"
	"github.com/sirupsen/logrus"
)

const (
	OPTION_LOGGER      = "logger"
	OPTION_HTTP_CLIENT = "httpClient"
)

var (
	WithLogger, GetLogger         = optiongo.New[*logrus.Entry](OPTION_LOGGER)
	WithHTTPClient, GetHTTPClient = optiongo.New[*http.Client](OPTION_HTTP_CLIENT)
)
