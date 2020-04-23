package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pbst "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
)

var (
	srv_ep_opt              cmd_contrib.ServiceEndpointOption
	base_opt                cmd_contrib.BaseOption
	device                  string
	module_descriptor_slice []string
)

func main() {
	pflag.StringVar(&srv_ep_opt.Address, "addr", "", "Deviced Service Address")
	pflag.BoolVar(&srv_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&srv_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&srv_ep_opt.CertFile, "certfile", "", "Cert File for connect to Deviced")
	pflag.StringVar(&srv_ep_opt.KeyFile, "keyfile", "", "Key File for connect to Deviced")
	pflag.StringVar(&device, "device", "", "Device Id")
	pflag.StringSliceVar(&module_descriptor_slice, "module-descriptor", nil, "Module and descriptor pairs. ex: module:desc_file_path")

	pflag.Parse()

	base_opt.Token = os.Getenv("MT_TOKEN")

	logger := logrus.New()

	ctx := context_helper.WithToken(context.TODO(), "Bearer "+base_opt.Token)
	var opts []grpc.DialOption
	if srv_ep_opt.PlainText {
		opts = append(opts, grpc.WithInsecure())
	} else if srv_ep_opt.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	}

	conn, err := grpc.Dial(srv_ep_opt.Address, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewDevicedServiceClient(conn)

	lcbd_req := &pb.ListConfigsByDeviceRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device},
		},
	}
	lcbd_res, err := cli.ListConfigsByDevice(ctx, lcbd_req)
	if err != nil {
		panic(err)
	}

	cfg, err := dsdk.LookupConfig(lcbd_res.GetConfigs(), deviced_helper.DEVICE_CONFIG_DESCRIPTOR)
	if err != nil {
		if err != dsdk.ErrConfigNotFound {
			panic(err)
		}

		cc_req := &pb.CreateConfigRequest{
			Config: &pb.OpConfig{
				Alias: &wrappers.StringValue{Value: deviced_helper.DEVICE_CONFIG_DESCRIPTOR},
				Body:  &pbst.Struct{},
			},
		}

		cc_res, err := cli.CreateConfig(ctx, cc_req)
		if err != nil {
			panic(err)
		}
		cfg_id := cc_res.GetConfig().GetId()
		logger.WithField("config", cfg_id).Infof("create config")

		actd_req := &pb.AddConfigsToDeviceRequest{
			Device: &pb.OpDevice{
				Id: &wrappers.StringValue{Value: device},
			},
			Configs: []*pb.OpConfig{
				{Id: &wrappers.StringValue{Value: cfg_id}},
			},
		}
		_, err = cli.AddConfigsToDevice(ctx, actd_req)
		if err != nil {
			panic(err)
		}
		logger.WithFields(logrus.Fields{
			"config": cfg_id,
			"device": device,
		}).Infof("add config to device")

		cfg, err = dsdk.FromConfig(cc_res.GetConfig())
		if err != nil {
			panic(err)
		}
	}

	for _, md := range module_descriptor_slice {
		xs := strings.Split(md, ":")
		if len(xs) != 2 {
			panic("bad module descriptor pairs: " + md)
		}

		module := xs[0]
		descriptor := xs[1]

		buf, err := ioutil.ReadFile(descriptor)
		if err != nil {
			panic(err)
		}

		ud_req := &pb.UploadDescriptorRequest{
			Descriptor_: &pb.OpDescriptor{
				Body: &wrappers.BytesValue{Value: buf},
			},
		}

		ud_res, err := cli.UploadDescriptor(ctx, ud_req)
		if err != nil {
			panic(err)
		}
		sha1 := ud_res.GetDescriptor_().GetSha1()
		logger.WithField("descriptor", sha1).Infof("upload descriptor")

		cfg.Set(fmt.Sprintf("modules.%v.sha1", module), sha1)
	}

	cfg_pb, err := dsdk.ToProtobuf(cfg)
	if err != nil {
		panic(err)
	}
	pc_req := &pb.PatchConfigRequest{
		Config: cfg_pb,
	}
	_, err = cli.PatchConfig(ctx, pc_req)
	if err != nil {
		panic(err)
	}

	logger.Infof("setup device configs...")
}
