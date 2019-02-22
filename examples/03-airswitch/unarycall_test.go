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
	air_switch_pb "github.com/nayotta/metathings/pkg/proto/esp32_air_switch"
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
	MtpDeviceID        = "airswitch01"
)

var (
	testGetSwitchsAddr        int32 = 0x0009 //switch 0 and 3
	testSwitchAddr0           int32 = 3
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

type airSwitchTestSuite struct {
	suite.Suite
	token         string
	devicedClient deviced_pb.DevicedServiceClient
}

func (suite *airSwitchTestSuite) encodeMqttDevicedUnaryCallRequest(payload *air_switch_pb.MqttDeviceRequest) (*deviced_pb.UnaryCallResponse, error) {
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

func (suite *airSwitchTestSuite) sendRequest(payload *air_switch_pb.MqttDeviceRequest) (*air_switch_pb.MqttDeviceResponse, error) {
	res, err := suite.encodeMqttDevicedUnaryCallRequest(payload)
	if err != nil {
		return nil, err
	}

	resMqtt := air_switch_pb.MqttDeviceResponse{}
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
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchsAddrReq{
			GetSwitchsAddrReq: &air_switch_pb.GetSwitchsAddrReq{
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
	case *air_switch_pb.MqttDeviceResponse_GetSwitchsAddrRes:
		resTypeCheck = true
		suite.Equal(res.GetGetSwitchsAddrRes().GetAddrRes(), testGetSwitchsAddr)
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchsState() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchsStateReq{
			GetSwitchsStateReq: &air_switch_pb.GetSwitchsStateReq{
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
	case *air_switch_pb.MqttDeviceResponse_GetSwitchsStateRes:
		resTypeCheck = true
		// 实时状态 无法检查
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchData() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchDataReq{
			GetSwitchDataReq: &air_switch_pb.GetSwitchDataReq{
				SwitchAddr:  testSwitchAddr0,
				Votage:      true,
				LeakCurrent: true,
				Power:       true,
				Temp:        true,
				Current:     true,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_GetSwitchDataRes:
		resTypeCheck = true
		suite.Equal(res.GetGetSwitchDataRes().GetSwitchAddr(), testSwitchAddr0)
		// 实时状态 无法检查
		suite.NotNil(res.GetGetSwitchDataRes().GetVotage())
		suite.NotNil(res.GetGetSwitchDataRes().GetLeakCurrent())
		suite.NotNil(res.GetGetSwitchDataRes().GetPower())
		suite.NotNil(res.GetGetSwitchDataRes().GetTemp())
		suite.NotNil(res.GetGetSwitchDataRes().GetCurrent())
		break
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchConfig() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchConfigReq{
			GetSwitchConfigReq: &air_switch_pb.GetSwitchConfigReq{
				SwitchAddr:      testSwitchAddr0,
				VotageHigh:      true,
				VotageLow:       true,
				LeakCurrentHigh: true,
				PowerHigh:       true,
				TempHigh:        true,
				CurrentHigh:     true,
				Model:           true,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_GetSwitchConfigRes:
		resTypeCheck = true
		fmt.Println(res.Payload)
		suite.Equal(testSwitchAddr0, res.GetGetSwitchConfigRes().GetSwitchAddr())
		suite.Equal(testSwitchVotageHigh, res.GetGetSwitchConfigRes().GetVotageHigh().GetValue())
		suite.Equal(testSwitchVotageLow, res.GetGetSwitchConfigRes().GetVotageLow().GetValue())
		suite.Equal(testSwitchLeakCurrentHigh, res.GetGetSwitchConfigRes().GetLeakCurrentHigh().GetValue())
		suite.Equal(testSwitchPowerHigh, res.GetGetSwitchConfigRes().GetPowerHigh().GetValue())
		suite.Equal(testSwitchTempHigh, res.GetGetSwitchConfigRes().GetTempHigh().GetValue())
		suite.Equal(testSwitchCurrentHigh, res.GetGetSwitchConfigRes().GetCurrentHigh().GetValue())
		suite.Equal(testSwitchType0, res.GetGetSwitchConfigRes().GetType().GetValue())
		suite.Equal(testSwitchModel0, res.GetGetSwitchConfigRes().GetModel().GetValue())
		break
	}

	suite.Equal(resTypeCheck, true)
}

// need after set config
func (suite *airSwitchTestSuite) TestSwitchConfig() {
	//set config
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_SetSwitchConfigReq{
			SetSwitchConfigReq: &air_switch_pb.SetSwitchConfigReq{
				SwitchAddr:      testSwitchAddr0,
				VotageHigh:      &wrappers.Int32Value{Value: testSwitchVotageHigh},
				VotageLow:       &wrappers.Int32Value{Value: testSwitchVotageLow},
				LeakCurrentHigh: &wrappers.Int32Value{Value: testSwitchLeakCurrentHigh},
				PowerHigh:       &wrappers.Int32Value{Value: testSwitchPowerHigh},
				TempHigh:        &wrappers.Int32Value{Value: testSwitchTempHigh},
				CurrentHigh:     &wrappers.Int32Value{Value: testSwitchCurrentHigh},
				WarnVotageHigh:  &wrappers.Int32Value{Value: testSwitchWarnVotageHigh},
				WarnVotageLow:   &wrappers.Int32Value{Value: testSwitchWarnVotageLow},
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_SetSwitchConfigRes:
		resTypeCheck = true
		suite.Equal(testSwitchAddr0, res.GetSetSwitchConfigRes().GetSwitchAddr())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetVotageHigh())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetVotageLow())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetLeakCurrentHigh())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetPowerHigh())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetTempHigh())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetCurrentHigh())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetWarnVotageHigh())
		suite.Equal(true, res.GetSetSwitchConfigRes().GetWarnVotageLow())
		break
	}
	suite.Equal(resTypeCheck, true)

	// get config
	resTypeCheck = false
	payload = &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchConfigReq{
			GetSwitchConfigReq: &air_switch_pb.GetSwitchConfigReq{
				SwitchAddr:      testSwitchAddr0,
				VotageHigh:      true,
				VotageLow:       true,
				LeakCurrentHigh: true,
				PowerHigh:       true,
				TempHigh:        true,
				CurrentHigh:     true,
				Model:           true,
			},
		},
	}

	res, err = suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_GetSwitchConfigRes:
		resTypeCheck = true
		suite.Equal(testSwitchAddr0, res.GetGetSwitchConfigRes().GetSwitchAddr())
		suite.Equal(testSwitchVotageHigh, res.GetGetSwitchConfigRes().GetVotageHigh().GetValue())
		suite.Equal(testSwitchVotageLow, res.GetGetSwitchConfigRes().GetVotageLow().GetValue())
		suite.Equal(testSwitchLeakCurrentHigh, res.GetGetSwitchConfigRes().GetLeakCurrentHigh().GetValue())
		suite.Equal(testSwitchPowerHigh, res.GetGetSwitchConfigRes().GetPowerHigh().GetValue())
		suite.Equal(testSwitchTempHigh, res.GetGetSwitchConfigRes().GetTempHigh().GetValue())
		suite.Equal(testSwitchCurrentHigh, res.GetGetSwitchConfigRes().GetCurrentHigh().GetValue())
		suite.Equal(testSwitchType0, res.GetGetSwitchConfigRes().GetType().GetValue())
		suite.Equal(testSwitchModel0, res.GetGetSwitchConfigRes().GetModel().GetValue())
		break
	}

	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestSwitchState() {
	// test switch off
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_SetSwitchStateReq{
			SetSwitchStateReq: &air_switch_pb.SetSwitchStateReq{
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
	case *air_switch_pb.MqttDeviceResponse_SetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetSetSwitchStateRes().GetSwitchAddr())
		suite.Equal(false, res.GetSetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// wait switch act over
	time.Sleep(1 * time.Second)

	// test switch get
	resTypeCheck = false
	payload = &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchStateReq{
			GetSwitchStateReq: &air_switch_pb.GetSwitchStateReq{
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
	case *air_switch_pb.MqttDeviceResponse_GetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchStateRes().GetSwitchAddr())
		suite.Equal(false, res.GetGetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// test switch on
	resTypeCheck = false
	payload = &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_SetSwitchStateReq{
			SetSwitchStateReq: &air_switch_pb.SetSwitchStateReq{
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
	case *air_switch_pb.MqttDeviceResponse_SetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetSetSwitchStateRes().GetSwitchAddr())
		suite.Equal(true, res.GetSetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// wait switch act over
	time.Sleep(1 * time.Second)

	// test switch get
	resTypeCheck = false
	payload = &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchStateReq{
			GetSwitchStateReq: &air_switch_pb.GetSwitchStateReq{
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
	case *air_switch_pb.MqttDeviceResponse_GetSwitchStateRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchStateRes().GetSwitchAddr())
		suite.Equal(true, res.GetGetSwitchStateRes().GetState())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// leak test
	resTypeCheck = false
	payload = &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_SetSwitchLeakTestReq{
			SetSwitchLeakTestReq: &air_switch_pb.SetSwitchLeakTest{
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
	case *air_switch_pb.MqttDeviceResponse_SetSwitchLeakTestRes:
		suite.Equal(testSwitchAddr0, res.GetSetSwitchLeakTestRes().GetSwitchAddr())
		suite.Equal(true, res.GetSetSwitchLeakTestRes().GetResult())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestSwitchKwh() {
	// test set kwh
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_SetSwitchKwhReq{
			SetSwitchKwhReq: &air_switch_pb.SetSwitchKWhReq{
				SwitchAddr: testSwitchAddr0,
				Kwh:        testSwitchKwh,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_SetSwitchKwhRes:
		suite.Equal(testSwitchAddr0, res.GetSetSwitchKwhRes().GetSwitchAddr())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)

	// wait switch act over
	time.Sleep(5 * time.Second)

	// test get kwh
	resTypeCheck = false
	payload = &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchKwhReq{
			GetSwitchKwhReq: &air_switch_pb.GetSwitchKWhReq{
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
	case *air_switch_pb.MqttDeviceResponse_GetSwitchKwhRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchKwhRes().GetSwitchAddr())
		suite.Equal(testSwitchKwh, res.GetGetSwitchKwhRes().GetKwh())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchKwh() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchKwhReq{
			GetSwitchKwhReq: &air_switch_pb.GetSwitchKWhReq{
				SwitchAddr: testSwitchAddr0,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_GetSwitchKwhRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchKwhRes().GetSwitchAddr())
		suite.Equal(testSwitchKwh, res.GetGetSwitchKwhRes().GetKwh())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchWarn() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchWarnReq{
			GetSwitchWarnReq: &air_switch_pb.GetSwitchWarnReq{
				SwitchAddr: testSwitchAddr0,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_GetSwitchWarnRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchWarnRes().GetSwitchAddr())
		suite.Equal((int32)(0), res.GetGetSwitchWarnRes().GetWarnRes())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetTimeTask() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetTimeTaskReq{
			GetTimeTaskReq: &air_switch_pb.GetTimeTaskReq{
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
	case *air_switch_pb.MqttDeviceResponse_GetTimeTaskRes:

		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestSetTimeTask() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_SetTimeTaskReq{
			SetTimeTaskReq: &air_switch_pb.SetTimeTaskReq{},
		},
	}
	payload.GetSetTimeTaskReq().TimeTasks = append(payload.GetSetTimeTaskReq().TimeTasks,
		&air_switch_pb.TimeTask{
			Time: "00 22 15 * * *",
		})
	payload.GetSetTimeTaskReq().TimeTasks[0].Tasks = append(payload.GetSetTimeTaskReq().TimeTasks[0].Tasks,
		&air_switch_pb.MqttDeviceRequest{
			SessionId: 99,
			Payload: &air_switch_pb.MqttDeviceRequest_SetSwitchStateReq{
				SetSwitchStateReq: &air_switch_pb.SetSwitchStateReq{
					SwitchAddr: 3,
					State:      false,
				},
			},
		})

	payload.GetSetTimeTaskReq().TimeTasks[0].Tasks = append(payload.GetSetTimeTaskReq().TimeTasks[0].Tasks,
		&air_switch_pb.MqttDeviceRequest{
			SessionId: 99,
			Payload: &air_switch_pb.MqttDeviceRequest_Cmd_06Req{
				Cmd_06Req: &air_switch_pb.CmdSubReq{
					SubCmd: 0x07,
					Target: 3,
					Value:  0xa5,
				},
			},
		})

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_SetTimeTaskRes:

		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetSwitchCtrl() {
	// test get kwh
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_GetSwitchCtrlReq{
			GetSwitchCtrlReq: &air_switch_pb.GetSwitchCtrlReq{
				SwitchAddr: testSwitchAddr0,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	fmt.Println(res.Payload)
	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_GetSwitchCtrlRes:
		suite.Equal(testSwitchAddr0, res.GetGetSwitchCtrlRes().GetSwitchAddr())
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func (suite *airSwitchTestSuite) TestGetRunTime() {
	resTypeCheck := false
	payload := &air_switch_pb.MqttDeviceRequest{
		Payload: &air_switch_pb.MqttDeviceRequest_RunTimeReq{
			RunTimeReq: &air_switch_pb.RunTimeReq{
				Time:0,
			},
		},
	}

	res, err := suite.sendRequest(payload)
	suite.Nil(err)
	if err != nil {
		return
	}

	switch res.Payload.(type) {
	case *air_switch_pb.MqttDeviceResponse_RunTimeRes:
		fmt.Println(res)
		resTypeCheck = true
	}
	suite.Equal(resTypeCheck, true)
}

func TestAirSwitchTestSuite(t *testing.T) {
	suite.Run(t, new(airSwitchTestSuite))
}
