package metathings_device_cloud_service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/protobuf/jsonpb"
	device_service "github.com/nayotta/metathings/pkg/device/service"
	log "github.com/sirupsen/logrus"
)

type IssueModuleTokenRequest struct {
	Credential struct {
		Id string
	}
	Timestamp string
	Nonce     int64
	Hmac      string
}

func (s *MetathingsDeviceCloudService) IssueModuleToken(w http.ResponseWriter, r *http.Request) {
	req := new(IssueModuleTokenRequest)
	err := ParseHttpRequestBody(r, req)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ts, err := time.Parse(time.RFC3339Nano, req.Timestamp)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to parse timestamp")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cli, cfn, err := s.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to connect identityd2 service")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cfn()

	tkn, err := device_service.IssueModuleTokenWithClient(cli, context.TODO(), req.Credential.Id, ts, req.Nonce, req.Hmac)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to issue module token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codec := jsonpb.Marshaler{}
	buf, err := codec.MarshalToString(tkn)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO(Peer): marshal by jsonpb.Marshaler not write by code
	buf = fmt.Sprintf(`{"token": %v}`, buf)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(buf))

	s.get_logger().WithFields(log.Fields{
		"module": tkn.Entity.Id,
	}).Debugf("issue module token")
}
