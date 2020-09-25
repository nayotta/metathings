package metathings_device_cloud_service

import (
	"context"
	"net/http"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	device_pb "github.com/nayotta/metathings/proto/device"
	log "github.com/sirupsen/logrus"
)

func (s *MetathingsDeviceCloudService) PushFrameToFlow(w http.ResponseWriter, r *http.Request) {
	logger := s.get_logger()

	tkn_txt := GetTokenFromHeader(r)
	tkn, err := s.tkvdr.Validate(context.TODO(), tkn_txt)
	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req := new(device_pb.PushFrameToFlowRequest)
	err = ParseHttpRequestBody(r, req)
	if err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mdl_id := tkn.Entity.Id
	logger = logger.WithField("module", mdl_id)

	dev, err := s.get_device_by_module_id(r.Context(), mdl_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get device by module id")
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
		logger.Errorf("failed to find flow in device")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sess := id_helper.NewId()
	logger = logger.WithFields(log.Fields{
		"flow":    flw_n,
		"device":  dev.Id,
		"session": sess,
	})

	err = s.start_push_frame_loop(dev.Id, req, sess)
	if err != nil {
		logger.WithError(err).Errorf("failed to start push frame loop")
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
		logger.WithError(err).Errorf("failed to marshal response to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
	return
}
