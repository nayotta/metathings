package metathings_component

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/wrapperspb"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	const_helper "github.com/nayotta/metathings/pkg/common/constant"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	identityd2_contrib "github.com/nayotta/metathings/pkg/identityd2/contrib"
	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
	identityd2_pb "github.com/nayotta/metathings/proto/identityd2"
)

type KernelConfig struct {
	*viper.Viper
}

func (kc *KernelConfig) Sub(key string) *KernelConfig {
	return &KernelConfig{kc.Viper.Sub(key)}
}

func (kc *KernelConfig) Raw() *viper.Viper {
	return kc.Viper
}

func new_kernel_config_viper_form_text(text string) (*viper.Viper, error) {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvPrefix(constant_helper.PREFIX_METATHINGS_COMPONENT)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	v.SetConfigType("yaml")
	err := v.ReadConfig(strings.NewReader(text))
	if err != nil {
		return nil, err
	}

	return v, nil
}

func NewKernelConfigFromText(text string) (*KernelConfig, error) {
	v, err := new_kernel_config_viper_form_text(text)
	if err != nil {
		return nil, err
	}

	return &KernelConfig{v}, nil
}

type KernelInterface interface {
	Context() context.Context
	Show() (*deviced_pb.Module, error)
	ShowFirmwareDescriptor() (*deviced_pb.FirmwareDescriptor, error)
	PutObject(name string, content io.Reader) error
	PutObjectStreaming(name string, content io.ReadSeeker, opt *PutObjectStreamingOption) error
	PutObjectStreamingWithCancel(name string, content io.ReadSeeker, opt *PutObjectStreamingOption) (context.CancelFunc, chan error, error)
	PutObjects(objects map[string]io.Reader) error
	GetObject(name string) (*deviced_pb.Object, error)
	GetObjectContent(name string) ([]byte, error)
	RemoveObjct(name string) error
	RemoveObjets(names []string) error
	RenameObject(src, dst string) error
	PushFrameToFlowOnce(name string, data interface{}, opt *PushFrameToFlowOnceOption) error
	Heartbeat() error
}

type Kernel struct {
	cli_fty *client_helper.ClientFactory
	cfg     *KernelConfig
	logger  log.FieldLogger
	ctx     context.Context
}

func (k *Kernel) Context() context.Context {
	if k.ctx == nil {
		cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
		if err != nil {
			panic(err)
		}
		defer cfn()

		kc := k.Config()
		id := kc.GetString("credential.id")
		secret := kc.GetString("credential.secret")
		tkn, err := _issue_module_token(cli, context.TODO(), id, secret)
		if err != nil {
			panic(err)
		}

		k.ctx = context_helper.WithToken(context.TODO(), tkn.GetText())
	}

	return k.ctx
}

func (k *Kernel) Show() (*deviced_pb.Module, error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	ctx := k.Context()
	mdl, err := _show_module(cli, ctx)
	if err != nil {
		return nil, err
	}

	return mdl, nil
}

func (k *Kernel) ShowFirmwareDescriptor() (*deviced_pb.FirmwareDescriptor, error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	ctx := k.Context()
	desc, err := _show_module_firmware_descriptor(cli, ctx)
	if err != nil {
		return nil, err
	}

	return desc, nil
}

func (k *Kernel) PutObject(name string, content io.Reader) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = _put_object(cli, k.Context(), name, content)
	if err != nil {
		return err
	}

	return nil
}

type PutObjectStreamingOption struct {
	Sha1   string
	Length int64
}

func (k *Kernel) PutObjectStreaming(name string, content io.ReadSeeker, opt *PutObjectStreamingOption) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = _put_object_streaming(cli, k.Context(), name, content, opt.Sha1, opt.Length)
	if err != nil {
		return err
	}

	return nil
}

func (k *Kernel) PutObjectStreamingWithCancel(name string, content io.ReadSeeker, opt *PutObjectStreamingOption) (cancel context.CancelFunc, errs chan error, err error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, nil, err
	}
	defer cfn()

	ctx := k.Context()
	ctx, cancel = context.WithCancel(ctx)

	errs = make(chan error, 1)
	go func() {
		defer close(errs)
		errs <- _put_object_streaming(cli, ctx, name, content, opt.Sha1, opt.Length)
	}()

	return cancel, errs, nil
}

