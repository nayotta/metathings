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
	username        = "admin"
	password        = "admin"
	defaultDomainID = "default"
)

func main() {
	var tokenStr string
	var ctx context.Context
	sepStr := "\n-----------------------------------------------------------"

	fmt.Println("dial")
	conn, err := grpc.Dial("metathings.ai:21733", grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                2 * time.Second,
		Timeout:             120 * time.Second,
		PermitWithoutStream: true,
	}))
	if err != nil {
		fmt.Printf("dial error:%v\n", err)
	}

	identityd2Client := identityd2_pb.NewIdentitydServiceClient(conn)
	devicedClient := deviced_pb.NewDevicedServiceClient(conn)

	// issue token
	tokenIn := &identityd2_pb.IssueTokenByPasswordRequest{
		Entity: &identityd2_pb.OpEntity{
			Name: &wrappers.StringValue{
				Value: username,
			},
			Password: &wrappers.StringValue{
				Value: password,
			},
		},
	}
	tokenIn.Entity.Domains = append(tokenIn.Entity.Domains, &identityd2_pb.OpDomain{
		Id: &wrappers.StringValue{
			Value: defaultDomainID,
		},
	})

	tokenRes, err := identityd2Client.IssueTokenByPassword(context.Background(), tokenIn)
	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}

	// token with mt
	tokenStr = "mt " + tokenRes.GetToken().GetText()
	fmt.Println("issue token res:", tokenStr, sepStr)

	// streamcall
	ctx = context.Background()
	ctx = ctx_help.WithToken(ctx, tokenStr)
	stream, err := devicedClient.StreamCall(ctx)

	go func() {
		defer fmt.Println("recv loop deinit")
		for {
			var res *deviced_pb.StreamCallResponse
			var resMqtt air_switch_pb.MqttDeviceResponse
			//fmt.Println("begin recv")
			res, err = stream.Recv()
			if err != nil {
				fmt.Printf("recv error:%v\n", err)
				return
			}
			//fmt.Println("recv msg", res.GetValue().GetValue().GetValue())
			resMqtt = air_switch_pb.MqttDeviceResponse{}
			proto.Unmarshal(res.GetValue().GetValue().GetValue(), &resMqtt)
			switch resMqtt.Payload.(type) {
			case *air_switch_pb.MqttDeviceResponse_Cmd_01Res:
				fmt.Println(resMqtt.GetCmd_01Res().RetCmd, resMqtt.GetCmd_01Res().Size, resMqtt.GetCmd_01Res().Value)
			case *air_switch_pb.MqttDeviceResponse_Cmd_02Res:
				fmt.Println(resMqtt.GetCmd_02Res().RetCmd, resMqtt.GetCmd_02Res().Size, resMqtt.GetCmd_02Res().Value)
			case *air_switch_pb.MqttDeviceResponse_Cmd_03Res:
				if (int32)(resMqtt.GetCmd_03Res().Value[0])*256+(int32)(resMqtt.GetCmd_03Res().Value[1]) != 0 {
					fmt.Println("debug", (int32)(resMqtt.GetCmd_03Res().Value[0])*256+(int32)(resMqtt.GetCmd_03Res().Value[1]))
				}
				/*
					if len(resMqtt.GetCmd_03Res().GetValue()) >= 2 {
						fmt.Println(resMqtt.GetCmd_03Res().RetCmd, resMqtt.GetCmd_03Res().Size, (int32)(resMqtt.GetCmd_03Res().Value[0])*256+(int32)(resMqtt.GetCmd_03Res().Value[1]))
					} else {
						fmt.Println("MqttDeviceResponse_Cmd_03 error")
					}*/
			case *air_switch_pb.MqttDeviceResponse_Cmd_04Res:
				fmt.Println(resMqtt.GetCmd_04Res().RetCmd, resMqtt.GetCmd_04Res().Size, (int32)(resMqtt.GetCmd_04Res().Value[0])*256+(int32)(resMqtt.GetCmd_04Res().Value[1]))
			case *air_switch_pb.MqttDeviceResponse_Cmd_05Res:
				fmt.Println(resMqtt.GetCmd_05Res().RetCmd, resMqtt.GetCmd_05Res().Size, (int32)(resMqtt.GetCmd_05Res().Value[0])*256+(int32)(resMqtt.GetCmd_05Res().Value[1]))
				break
			case *air_switch_pb.MqttDeviceResponse_Cmd_06Res:
				fmt.Println(resMqtt.GetCmd_06Res().RetCmd, resMqtt.GetCmd_06Res().Size, (int32)(resMqtt.GetCmd_06Res().Value[0])*256+(int32)(resMqtt.GetCmd_06Res().Value[1]))
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
			Payload: &air_switch_pb.MqttDeviceRequest_Cmd_03Req{
				Cmd_03Req: &air_switch_pb.CmdSubReq{
					SubCmd: 0x05,
					Target: 0,
					Value:  0x01,
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
		time.Sleep(100 * time.Millisecond) 
	}

}
