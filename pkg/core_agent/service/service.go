package meatathings_core_agent_service

import (
	"context"
	"os"
	"path"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	app_cred_mgr "github.com/bigdatagz/metathings/pkg/common/application_credential_manager"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	cored_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
)

type options struct {
	metathings_addr string
	logLevel        string

	core_agent_home               string
	core_id                       string
	application_credential_id     string
	application_credential_secret string
}

var defaultServiceOptions = options{
	logLevel: "info",
}

type ServiceOptions func(*options)

func SetMetathingsAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.metathings_addr = addr
	}
}

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
	}
}

func SetCoreId(id string) ServiceOptions {
	return func(o *options) {
		o.core_id = id
	}
}

func SetApplicationCredential(id, secret string) ServiceOptions {
	return func(o *options) {
		o.application_credential_id = id
		o.application_credential_secret = secret
	}
}

type coreAgentService struct {
	appCredMgr *app_cred_mgr.ApplicationCredentialManager
	logger     log.FieldLogger
	opts       options
}

func (srv *coreAgentService) CreateEntity(context.Context, *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) DeleteEntity(context.Context, *pb.DeleteEntityRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) PatchEntity(context.Context, *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) GetEntity(context.Context, *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) ListEntities(context.Context, *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func getCoreIdFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data := make([]byte, 64)
	_, err = f.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getCoreIdFromService(opts options, token string) (string, error) {
	ctx := context.Background()
	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	grpc_opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(opts.metathings_addr, grpc_opts...)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cli := cored_pb.NewCoreServiceClient(conn)
	req := &cored_pb.CreateCoreRequest{}
	res, err := cli.CreateCore(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Core.Id, nil
}

func saveCoreIdToPath(id, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(id))
	if err != nil {
		return err
	}

	return nil
}

func NewCoreAgentService(opt ...ServiceOptions) (*coreAgentService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("core-agent", opts.logLevel)
	if err != nil {
		log.WithField("error", err).Errorf("failed to new logger")
		return nil, err
	}

	appCredMgr, err := app_cred_mgr.NewApplicationCredentialManager(
		opts.metathings_addr,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.WithField("error", err).Errorf("failed to new application credential manager")
		return nil, err
	}

	if opts.core_id == "" {
		core_id_path := path.Join(opts.core_agent_home, "core-id")
		core_id, err := getCoreIdFromFile(core_id_path)
		if err != nil {
			core_id, err = getCoreIdFromService(opts, appCredMgr.GetToken())
			if err != nil {
				return nil, err
			}
			saveCoreIdToPath(core_id, core_id_path)
		}
		opts.core_id = core_id

	}

	srv := &coreAgentService{
		appCredMgr: appCredMgr,
		logger:     logger,
		opts:       opts,
	}
	return srv, nil
}
