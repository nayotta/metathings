package metathings_device_cloud_service

import (
	"net/http"

	pb "github.com/nayotta/metathings/proto/device"
)

func (s *MetathingsDeviceCloudService) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger().WithField("#method", "DeleteProfile")

	ctx := s.context()
	tkn_txt := GetTokenFromHeader(r)
	_, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req pb.DeleteProfileRequest
	if err = ParseHttpRequestBody(r, &req); err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	profile := req.GetProfile()

	if err = s.profile_storage.DeleteProfile(ctx, profile.GetName().GetValue()); err != nil {
		logger.WithError(err).Errorf("failed to delete profile in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Infof("delete profile")
	w.WriteHeader(http.StatusNoContent)

	return
}
