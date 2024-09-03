package metathings_device_cloud_service

import (
	"net/http"

	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
	pb "github.com/nayotta/metathings/proto/device"
)

func (s *MetathingsDeviceCloudService) CreateOrUpdateProfile(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger().WithField("#method", "CreateOrUpdateProfile")

	ctx := s.context()
	tkn_txt := GetTokenFromHeader(r)
	_, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req pb.CreateOrUpdateProfileRequest
	if err = ParseHttpRequestBody(r, &req); err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	profile := req.GetProfile()

	if err = s.profile_storage.CreateOrUpdateProfile(ctx, intf.Profile{
		Name:    profile.GetName().GetValue(),
		Driver:  profile.GetDriver().GetValue(),
		Address: profile.GetAddress().GetValue(),
		Port:    profile.GetPort().GetValue(),
	}); err != nil {
		logger.WithError(err).Errorf("failed to create or update profile in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	logger.Infof("create or update profile")

	return
}
