package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	ctx_help "github.com/nayotta/metathings/pkg/common/context"
	kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	switch_pb "github.com/nayotta/metathings/pkg/proto/esp8266_01_switch"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// MtpURL MtpURL
var (
	MtpURL             = "metathings.ai:21733"
	MtpUsername        = "admin"
	MtpPassword        = "admin"
	MtpDefaultDomainID = "default"
	MtpDeviceID        = "test"
)

var (
	testGetSwitchsAddr        int32 = 0x0008 //switch  3
	testSwitchAddr0           int32 = 0x0003 //switch  3
)

func main() {
	var tokenStr string

	conn, err := grpc.Dial(MtpURL, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                2 * time.Second,
		Timeout:             120 * time.Second,
		PermitWithoutStream: true,
	}))

	if err != nil {
		fmt.Printf("dial error:%v\n", err)
		return
	}

	identityd2Client := identityd2_pb.NewIdentitydServiceClient(conn)
	devicedClient := deviced_pb.NewDevicedServiceClient(conn)

	// issue token
	tokenIn := &identityd2_pb.IssueTokenByPasswordRequest{
		Entity: &identityd2_pb.OpEntity{
			Name: &wrappers.StringValue{
				Value: MtpUsername,
			},
			Password: &wrappers.StringValue{
				Value: MtpPassword,
			},
		},
	}
	tokenIn.Entity.Domains = append(tokenIn.Entity.Domains, &identityd2_pb.OpDomain{
		Id: &wrappers.StringValue{
			Value: MtpDefaultDomainID,
		},
	})

	tokenRes, err := identityd2Client.IssueTokenByPassword(context.Background(), tokenIn)

	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}

	// token with mt
	tokenStr = "mt " + tokenRes.GetToken().GetText()
	fmt.Println("issue token res:", tokenStr)

	// steamcall config
	ctx := context.Background()
	ctx = ctx_help.WithToken(ctx, tokenStr)
	stream, err := devicedClient.StreamCall(ctx)
	if err != nil {
		fmt.Println("stream create error", err)
		return
	}
	req := &deviced_pb.StreamCallRequest{
		Device: &deviced_pb.OpDevice{
			Id: &wrappers.StringValue{
				Value: MtpDeviceID,
			},
			Kind: kind.DeviceKind_DEVICE_KIND_SIMPLE,
		},
	}

	err = stream.Send(req)
	if err != nil {
		fmt.Println("stream config error", err)
		return
	}
	println("test setup")
	go func() {
		for {
			var res *deviced_pb.StreamCallResponse
			var resMqtt switch_pb.MqttDeviceResponse
			//fmt.Println("begin recv")
			res, err := stream.Recv()
			if err != nil {
				fmt.Printf("recv error:%v\n", err)
				return
			}

			resMqtt = switch_pb.MqttDeviceResponse{}
			proto.Unmarshal(res.GetValue().GetValue().GetValue(), &resMqtt)
			switch resMqtt.Payload.(type) {
			case *switch_pb.MqttDeviceResponse_GetSwitchStateRes:
				fmt.Println("STAT-->", resMqtt.GetGetSwitchStateRes(), "<--")
				break
			default:
				fmt.Println("UNKNOWN type")
				break
			}
		}
	}()

	var i int32
	//var cmd int32
	for i = 0; i < 10000; i++ {
		//fmt.Println("begin sent")
		/*
			if i%2 == 0 {
				cmd = 0xFF00
			} else {
				cmd = 0x0000
			}*/
		payload := &switch_pb.MqttDeviceRequest{
			Payload: &switch_pb.MqttDeviceRequest_GetSwitchStateReq{
				GetSwitchStateReq: &switch_pb.GetSwitchStateReq{
					SwitchAddr: testSwitchAddr0,
				},
			},
		}
		payloadByte, err := proto.Marshal(payload)
		if err != nil {
			fmt.Printf("proto marshal error:%v\n", err)
		}
		req := &deviced_pb.StreamCallRequest{
			Device: &deviced_pb.OpDevice{
				Id: &wrappers.StringValue{
					Value: MtpDeviceID,
				},
				Kind: kind.DeviceKind_DEVICE_KIND_SIMPLE,
			},
			Value: &deviced_pb.OpStreamCallValue{
				Union: &deviced_pb.OpStreamCallValue_Value{
					Value: &any.Any{
						Value: payloadByte,
					},
				},
			},
		}

		err = stream.Send(req)
		if err != nil {
			fmt.Println("stream send error", err)
			return
		}
		//time.Sleep(1 * time.Second)
		time.Sleep(20000000 * time.Millisecond)
	}
}
