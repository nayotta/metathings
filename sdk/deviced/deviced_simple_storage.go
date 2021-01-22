package metathings_deviced_sdk

import (
	"context"
	"io"
	"io/ioutil"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/objx"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb "github.com/nayotta/metathings/proto/deviced"
)

type DevicedSimpleStorage struct {
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (s *DevicedSimpleStorage) get_logger() log.FieldLogger {
	return s.logger
}

func (s *DevicedSimpleStorage) toOpobject(x *pb.Object) *pb.OpObject {
	var y pb.OpObject
	buf, _ := grpc_helper.JSONPBMarshaler.MarshalToString(x)
	_ = grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(buf), &y)
	return &y
}

func (s *DevicedSimpleStorage) Put(ctx context.Context, obj *pb.Object, rd io.Reader) error {
	logger := s.get_logger()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}
	defer cfn()

	buf, err := ioutil.ReadAll(rd)
	if err != nil {
		logger.WithError(err).Debugf("failed to read content in reader")
		return err
	}

	req := &pb.PutObjectRequest{
		Object:  s.toOpobject(obj),
		Content: &wrappers.BytesValue{Value: buf},
	}

	_, err = cli.PutObject(ctx, req)
	if err != nil {
		logger.WithError(err).Debugf("failed to put object in deviced service")
		return err
	}

	return nil
}

func (s *DevicedSimpleStorage) Remove(ctx context.Context, obj *pb.Object) error {
	logger := s.get_logger()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}
	defer cfn()

	req := &pb.RemoveObjectRequest{
		Object: s.toOpobject(obj),
	}

	if _, err = cli.RemoveObject(ctx, req); err != nil {
		logger.WithError(err).Debugf("failed to remove object in deviced service")
		return err
	}

	return nil
}

func (s *DevicedSimpleStorage) Rename(ctx context.Context, src, dst *pb.Object) error {
	logger := s.get_logger()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}
	defer cfn()

	req := &pb.RenameObjectRequest{
		Source:      s.toOpobject(src),
		Destination: s.toOpobject(dst),
	}

	if _, err = cli.RenameObject(ctx, req); err != nil {
		logger.WithError(err).Debugf("failed to rename object in deviced service")
		return err
	}

	return nil
}

func (s *DevicedSimpleStorage) Get(ctx context.Context, obj *pb.Object) (*pb.Object, error) {
	logger := s.get_logger()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return nil, err
	}
	defer cfn()

	req := &pb.GetObjectRequest{
		Object: s.toOpobject(obj),
	}

	res, err := cli.GetObject(ctx, req)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object in deviced service")
		return nil, err
	}

	return res.GetObject(), nil
}

func (s *DevicedSimpleStorage) GetContent(ctx context.Context, obj *pb.Object) ([]byte, error) {
	logger := s.get_logger()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return nil, err
	}
	defer cfn()

	req := &pb.GetObjectContentRequest{
		Object: s.toOpobject(obj),
	}

	res, err := cli.GetObjectContent(ctx, req)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object content in deviced service")
		return nil, err
	}

	return res.GetContent(), nil
}

func (s *DevicedSimpleStorage) List(ctx context.Context, obj *pb.Object, opts ...SimpleStorageListOption) ([]*pb.Object, error) {
	logger := s.get_logger()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return nil, err
	}
	defer cfn()

	o := make(map[string]interface{})
	for _, opt := range opts {
		opt(o)
	}
	ox := objx.New(o)

	req := &pb.ListObjectsRequest{
		Object:    s.toOpobject(obj),
		Recursive: &wrappers.BoolValue{Value: ox.Get("recursive").Bool()},
		Depth:     &wrappers.Int32Value{Value: ox.Get("depth").Int32()},
	}

	res, err := cli.ListObjects(ctx, req)
	if err != nil {
		logger.WithError(err).Debugf("failed to list objects in deviced service")
		return nil, err
	}

	return res.GetObjects(), nil
}

func NewDevicedSimpleStorage(args ...interface{}) (SimpleStorage, error) {
	var logger log.FieldLogger
	var cli_fty *client_helper.ClientFactory

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":         opt_helper.ToLogger(&logger),
		"client_factory": client_helper.ToClientFactory(&cli_fty),
	})(args...); err != nil {
		return nil, err
	}

	return &DevicedSimpleStorage{
		logger:  logger,
		cli_fty: cli_fty,
	}, nil
}

func init() {
	register_simple_storage_factory("default", NewDevicedSimpleStorage)
}
