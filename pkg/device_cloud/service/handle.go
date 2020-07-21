package metathings_device_cloud_service

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	cpt "github.com/nayotta/metathings/pkg/component"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (dc *DeviceConnection) handle(req *pb.ConnectRequest) error {
	switch req.Kind {
	case pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM:
		return dc.handle_system_request(req)
	case pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER:
		return dc.handle_user_request(req)
	default:
		dc.logger.Warningf("unexpected request data")
		return nil
	}
}

func (dc *DeviceConnection) handle_system_request(req *pb.ConnectRequest) error {
	switch req.Union.(type) {
	case *pb.ConnectRequest_UnaryCall:
		return dc.handle_system_unary_request(req)
	case *pb.ConnectRequest_StreamCall:
		return dc.handle_system_stream_request(req)
	default:
		dc.logger.WithField("union", req.Union).Debugf("unsupported union type")
		return nil
	}
}

func (dc *DeviceConnection) handle_system_unary_request(req *pb.ConnectRequest) error {
	req_val := req.GetUnaryCall()
	sess := req.GetSessionId().GetValue()
	component := req_val.GetComponent().GetValue()
	name := req_val.GetName().GetValue()
	method := req_val.GetMethod().GetValue()

	logger := dc.logger.WithFields(log.Fields{
		"#session":   sess,
		"#component": component,
		"#name":      name,
		"#method":    method,
	})

	req_sign := fmt.Sprintf("%v$%v$%v", component, name, method)

	switch req_sign {
	case "system$system$pong":
		logger.Debugf("receive pong response")
		return nil
	case "system$system$sync_firmware":
		return dc.handle_system_unary_request_sync_firmware(req)
	default:
		dc.logger.WithField("sign", req_sign).Warningf("unsupported request sign")
		return nil
	}
}

func (dc *DeviceConnection) handle_system_unary_request_sync_firmware(req *pb.ConnectRequest) error {
	return dc.sync_firmware()
}

func (dc *DeviceConnection) handle_system_stream_request(req *pb.ConnectRequest) error {
	panic("unimplemented")
}

func (dc *DeviceConnection) handle_user_request(req *pb.ConnectRequest) error {
	switch req.Union.(type) {
	case *pb.ConnectRequest_UnaryCall:
		return dc.handle_user_unary_request(req)
	case *pb.ConnectRequest_StreamCall:
		return dc.handle_user_stream_request(req)
	default:
		dc.logger.WithField("union", req.Union).Debugf("unsupported union type")
		return nil
	}
}

func (dc *DeviceConnection) handle_user_unary_request(req *pb.ConnectRequest) error {
	req_val := req.GetUnaryCall()
	sess := req.GetSessionId().GetValue()
	kind := req.GetKind()
	component := req_val.GetComponent().GetValue()
	name := req_val.GetName().GetValue()
	method := req_val.GetMethod().GetValue()

	logger := dc.logger.WithFields(log.Fields{
		"#session":   sess,
		"#component": component,
		"#name":      name,
		"#method":    method,
	})

	mdl_prx, err := dc.get_module_proxy(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to get module proxy")
		return err
	}
	defer mdl_prx.Close()

	res_any, err := mdl_prx.UnaryCall(context.TODO(), method, req_val.GetValue())
	var res *pb.ConnectResponse

	if err != nil {
		logger.WithError(err).Debugf("failed to uanry call in module proxy")
		s := status.Convert(err)
		res = &pb.ConnectResponse{
			SessionId: sess,
			Kind:      kind,
			Union: &pb.ConnectResponse_Err{
				Err: &pb.ErrorValue{
					Name:      name,
					Component: component,
					Method:    method,
					Code:      uint32(s.Code()),
					Message:   s.Message(),
				},
			},
		}
	} else {
		res = &pb.ConnectResponse{
			SessionId: sess,
			Kind:      kind,
			Union: &pb.ConnectResponse_UnaryCall{
				UnaryCall: &pb.UnaryCallValue{
					Name:      name,
					Component: component,
					Method:    method,
					Value:     res_any,
				},
			},
		}
	}

	err = dc.stm.Send(res)
	if err != nil {
		logger.WithError(err).Debugf("failed to send msg")
		return err
	}

	logger.Debugf("send msg")

	return nil
}

