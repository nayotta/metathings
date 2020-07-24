package metathings_device_cloud_service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
)

func (dc *DeviceConnection) sync_firmware() error {
	go dc.do_sync_modules_firmware()

	return nil
}

func (dc *DeviceConnection) do_sync_modules_firmware() error {
	logger := dc.logger.WithFields(log.Fields{
		"#method": "do_sync_modules_firmware",
		"device":  dc.info.Id,
	})

	for _, m := range dc.info.Modules {
		loop_logger := logger.WithField("module", m.Name)
		mdl_prx, err := dc.get_module_proxy(m.Name)
		if err != nil {
			loop_logger.WithError(err).Warningf("failed to get module proxy")
			continue
		}

		sf_req := &empty.Empty{}
		any_req, err := ptypes.MarshalAny(sf_req)
		if err != nil {
			loop_logger.WithError(err).Debugf("failed to marshal request to ANY type")
		}

		_, err = mdl_prx.UnaryCall(context.Background(), "SyncFirmware", any_req)
		if err != nil {
			loop_logger.WithError(err).Warningf("failed to send sync firmware unary call")
		}
	}

	logger.Debugf("sync modules firmware")

	return nil
}
