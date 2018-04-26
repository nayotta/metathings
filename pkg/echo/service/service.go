package metathings_echo_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type metathingsEchoService struct {
	logger log.FieldLogger
}

func (srv *metathingsEchoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	var text_str string
	text := req.GetText()
	if text != nil {
		text_str = text.Value
		srv.logger.Infof("echo: %v", text_str)
		return &pb.EchoResponse{text_str}, nil
	}
	return nil, grpc.Errorf(codes.InvalidArgument, "empty body")
}

func NewEchoService(logger log.FieldLogger) *metathingsEchoService {
	return &metathingsEchoService{
		logger: logger,
	}
}
