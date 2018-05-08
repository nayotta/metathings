package metathings_echo_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	cs_helper "github.com/bigdatagz/metathings/pkg/common/core_service"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type metathingsEchoService struct {
	mt_plugin.CoreService
	opts    cs_helper.Options
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

func NewEchoService(opts cs_helper.Options) (*metathingsEchoService, error) {
	opts.Set("service_name", "echo")

	logger, err := log_helper.NewLogger("echod", opts.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENTD_CONFIG] = client_helper.ServiceConfig{opts.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	srv := &metathingsEchoService{
		opts:    opts,
		logger:  logger,
		cli_fty: cli_fty,
	}
	srv.CoreService = mt_plugin.MakeCoreService(srv.opts, srv.cli_fty)

	return srv, nil
}
