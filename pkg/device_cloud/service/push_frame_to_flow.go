package metathings_device_cloud_service

import (
	"net/http"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (s *MetathingsDeviceCloudService) PushFrameToFlow(w http.ResponseWriter, r *http.Request) {
	tkn_txt := GetTokenFromHeader(r)
	tkn, err := s.tkvdr.Validate(tkn_txt)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req := new(device_pb.PushFrameToFlowRequest)
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

	flw_n := req.GetConfig().GetFlow().GetName().GetValue()

	found := false
	for _, f := range dev.Flows {
		if f.Name == flw_n {
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
	err = s.start_push_frame_loop(dev.Id, req, sess)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to start push frame loop")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := &device_pb.PushFrameToFlowResponse{
		Id: dev.Id,
		Response: &device_pb.PushFrameToFlowResponse_Config_{
			Config: &device_pb.PushFrameToFlowResponse_Config{
				Session: sess,
			},
		},
	}

	buf, err := ParseHttpResponseBody(res)
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
