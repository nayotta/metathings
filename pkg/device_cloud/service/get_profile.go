package metathings_device_cloud_service

import (
	"net/http"

	pb "github.com/nayotta/metathings/proto/device"
)

func (s *MetathingsDeviceCloudService) GetProfile(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger().WithField("#method", "GetProfile")

	ctx := s.context()
	tkn_txt := GetTokenFromHeader(r)
	_, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req pb.GetProfileRequest
	if err = ParseHttpRequestBody(r, &req); err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	profile := req.GetProfile()
	profileName := profile.GetName().GetValue()
	logger = logger.WithField("profile.name", profileName)

	p, err := s.profile_storage.GetProfile(ctx, profileName)
	if err != nil {
		logger.WithError(err).Errorf("failed to get profile in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resProfile := ProfileToPb(p)
	res := &pb.GetProfileResponse{
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
	logger.Debugf("get profile")

	return
}
