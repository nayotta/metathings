package metathings_device_cloud_profile_storage_core

import "github.com/PeerXu/option-go"

const (
	OPTION_DEFAULT_PROFILE         = "defaultProfile"
	OPTION_DISABLE_DEFAULT_PROFILE = "disableDefaultProfile"
)

var (
	WithDefaultProfile, GetDefaultProfile               = option.New[string](OPTION_DEFAULT_PROFILE)
	WithDisableDefaultProfile, GetDisableDefaultProfile = option.New[bool](OPTION_DISABLE_DEFAULT_PROFILE)
)

func DefaultNewStorageOptions() option.Option {
	return option.NewOption(map[string]any{
		OPTION_DEFAULT_PROFILE: "default",
	})
}

func DefaultGetProfileByDeviceOptions() option.Option {
	return option.NewOption(map[string]any{
		OPTION_DISABLE_DEFAULT_PROFILE: false,
	})
}
