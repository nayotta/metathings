package metathings_device_cloud_service

import (
	"net/http"

	device_pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (s *MetathingsDeviceCloudService) Heartbeat(w http.ResponseWriter, r *http.Request) {
	tkn_txt := GetTokenFromHeader(r)
	req_mdl_sess := GetSessionFromHeader(r)
	tkn, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	mdl_id := tkn.Entity.Id

	// TODO(Peer): match token module name with request body module name
	req := new(device_pb.HeartbeatRequest)
	err = ParseHttpRequestBody(r, req)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cur_mdl_sess, err := s.storage.GetModuleSession(mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to get module session")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cur_mdl_sess != 0 && cur_mdl_sess != req_mdl_sess {
		s.get_logger().Warningf("current module session not 0, maybe duplicated")
		w.WriteHeader(http.StatusConflict)
		return
	}

	err = s.storage.SetModuleSession(mdl_id, req_mdl_sess)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to set module session in storage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.storage.Heartbeat(mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to heartbeat in storage")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO(Peer): cache device data in device cloud
	dev, err := s.get_device_by_module_id(r.Context(), mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to get device by module id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go s.try_to_build_device_connection(dev)

	s.get_logger().Debugf("heartbeat")
	w.WriteHeader(http.StatusNoContent)
}
