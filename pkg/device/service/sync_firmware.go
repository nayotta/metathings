package metathings_device_service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/objx"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
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

	var postfix string
	if len(sha256sum) > 12 {
		postfix = sha256sum[:12]
	} else {
		postfix = fmt.Sprintf("%d", time.Now().Unix())
	}

	dst := fmt.Sprintf("%v.%v", filepath.Base(src), postfix)
	if err = self.bs.Sync(self.context(), src, dst, uri, sha256sum); err != nil {
		logger.WithError(err).Debugf("failed to sync firmware")
		return err
	}

	logger.Infof("sync firmware, please restart!")

	return nil
}

func (self *MetathingsDeviceServiceImpl) do_sync_modules_firmware() {
	logger := self.logger.WithField("#method", "do_sync_modules_firmware")

	for _, m := range self.info.Modules {
		loop_logger := logger.WithField("module", m.Name)
		mdl, err := self.mdl_db.Lookup(m.Name)
		if err != nil {
			loop_logger.WithError(err).Warningf("failed to get module in database")
			continue
		}

		sf_req := &deviced_pb.OpUnaryCallValue{
			Method: &wrappers.StringValue{Value: "SyncFirmware"},
		}

		_, err = mdl.UnaryCall(context.Background(), sf_req)
		if err != nil {
			loop_logger.WithError(err).Warningf("failed to send sync firmware unary call")
		}
	}

	logger.Debugf("sync modules firmware")
}

func (self *MetathingsDeviceServiceImpl) get_device_firmware_descriptor(cli deviced_pb.DevicedServiceClient, ctx context.Context) (*deviced_pb.FirmwareDescriptor, error) {
	res, err := cli.ShowDeviceFirmwareDescriptor(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res.GetFirmwareDescriptor(), nil
}

func (self *MetathingsDeviceServiceImpl) sync_firmware() (err error) {
	logger := self.logger.WithField("#method", "sync_firmware")

	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to new deviced service client")
		return err
	}
	defer cfn()

	desc, err := self.get_device_firmware_descriptor(cli, self.context())
	if err != nil {
		logger.WithError(err).Errorf("failed to get device firmware descriptor")
		return err
	}

	desc_str, err := new(jsonpb.Marshaler).MarshalToString(desc.GetDescriptor_())
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal descriptor to json string")
		return err
	}

	descx, err := objx.FromJSON(desc_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal json string to objx")
		return err
	}

	dev_nxt_ver := descx.Get("device.version.next").String()
	if dev_nxt_ver != "" && dev_nxt_ver != self.GetVersion() {
		uri := descx.Get("device.uri.next").String()
		sha256sum := descx.Get("device.sha256.next").String()
		go self.do_sync_firmware(uri, sha256sum)
	}

	go self.do_sync_modules_firmware()

	return nil
}