func (k *Kernel) PutObjects(objects map[string]io.Reader) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	for name, content := range objects {
		if err = _put_object(cli, k.Context(), name, content); err != nil {
			return err
		}
	}

	return nil
}

func (k *Kernel) GetObject(name string) (*deviced_pb.Object, error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	obj, err := _get_object(cli, k.Context(), name)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (k *Kernel) GetObjectContent(name string) ([]byte, error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	content, err := _get_object_content(cli, k.Context(), name)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (k *Kernel) RemoveObject(name string) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = _remove_object(cli, k.Context(), name)
	if err != nil {
		return err
	}

	return nil
}

func (k *Kernel) RemoveObjects(names []string) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	for _, name := range names {
		if err = _remove_object(cli, k.Context(), name); err != nil {
			return err
		}
	}

	return nil
}

func (k *Kernel) RenameObject(src, dst string) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = _rename_object(cli, k.Context(), src, dst)
	if err != nil {
		return err
	}

	return nil
}

type ListObjectsOption struct {
	Recursive bool
	Depth     int32
}

func (k *Kernel) ListObjects(name string, opt *ListObjectsOption) ([]*deviced_pb.Object, error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	objs, err := _list_objects(cli, k.Context(), name, opt.Recursive, opt.Depth)
	if err != nil {
		return nil, err
	}

	return objs, nil
}

type PushFrameToFlowOnceOption struct {
	Id *string
	Ts *time.Time
}

func (k *Kernel) PushFrameToFlowOnce(name string, data interface{}, opt *PushFrameToFlowOnceOption) error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	id := id_helper.NewId()
	ts := time.Now()

	if opt == nil {
		opt = &PushFrameToFlowOnceOption{}
	}

	if opt.Id != nil {
		id = *opt.Id
	}

	if opt.Ts != nil {
		ts = *opt.Ts
	}

	if err = _push_frame_to_flow_once(cli, k.Context(), id, name, ts, data); err != nil {
		return err
	}

	return nil
}

func (k *Kernel) Heartbeat() error {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	kc := k.Config()
	name := kc.GetString("name")

	err = _heartbeat(cli, k.Context(), name)
	if err != nil {
		return err
	}

	return nil
}

type FrameStream struct {
	stream   pb.DeviceService_PushFrameToFlowClient
	push_ack bool
	close_cb func() error
}

func (fs *FrameStream) PushFrame(frm *deviced_pb.OpFrame) error {
	req_id := id_helper.NewId()
	req := &pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: req_id},
		Request: &pb.PushFrameToFlowRequest_Frame{
			Frame: frm,
		},
	}

	err := fs.stream.Send(req)
	if err != nil {
		return err
	}

	if fs.push_ack {
		res, err := fs.stream.Recv()
		if err != nil {
			return err
		}

		if res.GetId() != req_id || res.GetAck() == nil {
			return ErrUnexceptedResponse
		}
	}

	return nil
}

func (fs *FrameStream) Push(dat interface{}) error {
	var dat_js string
	var err error

	switch msg := dat.(type) {
	case proto.Message:
		dat_js, err = grpc_helper.JSONPBMarshaler.MarshalToString(msg)
		if err != nil {
			return err
		}
	case map[string]interface{}:
		buf, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		dat_js = string(buf)
	}

	var st stpb.Struct
	err = grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(dat_js), &st)
	if err != nil {
		return err
	}

	frm := &deviced_pb.OpFrame{
		Ts:   ptypes.TimestampNow(),
		Data: &st,
	}

	return fs.PushFrame(frm)
}

func (fs *FrameStream) Close() error {
	return fs.close_cb()
}

func (k *Kernel) NewFrameStream(flow string) (*FrameStream, error) {
	cli, cfn, err := k.cli_fty.NewDeviceServiceClient()
	if err != nil {
		return nil, err
	}

	kc := k.Config()
	config_ack := kc.GetBool("flow.config_ack")
	push_ack := kc.GetBool("flow.push_ack")

	stm, err := _build_push_frame_to_flow_stream(cli, k.Context(), flow, config_ack, push_ack)
	if err != nil {
		defer cfn()
		return nil, err
	}

	fs := &FrameStream{
		stream:   stm,
		push_ack: push_ack,
		close_cb: func() error {
			cfn()
			return nil
		},
	}

	return fs, nil
}

