package metathings_echo_service

import (
	"context"
	"time"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	agentd_pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type options struct {
	id              string
	name            string
	logLevel        string
	agentdAddr      string
	metathingsdAddr string
	endpoint        string

	heartbeat_interval int64
}

var (
	default_options = options{
		heartbeat_interval: 15,
	}
)

type ServiceOptions func(*options)

func SetName(name string) ServiceOptions {
	return func(o *options) {
		o.name = name
	}
}

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
	}
}

func SetAgentdAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.agentdAddr = addr
	}
}

func SetMetathingsdAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.metathingsdAddr = addr
	}
}

func SetEndpoint(ep string) ServiceOptions {
	return func(o *options) {
		o.endpoint = ep
	}
}

type metathingsEchoService struct {
	opts    options
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (srv *metathingsEchoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	var text_str string
	text := req.GetText()
	if text != nil {
		text_str = text.Value
		srv.logger.Infof("echo: %v", text_str)
		return &pb.EchoResponse{Text: text_str}, nil
	}
	return nil, grpc.Errorf(codes.InvalidArgument, "empty body")
}

func (srv *metathingsEchoService) ConnectToAgent() error {
	ctx := context.Background()

	cli, closeFn, err := srv.cli_fty.NewCoreAgentServiceClient()
	if err != nil {
		return err
	}
	defer closeFn()

	req := &agentd_pb.CreateOrGetEntityRequest{
		Name:        &gpb.StringValue{Value: srv.opts.name},
		ServiceName: &gpb.StringValue{Value: "echo"},
		Endpoint:    &gpb.StringValue{Value: srv.opts.endpoint},
	}
	res, err := cli.CreateOrGetEntity(ctx, req)
	if err != nil {
		return err
	}
	srv.opts.id = res.Entity.Id

	errs := make(chan error)
	go func() {
		for {
			req := &agentd_pb.HeartbeatRequest{
				EntityId: &gpb.StringValue{Value: srv.opts.id},
			}

			_, err := cli.Heartbeat(ctx, req)
			if err != nil {
				errs <- err
				return
			}

			srv.logger.Debugf("Heartbeat")
			<-time.After(time.Duration(srv.opts.heartbeat_interval) * time.Second)
		}
	}()

	return <-errs
}

func NewEchoService(opt ...ServiceOptions) (*metathingsEchoService, error) {
	opts := default_options
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("echod", opts.logLevel)
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.metathingsdAddr)
	cli_fty_cfgs[client_helper.AGENTD_CONFIG] = client_helper.ServiceConfig{opts.agentdAddr}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	return &metathingsEchoService{
		opts:    opts,
		logger:  logger,
		cli_fty: cli_fty,
	}, nil
}
