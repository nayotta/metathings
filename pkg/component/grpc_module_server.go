package metathings_component

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	component_pb "github.com/nayotta/metathings/pkg/proto/component"
)

type GrpcModuleServer struct {
	m   *Module
	srv *grpc.Server
}

func (s *GrpcModuleServer) Serve() error {
	cfg := s.m.Kernel().Config()
	host := cfg.GetString("service.host")
	port := cfg.GetInt("service.port")

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return err
	}

	s.srv = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_helper.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_helper.StreamServerInterceptor()),
	)

	component_pb.RegisterModuleServiceServer(s.srv, NewGrpcModuleWrapper(s.m.Target(), s.m.Logger()))

	return s.srv.Serve(lis)
}

func (s *GrpcModuleServer) Stop() {
	s.srv.Stop()
}

func NewGrpcModuleServer(m *Module) (ModuleServer, error) {
	return &GrpcModuleServer{
		m: m,
	}, nil
}

func init() {
	register_module_server_factory("grpc", NewGrpcModuleServer)
}
