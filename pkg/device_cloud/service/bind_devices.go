package metathings_device_cloud_service

import (
	"net/http"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"

	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (s *MetathingsDeviceCloudService) BindDevices(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger().WithField("#method", "BindDevices")

	ctx := s.context()
	tkn_txt := GetTokenFromHeader(r)
	_, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req pb.BindDevicesRequest
	if err = ParseHttpRequestBody(r, &req); err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deviceIds := lo.Map(req.GetDevices(), func(dev *deviced_pb.OpDevice, _ int) string {
		return dev.GetId().GetValue()
	})
	deviceIds = lo.Uniq(deviceIds)
	profile := req.GetProfile()
	profileName := profile.GetName().GetValue()

	logger = logger.WithFields(logrus.Fields{
		"devices.id":   deviceIds,
		"profile.name": profileName,
	})

	if err = s.profile_storage.BindDevices(ctx, deviceIds, profileName); err != nil {
		logger.WithError(err).Errorf("failed to bind devices in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	logger.Infof("bind devices")

	return
}
