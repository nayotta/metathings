package metathings_component_ext_firmware_service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nayotta/metathings/pkg/common/binary_synchronizer"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	component "github.com/nayotta/metathings/pkg/component"
)

type ComponentExtFirmwareService struct {
	module *component.Module
	bs     binary_synchronizer.BinarySynchronizer

	synchronizing_firmware_mtx   sync.Mutex
	stats_synchronizing_firmware bool
}

func (svc *ComponentExtFirmwareService) is_synchronizing_firmware() bool {
	return svc.stats_synchronizing_firmware
}

func (svc *ComponentExtFirmwareService) HANDLE_GRPC_SyncFirmware(ctx context.Context, in *any.Any) (*any.Any, error) {
	var err error

	if _, err = svc.SyncFirmware(ctx, &empty.Empty{}); err != nil {
		return nil, err
	}

	out, err := ptypes.MarshalAny(&empty.Empty{})
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (svc *ComponentExtFirmwareService) SyncFirmware(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	logger := svc.module.Logger()

	mdl_name := svc.module.Name()

	desc, err := svc.module.Kernel().ShowFirmwareDescriptor()
	if err != nil {
		logger.WithError(err).Errorf("failed to show firmware descriptor")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	desc_str, err := new(jsonpb.Marshaler).MarshalToString(desc.GetDescriptor_())
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal descriptor to json string")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	descx, err := objx.FromJSON(desc_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal json string to objx")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	mdl_nxt_ver := descx.Get(fmt.Sprintf("modules.%s.version.next", mdl_name)).String()
	if mdl_nxt_ver != svc.module.GetVersion() {
		uri := descx.Get(fmt.Sprintf("modules.%s.uri.next", mdl_name)).String()
		sha256sum := descx.Get(fmt.Sprintf("modules.%s.sha256.next", mdl_name)).String()
		go svc.do_sync_firmware(uri, sha256sum)
	}

	return &empty.Empty{}, nil
}

func (svc *ComponentExtFirmwareService) do_sync_firmware(uri, sha256sum string) error {
	logger := svc.module.Logger().WithFields(logrus.Fields{
		"uri":       uri,
		"sha256sum": sha256sum,
	})

	if svc.is_synchronizing_firmware() {
		logger.Wanringf("synchronizing firmware, please wait a minutes")
		return nil
	}

	svc.synchronizing_firmware_mtx.Lock()
	svc.stats_synchronizing_firmware = true
	defer func() {
		svc.stats_synchronizing_firmware = false
		svc.synchronizing_firmware_mtx.Unlock()
	}()

	src, err := os.Executable()
	if err != nil {
		logger.WithError(err).Debugf("failed to get executable info")
		return err
	}

	var postfix string
	if len(sha256sum) > 12 {
		postfix = sha256sum[:12]
	} else {
		postfix = fmt.Sprintf("%d", time.Now().Unix())
	}

	dst := fmt.Sprintf("%v.%v", filepath.Base(src), postfix)
	if err = svc.bs.Sync(svc.module.Kernel().Context(), src, dst, uri, sha256sum); err != nil {
		logger.WithError(err).Debugf("failed to sync firmware")
		return err
	}

	logger.Infof("sync firmware, please restart!")

	return nil
}

func NewComponentExtFirmwareService(m *component.Module) (*ComponentExtFirmwareService, error) {
	cfg := cast.ToStringMap(m.Kernel().Config().Get("binary_synchronizer"))
	args := cfg_helper.FlattenConfigOption(cfg, "logger", m.Logger())
	bs, err := binary_synchronizer.NewBinarySynchronizer(args...)
	if err != nil {
		return nil, err
	}

	srv := &ComponentExtFirmwareService{
		module: m,
		bs:     bs,

		stats_synchronizing_firmware: false,
	}

	return srv, nil
}
