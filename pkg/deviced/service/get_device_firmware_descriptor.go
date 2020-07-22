package metathings_deviced_service

import (
	"context"
	"fmt"
	"os"

	"github.com/stretchr/objx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) get_device_firmware_descriptor(ctx context.Context, dev_id_str string) (fd_s *storage.FirmwareDescriptor, error error) {
	logger := self.logger.WithField("device", dev_id_str)

	dev_s, err := self.storage.GetDevice(ctx, dev_id_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to get device in storage")
		return nil, self.ParseError(err)
	}

	var fd_s_desc_str string
	fd_s, err = self.storage.GetDeviceFirmwareDescriptor(ctx, dev_id_str)
	if err != nil {
		if err != storage.RecordNotFound {
			logger.WithError(err).Errorf("failed to get device firmware descriptor in storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		} else {
			fd_name_str := "unknown"
			fd_id_str := "unknown"
			fd_s = &storage.FirmwareDescriptor{
				Id:   &fd_id_str,
				Name: &fd_name_str,
			}
			fd_s_desc_str = "{}"
		}
	} else {
		fd_s_desc_str = *fd_s.Descriptor
	}

	fdsx, err := objx.FromJSON(fd_s_desc_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to parse device firmware descriptor")
		return nil, self.ParseError(err)
	}

	fdx := objx.New(map[string]interface{}{})
	dev_cur_ver, err := self.simple_storage.GetObjectContentSync(&simple_storage.Object{
		Device: dev_id_str,
		Prefix: "/sys/firmware/device/version",
		Name:   "current",
	})
	if err != nil {
		if err != os.ErrNotExist {
			logger.WithError(err).Errorf("failed to get current device version")
			return nil, self.ParseError(err)
		}

		dev_cur_ver = []byte("unknown")
	}

	fdx.Set("device.version.current", string(dev_cur_ver))

	if val := fdsx.Get("device.version.next").String(); val != "" && val != string(dev_cur_ver) {
		fdx.Set("device.version.next", val)

		if val = fdsx.Get("device.sha256.next").String(); val != "" {
			fdx.Set("device.sha256.next", val)
		}

		if val = fdsx.Get("device.uri.next").String(); val != "" {
			fdx.Set("device.uri.next", val)
		}
	}

	for _, mdl_s := range dev_s.Modules {
		mdl_cur_ver, err := self.simple_storage.GetObjectContentSync(&simple_storage.Object{
			Device: dev_id_str,
			Prefix: fmt.Sprintf("/sys/firmware/modules/%s/version", *mdl_s.Name),
			Name:   "current",
		})

		if err != nil {
			if err != os.ErrNotExist {
				logger.WithError(err).Errorf("failed to get current module version")
				return nil, self.ParseError(err)
			}

			mdl_cur_ver = []byte("unknown")
		}

		fdx.Set(fmt.Sprintf("modules.%s.version.current", *mdl_s.Name), string(mdl_cur_ver))

		if val := fdsx.Get(fmt.Sprintf("modules.%s.version.next", *mdl_s.Name)).String(); val != "" && val != string(dev_cur_ver) {
			fdx.Set(fmt.Sprintf("modules.%s.version.next", *mdl_s.Name), val)

			if val = fdsx.Get(fmt.Sprintf("modules.%s.sha256.next", *mdl_s.Name)).String(); val != "" {
				fdx.Set(fmt.Sprintf("modules.%s.sha256.next", *mdl_s.Name), val)
			}

			if val = fdsx.Get(fmt.Sprintf("modules.%s.uri.next", *mdl_s.Name)).String(); val != "" {
				fdx.Set(fmt.Sprintf("modules.%s.uri.next", *mdl_s.Name), val)
			}
		}
	}

	var buf string
	if buf, err = fdx.JSON(); err != nil {
		logger.WithError(err).Errorf("failed to parse firmware descriptor map to json string")
		return nil, self.ParseError(err)
	}

	fd_s.Descriptor = &buf

	return fd_s, nil
}

func (self *MetathingsDevicedService) GetDeviceFirmwareDescriptor(ctx context.Context, req *pb.GetDeviceFirmwareDescriptorRequest) (*pb.GetDeviceFirmwareDescriptorResponse, error) {
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()
	logger := self.logger.WithField("device", dev_id_str)

	fd_s, err := self.get_device_firmware_descriptor(ctx, dev_id_str)
	if err != nil {
		return nil, err
	}

	res := &pb.GetDeviceFirmwareDescriptorResponse{
		FirmwareDescriptor: copy_firmware_descriptor(fd_s),
	}

	logger.Infof("get device firmware descriptor")

	return res, nil
}
