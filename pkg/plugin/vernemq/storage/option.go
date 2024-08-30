package metathings_plugin_vernemq_storage

import (
	"github.com/PeerXu/option-go"
	"github.com/PeerXu/registerer-go"
)

const (
	OPTION_STORAGE = "storage"
)

var (
	WithStorage, GetStorage = option.New[Storage](OPTION_STORAGE)

	Register, New = registerer.Pair[Storage]()
)
