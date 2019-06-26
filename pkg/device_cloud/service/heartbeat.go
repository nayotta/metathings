package metathings_device_cloud_service

import (
	"encoding/json"
	"net/http"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type HeartbeatRequest struct {
	Module pb.OpModule
}

func (s *MetathingsDeviceCloudService) Heartbeat(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	req := new(HeartbeatRequest)
	err := dec.Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO(Peer): Get module id from token not request
	mdl_id := req.Module.GetId().GetValue()

	err = s.storage.Heartbeat(mdl_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	go s.try_to_build_device_connection_by_module_id(mdl_id)

	w.WriteHeader(http.StatusNoContent)
}