func (k *Kernel) Config() *KernelConfig {
	return k.cfg.Sub(k.cfg.GetString("stage"))
}

type NewKernelOption struct {
	Credential struct {
		Id     string
		Secret string
	}
	TransportCredential TransportCredential
	ServiceEndpoints    map[string]ServiceEndpoint
	ConfigText          string
}

func new_service_config_from_service_endpoint(ep ServiceEndpoint) (client_helper.ServiceConfig, error) {
	cred, err := client_helper.NewClientTransportCredentials(ep.CertFile, ep.KeyFile, ep.PlainText, ep.Insecure)
	if err != nil {
		return client_helper.ServiceConfig{}, err
	}

	return client_helper.ServiceConfig{
		Address:              ep.Address,
		TransportCredentials: cred,
	}, nil
}

func new_client_factory_from_new_kernel_option(opt *NewKernelOption) (*client_helper.ClientFactory, error) {
	var err error
	var cli_fty *client_helper.ClientFactory

	ep, ok := opt.ServiceEndpoints["default"]
	if !ok {
		return nil, ErrDefaultAddressRequired
	}
	srv_cfgs := client_helper.ServiceConfigs{}
	srv_cfgs[client_helper.DEFAULT_CONFIG], err = new_service_config_from_service_endpoint(ep)
	if err != nil {
		return nil, err
	}

	ep, ok = opt.ServiceEndpoints["device"]
	if !ok {
		return nil, ErrDeviceAddressRequired
	}
	srv_cfgs[client_helper.DEVICE_CONFIG], err = new_service_config_from_service_endpoint(ep)
	if err != nil {
		return nil, err
	}

	cli_fty, err = client_helper.NewClientFactory(srv_cfgs, client_helper.DefaultDialOption())
	if err != nil {
		return nil, err
	}
	return cli_fty, nil
}

func new_service_endpoint_from_kernel_config(cfg *KernelConfig, name string) (ServiceEndpoint, error) {
	var err error
	var srv_ep ServiceEndpoint

	if err = cfg.Sub("service_endpoint." + name).Unmarshal(&srv_ep); err != nil {
		return srv_ep, err
	}

	return srv_ep, nil
}

func new_client_factory_from_kernel_config(cfg *KernelConfig) (*client_helper.ClientFactory, error) {
	var err error

	opt := NewKernelOption{
		ServiceEndpoints: map[string]ServiceEndpoint{},
	}
	opt.ServiceEndpoints["default"], err = new_service_endpoint_from_kernel_config(cfg, "default")
	if err != nil {
		return nil, err
	}
	opt.ServiceEndpoints["device"], err = new_service_endpoint_from_kernel_config(cfg, "device")
	if err != nil {
		return nil, err
	}

	return new_client_factory_from_new_kernel_option(&opt)
}

func _get_module_config_text(cli pb.DeviceServiceClient, ctx context.Context, mdl_name string) (string, error) {
	req := &pb.GetObjectContentRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: fmt.Sprintf("%s/sys/config", mdl_name)},
		},
	}
	res, err := cli.GetObjectContent(ctx, req)
	if err != nil {
		return "", err
	}

	return string(res.GetContent()), nil
}

