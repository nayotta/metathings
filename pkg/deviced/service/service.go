package metathings_deviced_service

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"

	afo_helper "github.com/nayotta/metathings/pkg/common/auth_func_overrider"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	session_storage "github.com/nayotta/metathings/pkg/deviced/session_storage"
	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	evaluatord_sdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type MetathingsDevicedServiceOption struct {
	Methods struct {
		PutObjectStreaming struct {
			Timeout            int64
			ChunkSize          int64
			ChunkPerRequest    int
			PullRequestRetry   int
			PullRequestTimeout int64
		}
		PullFrameFromFlow struct {
			AliveInterval int64
		}
		PullFrameFromFlowSet struct {
			AliveInterval int64
		}
	}
}

type MetathingsDevicedService struct {
	grpc_auth.ServiceAuthFuncOverride
	tknr            token_helper.Tokener
	cli_fty         *client_helper.ClientFactory
	opt             *MetathingsDevicedServiceOption
	logger          log.FieldLogger
	storage         storage.Storage
	session_storage session_storage.SessionStorage
	simple_storage  simple_storage.SimpleStorage
	authorizer      identityd_authorizer.Authorizer
	validator       identityd_validator.Validator
	tkvdr           token_helper.TokenValidator
	cc              connection.ConnectionCenter
	flw_fty         flow.FlowFactory
	flwst_fty       flow.FlowSetFactory
	data_launcher   evaluatord_sdk.DataLauncher
}

func (self *MetathingsDevicedService) get_device_by_context(ctx context.Context) (*storage.Device, error) {
	var tkn *identityd_pb.Token
	var dev_s *storage.Device
	var err error

	// try to get device id from context
	dev_id := metautils.ExtractIncoming(ctx).Get("MT-Device")
	if dev_id == "" {
		tkn = context_helper.ExtractToken(ctx)
		dev_id = tkn.Entity.Id
	}

	if dev_s, err = self.storage.GetDevice(ctx, dev_id); err != nil {
		return nil, err
	}

	return dev_s, nil
}

func (self *MetathingsDevicedService) new_flow(dev_id, flw_id string) (flow.Flow, error) {
	return self.flw_fty.New(&flow.FlowOption{
		FlowId:   flw_id,
		DeviceId: dev_id,
	})
}

func (self *MetathingsDevicedService) new_flow_sets(flwsts_s []*storage.FlowSet) ([]flow.FlowSet, error) {
	var flwsts []flow.FlowSet

	for _, flwst_s := range flwsts_s {
		if flwst, err := self.flwst_fty.New(&flow.FlowSetOption{
			FlowSetId: *flwst_s.Id,
		}); err != nil {
			for _, flwst := range flwsts {
				flwst.Close()
			}

			return nil, err
		} else {
			flwsts = append(flwsts, flwst)
		}
	}

	return flwsts, nil
}

func (self *MetathingsDevicedService) get_flow_ids_by_devices(ctx context.Context, devs []*pb.OpDevice) ([]string, error) {
	flw_ids := []string{}

	for _, dev_r := range devs {
		dev_id := dev_r.GetId().GetValue()
		dev_s, err := self.storage.GetDevice(ctx, dev_id)
		if err != nil {
			return nil, err
		}

		for _, flw_r := range dev_r.Flows {
			ok := false
			for _, flw_s := range dev_s.Flows {
				if flw_r.GetName().GetValue() == *flw_s.Name {
					flw_ids = append(flw_ids, *flw_s.Id)
					ok = true
					break
				}
			}
			if !ok {
				return nil, ErrFlowNotFound
			}
		}
	}

	return flw_ids, nil
}

func (self *MetathingsDevicedService) offline_device(ctx context.Context, dev_id string) (err error) {
	var dev_s *storage.Device
	defer func() {
		if err != nil {
			self.logger.WithField("device", dev_id).WithError(err).Debugf("failed to offline device")
		} else {
			self.logger.WithField("device", dev_id).Debugf("device offline")
		}

	}()

	dev_s, err = self.storage.GetDevice(ctx, dev_id)
	if err != nil {
		return err
	}

	state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(state_pb.DeviceState_DEVICE_STATE_OFFLINE)
	_, err = self.storage.PatchDevice(ctx, dev_id, &storage.Device{
		State: &state_str,
	})
	if err != nil {
		return err
	}

	state_str = deviced_helper.MODULE_STATE_ENUMER.ToString(state_pb.ModuleState_MODULE_STATE_OFFLINE)
	for _, mdl_s := range dev_s.Modules {
		_, err = self.storage.PatchModule(ctx, *mdl_s.Id, &storage.Module{
			State: &state_str,
		})
		if err != nil {
			return err
		}
	}

	return
}

func (self *MetathingsDevicedService) IsIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	return false
}

func NewMetathingsDevicedService(
	opt *MetathingsDevicedServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	session_storage session_storage.SessionStorage,
	simple_storage simple_storage.SimpleStorage,
	authorizer identityd_authorizer.Authorizer,
	validator identityd_validator.Validator,
	tkvdr token_helper.TokenValidator,
	cc connection.ConnectionCenter,
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
	flw_fty flow.FlowFactory,
	flwst_fty flow.FlowSetFactory,
	data_launcher evaluatord_sdk.DataLauncher,
) (pb.DevicedServiceServer, error) {
	srv := &MetathingsDevicedService{
		opt:             opt,
		logger:          logger,
		storage:         storage,
		session_storage: session_storage,
		simple_storage:  simple_storage,
		authorizer:      authorizer,
		validator:       validator,
		tkvdr:           tkvdr,
		cc:              cc,
		tknr:            tknr,
		cli_fty:         cli_fty,
		flw_fty:         flw_fty,
		flwst_fty:       flwst_fty,
		data_launcher:   data_launcher,
	}
	srv.ServiceAuthFuncOverride = afo_helper.NewAuthFuncOverrider(tkvdr, srv, logger)

	return srv, nil
}
