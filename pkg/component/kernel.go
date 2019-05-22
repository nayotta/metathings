package metathings_component

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	const_helper "github.com/nayotta/metathings/pkg/common/constant"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	identityd2_contrib "github.com/nayotta/metathings/pkg/identityd2/contrib"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"github.com/nayotta/viper"
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
	v.SetEnvPrefix(METATHINGS_COMPONENT_PREFIX)
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

func (k *Kernel) Config() *KernelConfig {
	return k.cfg.Sub(k.cfg.GetString("stage"))
}

type NewKernelOption struct {
	Credential struct {
		Id     string
		Secret string
	}
	ServiceEndpoints map[string]string
	ConfigText       string
}

func new_client_factory_from_new_kernel_option(opt *NewKernelOption) (*client_helper.ClientFactory, error) {
	var err error
	var cli_fty *client_helper.ClientFactory

	addr, ok := opt.ServiceEndpoints["default"]
	if !ok {
		return nil, ErrDefaultAddressRequired
	}
	srv_cfgs := client_helper.NewDefaultServiceConfigs(addr)
	addr, ok = opt.ServiceEndpoints["device"]
	if !ok {
		return nil, ErrDeviceAddressRequired
	}
	srv_cfgs[client_helper.DEVICED_CONFIG] = client_helper.ServiceConfig{addr}
	cli_fty, err = client_helper.NewClientFactory(srv_cfgs, client_helper.DefaultDialOptionFn())
	if err != nil {
		return nil, err
	}
	return cli_fty, nil
}

func new_client_factory_from_kernel_config(cfg *KernelConfig) (*client_helper.ClientFactory, error) {
	var err error
	var cli_fty *client_helper.ClientFactory

	addr := cfg.GetString("service_endpoint.default.address")
	if addr == "" {
		return nil, ErrDefaultAddressRequired
	}
	srv_cfgs := client_helper.NewDefaultServiceConfigs(addr)
	addr = cfg.GetString("service_endpoint.device.address")
	if addr == "" {
		return nil, ErrDeviceAddressRequired
	}
	srv_cfgs[client_helper.DEVICE_CONFIG] = client_helper.ServiceConfig{addr}
	cli_fty, err = client_helper.NewClientFactory(srv_cfgs, client_helper.DefaultDialOptionFn())
	if err != nil {
		return nil, err
	}
	return cli_fty, nil
}

func _get_module_config_text(cli pb.DeviceServiceClient, ctx context.Context, mdl_id string) (string, error) {
	req := &device_pb.GetObjectContentRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: fmt.Sprintf("%s/sys/config", mdl_id)},
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

func _put_object(cli pb.DeviceServiceClient, ctx context.Context, name string, content io.Reader) error {
	buf, err := ioutil.ReadAll(content)
	if err != nil {
		return err
	}

	req := &device_pb.PutObjectRequest{
		Object: &deviced_pb.OpObject{
			Name: &wrappers.StringValue{Value: name},
		},
		Content: &wrappers.BytesValue{Value: buf},
	}

	_, err = cli.PutObject(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func _get_object(cli pb.DeviceServiceClient, ctx context.Context, name string) (*deviced_pb.Object, error) {
	req := &device_pb.GetObjectRequest{
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
	req := &device_pb.GetObjectContentRequest{
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
	req := &device_pb.RemoveObjectRequest{
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
	req := &device_pb.RenameObjectRequest{
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

func _heartbeat(cli pb.DeviceServiceClient, ctx context.Context, name string) error {
	req := &device_pb.HeartbeatRequest{
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

func new_issue_module_token_request(domain, id, secret string) *device_pb.IssueModuleTokenRequest {
	itbc_req := identityd2_contrib.NewIssueTokenByCredentialRequest(domain, id, secret)
	return &device_pb.IssueModuleTokenRequest{
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

		cfg_txt, err = _get_module_config_text(cli, ctx, mdl.Id)
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