func _show_module(cli pb.DeviceServiceClient, ctx context.Context) (*deviced_pb.Module, error) {
	res, err := cli.ShowModule(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res.GetModule(), nil
}

func _show_module_firmware_descriptor(cli pb.DeviceServiceClient, ctx context.Context) (*deviced_pb.FirmwareDescriptor, error) {
	res, err := cli.ShowModuleFirmwareDescriptor(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res.GetFirmwareDescriptor(), nil
}

func _put_object(cli pb.DeviceServiceClient, ctx context.Context, name string, content io.Reader) error {
	buf, err := ioutil.ReadAll(content)
	if err != nil {
		return err
	}

	req := &pb.PutObjectRequest{
		Object: &deviced_pb.OpObject{
			Name:   &wrappers.StringValue{Value: name},
			Length: &wrapperspb.Int64Value{Value: int64(len(buf))},
		},
		Content: &wrappers.BytesValue{Value: buf},
	}

	_, err = cli.PutObject(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func _put_object_streaming(cli pb.DeviceServiceClient, ctx context.Context, name string, content io.ReadSeeker, sha1 string, length int64) error {
	stm, err := cli.PutObjectStreaming(ctx)
	if err != nil {
		return err
	}

	req := &pb.PutObjectStreamingRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Request: &pb.PutObjectStreamingRequest_Metadata_{
			Metadata: &pb.PutObjectStreamingRequest_Metadata{
				Object: &deviced_pb.OpObject{
					Name:   &wrappers.StringValue{Value: name},
					Length: &wrappers.Int64Value{Value: length},
				},
				Sha1: &wrappers.StringValue{Value: sha1},
			},
		},
	}

	if err = stm.Send(req); err != nil {
		return err
	}

	errs := make(chan error)
	defer close(errs)
	go _put_object_streaming_loop(stm, content, errs)
	if err = <-errs; err != nil {
		return err
	}

	return nil
}

func _put_object_streaming_loop(stm pb.DeviceService_PutObjectStreamingClient, content io.ReadSeeker, errs chan error) {
	for {
		res, err := stm.Recv()
		if err != nil {
			errs <- err
			return
		}

		chunks := res.GetChunks()
		if chunks == nil {
			continue
		}

		req := &pb.PutObjectStreamingRequest{
			Id: &wrappers.StringValue{Value: res.GetId()},
		}
		req_chks := []*deviced_pb.OpObjectChunk{}
		for _, chk := range chunks.GetChunks() {
			offset := chk.GetOffset()
			length := chk.GetLength()
			buf := make([]byte, length)

			if _, err = content.Seek(offset, 0); err != nil {
				errs <- err
				return
			}

			n, err := content.Read(buf)
			if err != nil {
				errs <- err
				return
			}
			req_chks = append(req_chks, &deviced_pb.OpObjectChunk{
				Offset: &wrappers.Int64Value{Value: offset},
				Data:   &wrappers.BytesValue{Value: buf[:n]},
				Length: &wrappers.Int64Value{Value: int64(n)},
			})
		}
		req.Request = &pb.PutObjectStreamingRequest_Chunks{
			Chunks: &deviced_pb.OpObjectChunks{
				Chunks: req_chks,
			},
		}

		if err = stm.Send(req); err != nil {
			errs <- err
			return
		}
	}
}

func _get_object(cli pb.DeviceServiceClient, ctx context.Context, name string) (*deviced_pb.Object, error) {
	req := &pb.GetObjectRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: name},
		},
	}

	res, err := cli.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.GetObject(), nil
}

func _get_object_content(cli pb.DeviceServiceClient, ctx context.Context, name string) ([]byte, error) {
	req := &pb.GetObjectContentRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: name},
		},
	}

	res, err := cli.GetObjectContent(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.GetContent(), nil
}

func _remove_object(cli pb.DeviceServiceClient, ctx context.Context, name string) error {
	req := &pb.RemoveObjectRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: name},
		},
	}

	_, err := cli.RemoveObject(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func _rename_object(cli pb.DeviceServiceClient, ctx context.Context, src, dst string) error {
	req := &pb.RenameObjectRequest{
		Source: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: src},
		},
		Destination: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: dst},
		},
	}

	_, err := cli.RenameObject(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func _list_objects(cli pb.DeviceServiceClient, ctx context.Context, name string, recursive bool, depth int32) ([]*deviced_pb.Object, error) {
	req := &pb.ListObjectsRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: name},
		},
		Recursive: &wrappers.BoolValue{Value: recursive},
		Depth:     &wrappers.Int32Value{Value: depth},
	}

	res, err := cli.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Objects, nil
}