func (dc *DeviceConnection) handle_user_stream_request(req *pb.ConnectRequest) error {
	var cli pb.DevicedServiceClient
	var cfn client_helper.CloseFn
	var stm pb.DevicedService_ConnectClient
	var err error

	req_val := req.GetStreamCall()
	req_sess := req.GetSessionId().GetValue()
	cfg := req_val.GetConfig()
	name := cfg.GetName().GetValue()
	component := cfg.GetComponent().GetValue()
	method := cfg.GetMethod().GetValue()

	logger := dc.logger.WithFields(log.Fields{
		"#session":   req_sess,
		"#component": component,
		"#name":      name,
		"#method":    method,
	})

	if cli, cfn, err = dc.cli_fty.NewDevicedServiceClient(); err != nil {
		logger.WithError(err).Debugf("failed to new deviced service client")
		return err
	}
	defer cfn()
	logger.Debugf("new deviced service client")

	ctx := context_helper.NewOutgoingContext(
		context.Background(),
		context_helper.WithTokenOp(dc.tknr.GetToken()),
		context_helper.WithSessionOp(req_sess),
		context_helper.WithDeviceOp(dc.opt.Device.Id),
	)
	if stm, err = cli.Connect(ctx); err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}
	logger.Debugf("create deviced stream")

	acked := make(chan struct{})
	acked_once := new(sync.Once)
	go func() {
		for cnt := 0; cnt < dc.opt.Config.MaxSendConfigResponseRetry; cnt++ {
			select {
			case _, ok := <-acked:
				if !ok {
					return
				}
			default:
			}

			cfg_res := &pb.ConnectResponse{
				SessionId: req_sess,
				Kind:      req.GetKind(),
				Union: &pb.ConnectResponse_StreamCall{
					StreamCall: &pb.StreamCallValue{
						Union: &pb.StreamCallValue_Config{
							Config: &pb.StreamCallConfig{
								Name:      name,
								Component: component,
								Method:    method,
							},
						},
					},
				},
			}

			if err = stm.Send(cfg_res); err != nil {
				logger.WithError(err).Debugf("failed to send config response")
				return
			}
			logger.Debugf("send config response")
			time.Sleep(time.Duration(math.Min((float64(cnt)*dc.opt.Config.SendConfigResponseIntervalA)+dc.opt.Config.SendConfigResponseIntervalB, dc.opt.Config.MaxSendConfigResponseInterval)) * time.Millisecond)
		}
		select {
		case _, ok := <-acked:
			if !ok {
				return
			}
		default:
			logger.Warningf("failed to recv config ack")
		}
	}()

	mdl_prx, err := dc.get_module_proxy(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to stream call")
		return err
	}

	stm = grpc_helper.NewHijackStream(stm, func(stm_ *grpc_helper.HijackStream, req *pb.ConnectRequest) error {
		switch req.GetStreamCall().Union.(type) {
		case *pb.OpStreamCallValue_Value:
			stm_.RecvChan() <- req
		case *pb.OpStreamCallValue_Config:
		case *pb.OpStreamCallValue_ConfigAck:
			logger.Debugf("recv config ack")
			acked_once.Do(func() { close(acked) })
		case *pb.OpStreamCallValue_Exit:
			logger.Debugf("recv exit")
			return context.Canceled
		}

		return nil
	})

	err = mdl_prx.StreamCall(context.Background(), method, cpt.NewModuleProxyStream(stm, req_sess))
	if err != nil {
		logger.WithError(err).Debugf("failed to stream call")
		return err
	}
	logger.Debugf("stream closed")

	return nil
}
