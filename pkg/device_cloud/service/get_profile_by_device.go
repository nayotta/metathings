package metathings_device_cloud_service

import (
	"net/http"

	pb "github.com/nayotta/metathings/proto/device"
)

func (s *MetathingsDeviceCloudService) GetProfileByDevice(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger().WithField("#method", "GetProfileByDevice")

	ctx := s.context()
	tkn_txt := GetTokenFromHeader(r)
	_, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req pb.GetProfileByDeviceRequest
	if err = ParseHttpRequestBody(r, &req); err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	device := req.GetDevice()
	deviceId := device.GetId().GetValue()
	logger = logger.WithField("device.id", deviceId)

	p, err := s.profile_storage.GetProfileByDevice(ctx, deviceId)
	if err != nil {
		logger.WithError(err).Errorf("failed to get profile by device in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resProfile := ProfileToPb(p)
	res := &pb.GetProfileByDeviceResponse{
		Profile: &resProfile,
	}

	buf, err := ParseHttpResponseBody(res)
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(buf))
	logger.Debugf("get profile by device")

	return
}
