package metathings_device_service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/objx"
)

func (self *MetathingsDeviceServiceImpl) do_sync_firmware(uri, sha256sum string) error {
	logger := self.logger.WithFields(logrus.Fields{
		"uri":       uri,
		"sha256sum": sha256sum,
	})

	src, err := os.Executable()
	if err != nil {
		logger.WithError(err).Debugf("failed to get executable info")
		return err
	}

	dst := fmt.Sprintf("%v.%v", filepath.Base(src), time.Now().Unix())
	if err = self.bs.Sync(self.context(), src, dst, uri, sha256sum); err != nil {
		logger.WithError(err).Debugf("failed to sync firmware")
		return err
	}

	logger.Infof("sync firmware, please restart!")

	return nil
}

func (self *MetathingsDeviceServiceImpl) sync_firmware() (err error) {
	logger := self.logger.WithField("#method", "sync_firmware")

	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to new deviced service client")
		return err
	}
	defer cfn()

	fd_res, err := cli.ShowDeviceFirmwareDescriptor(self.context(), &empty.Empty{})
	if err != nil {
		logger.WithError(err).Errorf("failed to show device firmware descriptor")
		return err
	}

	fd := fd_res.GetFirmwareDescriptor()
	desc := fd.GetDescriptor_()
	desc_str, err := new(jsonpb.Marshaler).MarshalToString(desc)
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal descriptor to json")
		return err
	}

	descx, err := objx.FromJSON(desc_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to convert descriptor json to objx")
		return err
	}

	dev_nxt_ver := descx.Get("device.version.next").String()

	if dev_nxt_ver != self.GetVersion() {
		uri := descx.Get("device.uri.next").String()
		sha256sum := descx.Get("device.sha256.next").String()
		go self.do_sync_firmware(uri, sha256sum)
	}

	return nil
}
