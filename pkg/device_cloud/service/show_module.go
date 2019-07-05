package metathings_device_cloud_service

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
)

func (s *MetathingsDeviceCloudService) ShowModule(w http.ResponseWriter, r *http.Request) {
	tkn_txt := GetTokenFromHeader(r)
	tkn, err := s.tkvdr.Validate(tkn_txt)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	mdl_id := tkn.Entity.Id

	dev, err := s.get_device_by_module_id(mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to get device by module id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, mdl := range dev.Modules {
		if mdl.Id == mdl_id {
			codec := jsonpb.Marshaler{}
			buf, err := codec.MarshalToString(mdl)
			if err != nil {
				s.get_logger().WithError(err).Errorf("failed to marshal response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			buf = fmt.Sprintf(`{"module": %v}`, buf)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(buf))
			return
		}
	}
}
