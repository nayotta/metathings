package metathings_plugin_vernemq_service

import (
	"github.com/PeerXu/option-go"
)

const (
	OPTION_WEBHOOK_SECRET = "webhookSecret"
)

var (
	WithWebhookSecret, GetWebhookSecret = option.New[string](OPTION_WEBHOOK_SECRET)
)
