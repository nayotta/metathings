package metathings_device_cloud_service

import (
	"context"
	"fmt"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
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
	default:
		dc.logger.WithField("sign", req_sign).Warningf("unsupported request sign")
		return nil
	}
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
	panic("unimplemented")
}
