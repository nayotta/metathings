package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

func unary_gpio_digital_write(cli pb.MqttdServiceClient, ctx context.Context, req *any.Any) (*any.Any, error) {
	req1 := &pb.GpioDigitalWriteRequest{}
	err := ptypes.UnmarshalAny(req, req1)
	if err != nil {
		return nil, err
	}

	res, err := cli.GpioDigitalWrite(ctx, req1)
	if err != nil {
		return nil, err
	}

	res1, err := ptypes.MarshalAny(res)
	if err != nil {
		return nil, err
	}

	return res1, nil

}

func init() {
	unaryCallMethods["gpio_digital_write"] = unary_gpio_digital_write
}
