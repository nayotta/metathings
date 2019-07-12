package metathings_device_cloud_service

import (
	"encoding/json"
	"net/http"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
)

type PushFrameToFlowRequest struct {
	Id   string
	Flow struct {
		Name string
	}
	ConfigAck bool `json:"config_ack"`
	PushAck   bool `json:"push_ack"`
}

type PushFrameToFlowResponse struct {
	Id      string
	Session string
}

func (s *MetathingsDeviceCloudService) PushFrameToFlow(w http.ResponseWriter, r *http.Request) {
	tkn_txt := GetTokenFromHeader(r)
	tkn, err := s.tkvdr.Validate(tkn_txt)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req := new(PushFrameToFlowRequest)
	err = ParseHttpRequestBody(r, req)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mdl_id := tkn.Entity.Id

	dev, err := s.get_device_by_module_id(mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to get device by module id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	found := false
	for _, f := range dev.Flows {
		if f.Name == req.Flow.Name {
			found = true
			break
		}
	}

	if !found {
		s.get_logger().Errorf("failed to find flow in device")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sess := id_helper.NewId()
	go s.start_push_frame_loop(dev.Id, req, sess)

	res := PushFrameToFlowResponse{
		Id:      dev.Id,
		Session: sess,
	}

	buf, err := json.Marshal(&res)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to marshal response to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
	return
}
