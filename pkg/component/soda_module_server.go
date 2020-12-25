package metathings_component

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	component_pb "github.com/nayotta/metathings/proto/component"
)

type SodaModuleServer struct {
	m *Module

	grpc_srv      *grpc.Server
	grpc_srv_quit chan struct{}

	backend SodaModuleBackend
}

func (s *SodaModuleServer) start_grpc() error {
	logger := s.m.Logger()

	cfg := s.m.Kernel().Config()
	host := cfg.GetString("service.host")
	port := cfg.GetInt("service.port")

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return err
	}

	s.grpc_srv = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_helper.UnaryServerInterceptor(s.m.Logger())),
		grpc.StreamInterceptor(grpc_helper.StreamServerInterceptor(s.m.Logger())),
	)

	wrap, err := NewSodaModuleWrapper(s.m)
	if err != nil {
		return err
	}

	component_pb.RegisterModuleServiceServer(s.grpc_srv, wrap)

	go func() {
		if err = s.grpc_srv.Serve(lis); err != nil {
			logger.WithError(err).Warningf("grpc server closed")
		}
		close(s.grpc_srv_quit)
	}()

	return nil
}

func (s *SodaModuleServer) stop_grpc() error {
	s.grpc_srv.Stop()

	return nil
}

func (s *SodaModuleServer) wait_grpc() <-chan struct{} {
	return s.grpc_srv_quit
}

func (s *SodaModuleServer) Start() error {
	var err error

	if err = s.start_grpc(); err != nil {
		return err
	}

	if err = s.backend.Start(); err != nil {
		return err
	}

	return nil
}

func (s *SodaModuleServer) Wait() {
	select {
	case <-s.wait_grpc():
	case <-s.backend.Done():
	}
}

func (s *SodaModuleServer) Serve() error {
	var err error

	if err = s.Start(); err != nil {
		return err
	}

	s.Wait()

	return nil

}

func (s *SodaModuleServer) Stop() {
	s.stop_grpc()
	s.backend.Stop()
}

func NewSodaModuleServer(m *Module) (ModuleServer, error) {
	cfg := m.Kernel().Config()

	be, err := NewSodaModuleBackend(cfg.GetString("backend.name"), m)
	if err != nil {
		return nil, err
	}

	return &SodaModuleServer{
		m:             m,
		grpc_srv_quit: make(chan struct{}),
		backend:       be,
	}, nil
}

func init() {
	register_module_server_factory("soda", NewSodaModuleServer)

}
