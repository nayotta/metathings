package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	ctx_help "github.com/nayotta/metathings/pkg/common/context"
	kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	switch_pb "github.com/nayotta/metathings/pkg/proto/esp8266_switch_1902"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	MtpURL             = "metathings.ai:21733"
	MtpUsername        = "admin"
	MtpPassword        = "admin"
	MtpDefaultDomainID = "default"
	MtpDeviceID        = "test"
)

var (
	testSwitchAddr0    int32 = 12     //switch 3
	testGetSwitchAddr0 int32 = 0x0008 //switch 3 bit 4
	testVotageHigh     int32 = 235000
	testVotageLow      int32 = 215000
	testCurrentHigh    int32 = 5000
)

type airSwitchTestSuite struct {
	suite.Suite
	token         string
	devicedClient deviced_pb.DevicedServiceClient
}

func (suite *airSwitchTestSuite) encodeMqttDevicedUnaryCallRequest(payload *switch_pb.MqttDeviceRequest) (*deviced_pb.UnaryCallResponse, error) {
	payloadByte, err := proto.Marshal(payload)
	if err != nil {
		fmt.Printf("proto marshal error:%v\n", err)
		return nil, err
	}

	req := &deviced_pb.UnaryCallRequest{
		Device: &deviced_pb.OpDevice{
			Id: &wrappers.StringValue{
				Value: MtpDeviceID,
			},
			Kind: kind.DeviceKind_DEVICE_KIND_SIMPLE,
		},
		Value: &deviced_pb.OpUnaryCallValue{
			Name: &wrappers.StringValue{
				Value: MtpDeviceID,
			},
			Component: &wrappers.StringValue{
				Value: MtpDeviceID,
			},
			Method: &wrappers.StringValue{
				Value: MtpDeviceID,
			},
			Value: &any.Any{
				Value: payloadByte,
			},
		},
	}

	ctx := context.Background()
	ctx = ctx_help.WithToken(ctx, suite.token)
	res, err := suite.devicedClient.UnaryCall(ctx, req)

	return res, nil
}

func (suite *airSwitchTestSuite) sendRequest(payload *switch_pb.MqttDeviceRequest) (*switch_pb.MqttDeviceResponse, error) {
	res, err := suite.encodeMqttDevicedUnaryCallRequest(payload)
	if err != nil {
		return nil, err
	}

	resMqtt := switch_pb.MqttDeviceResponse{}
	err = proto.Unmarshal(res.GetValue().GetValue().GetValue(), &resMqtt)
	if err != nil {
		return nil, err
	}

	return &resMqtt, nil
}

func (suite *airSwitchTestSuite) SetupTest() {
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
	suite.devicedClient = devicedClient

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

	suite.token = tokenStr
	println("test setup")
}

func (suite *airSwitchTestSuite) TestGetSwitchsAddr() {
	resTypeCheck := false
	payload := &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchsAddrReq{
			GetSwitchsAddrReq: &switch_pb.GetSwitchsAddrReq{
				Code: 0,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchsAddrRes:
		resTypeCheck = true
		suite.Equal(res.GetGetSwitchsAddrRes().GetAddrRes(), testGetSwitchAddr0)
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchsState() {
	resTypeCheck := false
	payload := &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchsStateReq{
			GetSwitchsStateReq: &switch_pb.GetSwitchsStateReq{
				Num: 9,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchsStateRes:
		resTypeCheck = true
		// 实时状态 无法检查
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestSwitchState() {
	// test switch off
	resTypeCheck := false
	payload := &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_SetSwitchStateReq{
			SetSwitchStateReq: &switch_pb.SetSwitchStateReq{
				SwitchAddr: testSwitchAddr0,
				State:      false,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_SetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetSetSwitchStateRes().GetSwitchAddr())
		suite.Equal(false, res.GetSetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// wait switch act over
	time.Sleep(1 * time.Second)

	// test switch get
	resTypeCheck = false
	payload = &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchStateReq{
			GetSwitchStateReq: &switch_pb.GetSwitchStateReq{
				SwitchAddr: testSwitchAddr0,
			},
		},
	}

	res, err = suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchStateRes().GetSwitchAddr())
		suite.Equal(false, res.GetGetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// test switch on
	resTypeCheck = false
	payload = &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_SetSwitchStateReq{
			SetSwitchStateReq: &switch_pb.SetSwitchStateReq{
				SwitchAddr: testSwitchAddr0,
				State:      true,
			},
		},
	}

	res, err = suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_SetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetSetSwitchStateRes().GetSwitchAddr())
		suite.Equal(true, res.GetSetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// wait switch act over
	time.Sleep(1 * time.Second)

	// test switch get
	resTypeCheck = false
	payload = &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchStateReq{
			GetSwitchStateReq: &switch_pb.GetSwitchStateReq{
				SwitchAddr: testSwitchAddr0,
			},
		},
	}

	res, err = suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchStateRes().GetSwitchAddr())
		suite.Equal(true, res.GetGetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchsData() {
	resTypeCheck := false
	payload := &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchsDataReq{
			GetSwitchsDataReq: &switch_pb.GetSwitchsDataReq{
				Votage:  true,
				Current: true,
				Power:   true,
				Kwh:     true,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchsDataRes:
		fmt.Println(res)
		resTypeCheck = true
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchsConfig() {
	resTypeCheck := false
	payload := &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchsConfigReq{
			GetSwitchsConfigReq: &switch_pb.GetSwitchsConfigReq{
				VotageHigh:   true,
				VotageLow:    true,
				CurrentHight: true,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchsConfigRes:
		fmt.Println(res)
		resTypeCheck = true
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestSwitchsConfig() {
	// set config
	resTypeCheck := false
	payload := &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_SetSwitchsConfigReq{
			SetSwitchsConfigReq: &switch_pb.SetSwitchsConfigReq{
				VotageHigh:  &wrappers.Int32Value{Value: testVotageHigh},
				VotageLow:   &wrappers.Int32Value{Value: testVotageLow},
				CurrentHigh: &wrappers.Int32Value{Value: testCurrentHigh},
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_SetSwitchsConfigRes:
		fmt.Println(res)
		resTypeCheck = true
		break
	}
	suite.Equal(resTypeCheck, true)

	// get config
	payload = &switch_pb.MqttDeviceRequest{
		Payload: &switch_pb.MqttDeviceRequest_GetSwitchsConfigReq{
			GetSwitchsConfigReq: &switch_pb.GetSwitchsConfigReq{
				VotageHigh:   true,
				VotageLow:    true,
				CurrentHight: true,
			},
		},
	}

	res, err = suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *switch_pb.MqttDeviceResponse_GetSwitchsConfigRes:
		fmt.Println(res)
		resTypeCheck = true
		break
	}
	suite.Equal(resTypeCheck, true)
}

func TestAirSwitchTestSuite(t *testing.T) {
	suite.Run(t, new(airSwitchTestSuite))
}
