package metathings_device_cloud_service

import (
	"net/http"

	"github.com/samber/lo"

	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (s *MetathingsDeviceCloudService) UnbindDevices(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger().WithField("#method", "UnbindDevices")

	ctx := s.context()
	tkn_txt := GetTokenFromHeader(r)
	_, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req pb.UnbindDevicesRequest
	if err = ParseHttpRequestBody(r, &req); err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	deviceIds := lo.Map(req.GetDevices(), func(dev *deviced_pb.OpDevice, _ int) string {
		return dev.GetId().GetValue()
	})
	deviceIds = lo.Uniq(deviceIds)

	logger = logger.WithField("devices.id", deviceIds)

	if err = s.profile_storage.UnbindDevices(ctx, deviceIds); err != nil {
		logger.WithError(err).Errorf("failed to unbind devices in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	logger.Infof("unbind devices")

	return
}
