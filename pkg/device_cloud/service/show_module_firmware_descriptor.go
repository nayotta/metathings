package metathings_device_cloud_service

import (
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *MetathingsDeviceCloudService) ShowModuleFirmwareDescriptor(w http.ResponseWriter, r *http.Request) {
	tkn_txt := GetTokenFromHeader(r)
	tkn, err := s.tkvdr.Validate(r.Context(), tkn_txt)
	logger := s.get_logger()

	if err != nil {
		logger.WithError(err).Errorf("failed to validate token")
		w.WriteHeader(http.StatusUnauthorized)
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

	logger = logger.WithField("device", dev.Id)

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to new deviced serivce client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cfn()

	res, err := cli.ShowDeviceFirmwareDescriptor(s.context_with_device(dev.Id), &empty.Empty{})
	if err != nil {
		logger.WithError(err).Errorf("failed to get device firmware descriptor")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf, err := ParseHttpResponseBody(res)
	if err != nil {
		logger.WithError(err).Errorf("failed to parse http response body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(buf))

	logger.Infof("show module firmware")

	return
}
