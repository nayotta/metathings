package metathings_device_cloud_service

import (
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type HeartbeatRequest struct {
	Module pb.OpModule
}

func (s *MetathingsDeviceCloudService) Heartbeat(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	req := new(HeartbeatRequest)
	err := dec.Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mdl_id := req.Module.GetId().GetValue()

	err = s.storage.Heartbeat(mdl_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	go s.try_to_build_device_connection_by_module_id(mdl_id)

	w.WriteHeader(http.StatusNoContent)
}

func (s *MetathingsDeviceCloudService) get_device_by_module_id(mdl_id string) (*pb.Device, error) {
	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	req := &pb.GetDeviceByModuleRequest{
		Module: &pb.OpModule{
			Id: &wrappers.StringValue{Value: mdl_id},
		},
	}

	res, err := cli.GetDeviceByModule(s.context(), req)
	if err != nil {
		return nil, err
	}

	return res.GetDevice(), nil
}

func (s *MetathingsDeviceCloudService) try_to_build_device_connection_by_module_id(mdl_id string) {
	err := s.storage.IsConnected(s.get_session_id(), mdl_id)
	switch err {
	case nil:
		// this instance are maintaining device connection, ignore
	case storage.ErrConnectedByOtherDeviceCloud:
		// other instance are maintaining device connection, ignore
	case storage.ErrNotConnected:
		// try to build device connection in current instance
		dev, err := s.get_device_by_module_id(mdl_id)
		if err != nil {
			s.get_logger().WithError(err).Errorf("failed to get device in deviced")
			return
		}

		dev_id := dev.Id
		// mark down instance session for the device
		err = s.storage.ConnectDevice(s.get_session_id(), dev_id)
		if err != nil {
			s.get_logger().WithError(err).Debugf("failed to lock connection in current instance, maybe locked by other instance")
			return
		}

		err = s.build_device_connection(dev_id)
		if err != nil {
			// unmark instance session on failed
			s.storage.UnconnectDevice(s.get_session_id(), dev_id)
			s.get_logger().WithError(err).Errorf("failed to build device connection")
			return
		}

		s.get_logger().WithFields(log.Fields{
			"module": mdl_id,
			"device": dev_id,
		}).Infof("build device connection")
	default:
		s.get_logger().WithError(err).Debugf("failed to get device connection status")
	}
}

func (s *MetathingsDeviceCloudService) build_device_connection(dev_id string) error {
	panic("unimplemented")
}
