package metathings_deviced_sdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	pb "github.com/nayotta/metathings/proto/deviced"
)

type DevicedCaller struct {
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (c *DevicedCaller) get_logger() log.FieldLogger {
	return c.logger
}

func (c *DevicedCaller) UnaryCall(ctx context.Context, device, module, method string, arguments map[string]interface{}) (map[string]interface{}, error) {
	logger := c.get_logger().WithFields(log.Fields{
		"device": device,
		"module": module,
		"method": method,
	})

	cli, cfn, err := c.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to new deviced service client")
		return nil, err
	}
	defer cfn()

	lcbd_req := &pb.ListConfigsByDeviceRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{
				Value: device,
			},
		},
	}
	lcbd_res, err := cli.ListConfigsByDevice(ctx, lcbd_req)
	if err != nil {
		logger.WithError(err).Debugf("failed to list configs by device in deviced")
		return nil, err
	}

	desc_cfg, err := LookupConfig(lcbd_res.GetConfigs(), deviced_helper.DEVICE_CONFIG_DESCRIPTOR)
	if err != nil {
		logger.WithError(err).Debugf("failed to lookup config")
		return nil, err
	}

	// TODO(Peer): cache descriptor
	sha1 := desc_cfg.Get(fmt.Sprintf("modules.%v.sha1", module)).String()
	if sha1 == "" {
		err = ErrModuleNotFound
		logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	gd_req := &pb.GetDescriptorRequest{
		Descriptor_: &pb.OpDescriptor{
			Sha1: &wrappers.StringValue{
				Value: sha1,
			},
		},
	}
	gd_res, err := cli.GetDescriptor(ctx, gd_req)
	if err != nil {
		logger.WithError(err).Debugf("failed to get descritpor")
		return nil, err
	}

	desc_buf := gd_res.GetDescriptor_().GetBody()
	var fds dpb.FileDescriptorSet
	if err = proto.Unmarshal(desc_buf, &fds); err != nil {
		logger.WithError(err).Debugf("failed to unmarshal descriptor")
		return nil, err
	}

	fd, err := desc.CreateFileDescriptorFromSet(&fds)
	if err != nil {
		logger.WithError(err).Debugf("failed to create file descritpor")
		return nil, err
	}

	var md *desc.MethodDescriptor
	var req_msg *dynamic.Message

	for _, sd := range fd.GetServices() {
		md = sd.FindMethodByName(method)
		if md == nil {
			continue
		}

		req_msg = dynamic.NewMessage(md.GetInputType())
	}

	if md == nil {
		err = ErrMethodNotFound
		logger.WithError(err).Debugf("failed to get method in descriptor")
		return nil, err
	}

	req_buf, err := json.Marshal(arguments)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal unary call arguments to json")
		return nil, err
	}

	err = req_msg.UnmarshalJSON(req_buf)
	if err != nil {
		logger.WithError(err).Debugf("failed to unmarshal json string to request message")
		return nil, err
	}

	any_msg, err := ptypes.MarshalAny(req_msg)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal request message to any message")
		return nil, err
	}

	uc_req := &pb.UnaryCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device},
		},
		Value: &pb.OpUnaryCallValue{
			Name:   &wrappers.StringValue{Value: module},
			Method: &wrappers.StringValue{Value: method},
			Value:  any_msg,
		},
	}
	uc_res, err := cli.UnaryCall(ctx, uc_req)
	if err != nil {
		logger.WithError(err).Debugf("failed to unary call in deviced")
		return nil, err
	}

	any_msg = uc_res.GetValue().GetValue()
	res_msg := dynamic.NewMessage(md.GetOutputType())
	if err = ptypes.UnmarshalAny(any_msg, res_msg); err != nil {
		logger.WithError(err).Debugf("failed to unmarshal any message to method output message")
		return nil, err
	}

	res_buf, err := new(jsonpb.Marshaler).MarshalToString(res_msg)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal output message to json string")
		return nil, err
	}

	res := make(map[string]interface{})
	err = json.Unmarshal([]byte(res_buf), &res)
	if err != nil {
		logger.WithError(err).Debugf("failed to unmarshal json string to map")
		return nil, err
	}

	return res, nil
}

func NewDevicedCaller(args ...interface{}) (Caller, error) {
	var logger log.FieldLogger
	var cli_fty *client_helper.ClientFactory

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":         opt_helper.ToLogger(&logger),
		"client_factory": client_helper.ToClientFactory(&cli_fty),
	})(args...); err != nil {
		return nil, err
	}

	return &DevicedCaller{
		logger:  logger,
		cli_fty: cli_fty,
	}, nil
}

func init() {
	register_caller_factory("default", NewDevicedCaller)
}
