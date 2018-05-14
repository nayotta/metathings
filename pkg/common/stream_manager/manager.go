package stream_manager

import (
	"errors"
	"io"
	"sync"
	"time"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	helper "github.com/nayotta/metathings/pkg/common"
	cored_pb "github.com/nayotta/metathings/pkg/proto/core"
)

var (
	NotFound   = errors.New("not found")
	Timeout    = errors.New("timeout")
	Registered = errors.New("registered")
)

type StreamManager interface {
	Register(core_id string, stream cored_pb.CoreService_StreamServer) (chan interface{}, error)
	UnaryCall(core_id string, req *cored_pb.UnaryCallRequestPayload) (*cored_pb.UnaryCallResponsePayload, error)
}

type streamManager struct {
	streams  map[string]cored_pb.CoreService_StreamServer
	sessions map[string]chan *cored_pb.StreamResponse

	lock   *sync.Mutex
	logger log.FieldLogger
}

func (mgr *streamManager) Register(core_id string, stream cored_pb.CoreService_StreamServer) (chan interface{}, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	if _, ok := mgr.streams[core_id]; ok {
		return nil, Registered
	}
	mgr.streams[core_id] = stream

	close := make(chan interface{})
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				mgr.lock.Lock()
				defer mgr.lock.Unlock()
				delete(mgr.streams, core_id)
				close <- nil
				if err == io.EOF {
					mgr.logger.
						WithField("core_id", core_id).
						Infof("core agent stream closed")
				} else {
					mgr.logger.
						WithError(err).
						WithField("core_id", core_id).
						Warningf("core agent stream closed with unexpected error")
				}
				return
			}

			if ch, ok := mgr.sessions[res.SessionId]; !ok {
				mgr.logger.
					WithField("session_id", res.SessionId).
					Errorf("unknown session id")
			} else {
				ch <- res
			}
		}
	}()

	return close, nil
}

func (mgr *streamManager) UnaryCall(core_id string, req *cored_pb.UnaryCallRequestPayload) (*cored_pb.UnaryCallResponsePayload, error) {
	stream, ok := mgr.streams[core_id]
	if !ok {
		mgr.logger.WithField("core_id", core_id).Warningf("core not found")
		return nil, NotFound
	}

	sess_id := helper.NewId()
	ch := make(chan *cored_pb.StreamResponse)
	mgr.sessions[sess_id] = ch

	stm_req := &cored_pb.StreamRequest{
		SessionId:   &gpb.StringValue{Value: sess_id},
		MessageType: cored_pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER,
		Payload:     &cored_pb.StreamRequest_UnaryCall{req},
	}

	if err := stream.Send(stm_req); err != nil {
		mgr.logger.WithError(err).Errorf("failed to send unary call")
		return nil, err
	}

	defer func() {
		close(ch)
		delete(mgr.sessions, sess_id)
		mgr.logger.WithField("session_id", sess_id).Debugf("close session receive channel")
	}()
	select {
	case stm_res := <-ch:
		res := stm_res.Payload.(*cored_pb.StreamResponse_UnaryCall).UnaryCall
		return res, nil
	case <-time.After(30 * time.Second):
		return nil, Timeout
	}
}

func NewStreamManager(logger log.FieldLogger) (StreamManager, error) {
	return &streamManager{
		streams:  make(map[string]cored_pb.CoreService_StreamServer),
		sessions: make(map[string]chan *cored_pb.StreamResponse),
		lock:     new(sync.Mutex),
		logger:   logger,
	}, nil
}
