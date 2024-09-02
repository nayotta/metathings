package metathings_device_cloud_profile_storage_core

import (
	"github.com/PeerXu/registerer-go"

	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
)

var (
	Register, New = registerer.Pair[intf.ProfileStorage]()
)