func _push_frame_to_flow_once(cli pb.DeviceServiceClient, ctx context.Context, id string, name string, ts time.Time, data interface{}) error {
	pbts, err := ptypes.TimestampProto(ts)
	if err != nil {
		return err
	}

	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var pbst stpb.Struct
	err = grpc_helper.JSONPBUnmarshaler.Unmarshal(bytes.NewReader(buf), &pbst)
	if err != nil {
		return err
	}

	req := &pb.PushFrameToFlowOnceRequest{
		Id: &wrappers.StringValue{Value: id},
		Flow: &deviced_pb.OpFlow{
			Name: &wrappers.StringValue{Value: name},
		},
		Frame: &deviced_pb.OpFrame{
			Ts:   pbts,
			Data: &pbst,
		},
	}

	_, err = cli.PushFrameToFlowOnce(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func _heartbeat(cli pb.DeviceServiceClient, ctx context.Context, name string) error {
	req := &pb.HeartbeatRequest{
		Module: &deviced_pb.OpModule{
			Name: &wrappers.StringValue{Value: name},
		},
	}

	_, err := cli.Heartbeat(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func new_issue_module_token_request(domain, id, secret string) *pb.IssueModuleTokenRequest {
	itbc_req := identityd2_contrib.NewIssueTokenByCredentialRequest(domain, id, secret)
	return &pb.IssueModuleTokenRequest{
		Credential: itbc_req.GetCredential(),
		Timestamp:  itbc_req.GetTimestamp(),
		Nonce:      itbc_req.GetNonce(),
		Hmac:       itbc_req.GetHmac(),
	}
}

func _issue_module_token(cli pb.DeviceServiceClient, ctx context.Context, id, secret string) (*identityd2_pb.Token, error) {
	req := new_issue_module_token_request(const_helper.DEFAULT_DOMAIN, id, secret)
	res, err := cli.IssueModuleToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.GetToken(), nil
}

func _build_push_frame_to_flow_stream(cli pb.DeviceServiceClient, ctx context.Context, flow string, config_ack, push_ack bool) (pb.DeviceService_PushFrameToFlowClient, error) {
	stm, err := cli.PushFrameToFlow(ctx)
	if err != nil {
		return nil, err
	}

	cfg_req_id := id_helper.NewId()
	cfg_req := &pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: cfg_req_id},
		Request: &pb.PushFrameToFlowRequest_Config_{
			Config: &pb.PushFrameToFlowRequest_Config{
				Flow: &deviced_pb.OpFlow{
					Name: &wrappers.StringValue{Value: flow},
				},
				ConfigAck: &wrappers.BoolValue{Value: config_ack},
				PushAck:   &wrappers.BoolValue{Value: push_ack},
			},
		},
	}
	err = stm.Send(cfg_req)
	if err != nil {
		return nil, err
	}

	if config_ack {
		res, err := stm.Recv()
		if err != nil {
			return nil, err
		}

		if res.GetId() != cfg_req_id || res.GetAck() == nil {
			return nil, ErrUnexceptedResponse
		}
	}

	return stm, nil
}

func NewKernel(opt *NewKernelOption) (*Kernel, error) {
	var err error
	var cfg_txt string
	var cli_fty *client_helper.ClientFactory

	if opt.ConfigText == "" {
		cli_fty, err = new_client_factory_from_new_kernel_option(opt)
		if err != nil {
			return nil, err
		}

		cli, cfn, err := cli_fty.NewDeviceServiceClient()
		if err != nil {
			return nil, err
		}
		defer cfn()

		tkn, err := _issue_module_token(cli, context.TODO(), opt.Credential.Id, opt.Credential.Secret)
		if err != nil {
			return nil, err
		}

		ctx := context_helper.WithToken(context.TODO(), tkn.GetText())

		mdl, err := _show_module(cli, ctx)
		if err != nil {
			return nil, err
		}

		cfg_txt, err = _get_module_config_text(cli, ctx, mdl.Name)
		if err != nil {
			return nil, err
		}
	} else {
		cfg_txt = opt.ConfigText
	}
	cfg, err := NewKernelConfigFromText(cfg_txt)
	if err != nil {
		return nil, err
	}
	krn := &Kernel{cfg: cfg}

	kc := krn.Config()

	logger, err := log_helper.NewLogger("kernel", kc.GetString("log.level"))
	if err != nil {
		return nil, err
	}
	krn.logger = logger

	if cli_fty == nil {
		cli_fty, err = new_client_factory_from_kernel_config(kc)
	}
	krn.cli_fty = cli_fty

	return krn, nil
}
