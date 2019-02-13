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
	air_switch_pb "github.com/nayotta/metathings/pkg/proto/esp32_air_switch"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	MtpURL             = "metathings.ai:21733"
	MtpUsername        = "admin"
	MtpPassword        = "admin"
	MtpDefaultDomainID = "default"
	MtpDeviceID        = "airswitch01"
)

var (
	testGetSwitchsAddr        int32 = 0x0009 //switch 0 and 3
	testSwitchAddr0           int32
	testSwitchType0           int32 = 0x0007
	testSwitchModel0          int32 = 0x0040
	testSwitchVotageHigh      int32 = 0x0104 //0x0104 = 260
	testSwitchVotageLow       int32 = 0x0064 //0x0064 = 100
	testSwitchLeakCurrentHigh int32 = 0x012c //0x012c = 300
	testSwitchPowerHigh       int32 = 0x3700 //0x3700 = 14080
	testSwitchTempHigh        int32 = 0x0320 //0x0320 = 800
	testSwitchCurrentHigh     int32 = 0x2580 //0x2580 = 9600
	testSwitchWarnVotageHigh  int32 = 0x00F0 //0x00F0 = 240
	testSwitchWarnVotageLow   int32 = 0x00C8 //0x00C8 = 200
	testSwitchKwh             int64 = 0x5432
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
				Value: "airswitch01",
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
			var resMqtt air_switch_pb.MqttDeviceResponse
			//fmt.Println("begin recv")
			res, err := stream.Recv()
			if err != nil {
				fmt.Printf("recv error:%v\n", err)
				return
			}

			resMqtt = air_switch_pb.MqttDeviceResponse{}
			proto.Unmarshal(res.GetValue().GetValue().GetValue(), &resMqtt)
			switch resMqtt.Payload.(type) {
			case *air_switch_pb.MqttDeviceResponse_GetSwitchCtrlRes:
				fmt.Println("CTRL-->", resMqtt.GetGetSwitchCtrlRes(), "<--")
				break
			case *air_switch_pb.MqttDeviceResponse_GetSwitchDataRes:
				fmt.Println("DATA-->", resMqtt.GetGetSwitchDataRes(), "<--")
				break
			case *air_switch_pb.MqttDeviceResponse_GetSwitchStateRes:
				fmt.Println("STAT-->", resMqtt.GetGetSwitchStateRes(), "<--")
				break
			case *air_switch_pb.MqttDeviceResponse_GetSwitchWarnRes:
				fmt.Println("WARN-->", resMqtt.GetGetSwitchWarnRes(), "<--")
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
		payload := &air_switch_pb.MqttDeviceRequest{
			Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchStateReq{
				GetSwitchStateReq: &air_switch_pb.GetSwitchStateReq{
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
					Value: "airswitch01",
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
