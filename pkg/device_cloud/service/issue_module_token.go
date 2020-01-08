package metathings_device_cloud_service

import (
	"net/http"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"

	device_service "github.com/nayotta/metathings/pkg/device/service"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (s *MetathingsDeviceCloudService) IssueModuleToken(w http.ResponseWriter, r *http.Request) {
	req := new(device_pb.IssueModuleTokenRequest)
	err := ParseHttpRequestBody(r, req)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ts, err := ptypes.Timestamp(req.GetTimestamp())
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to parse timestamp")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cred_id := req.GetCredential().GetId().GetValue()
	nonce := req.GetNonce().GetValue()
	hmac := req.GetHmac().GetValue()

	cli, cfn, err := s.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to connect identityd2 service")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cfn()

	tkn, err := device_service.IssueModuleTokenWithClient(cli, r.Context(), cred_id, ts, nonce, hmac)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to issue module token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := &device_pb.IssueModuleTokenResponse{
		Token: tkn,
	}

	buf, err := ParseHttpResponseBody(res)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(buf))

	s.get_logger().WithFields(log.Fields{
		"module": tkn.Entity.Id,
	}).Debugf("issue module token")
}
