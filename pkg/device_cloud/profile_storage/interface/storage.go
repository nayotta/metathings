package metathings_device_cloud_profile_storage_interface

import (
	"context"

	"github.com/PeerXu/option-go"
)

type Profile struct {
	Name    string `json:"name"`
	Driver  string `json:"driver"`
	Address string `json:"address"`
	Port    string `json:"port"`
}

type ProfileStorage interface {
	CreateOrUpdateProfile(ctx context.Context, profile Profile) error
	DeleteProfile(ctx context.Context, profileName string) error
	GetProfile(ctx context.Context, profileName string) (Profile, error)
	GetProfileByDevice(ctx context.Context, deviceId string, opts ...option.ApplyOption) (Profile, error)
	BindDevices(ctx context.Context, deviceIds []string, profileName string) error
	UnbindDevices(ctx context.Context, deviceIds []string) error
}
