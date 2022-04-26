package metathings_deviced_flow

import (
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	struct_ "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
	pb "github.com/nayotta/metathings/proto/deviced"
)

const (
	frm0_dat = `{"v": 0}`
	frm1_dat = `{"v": 1}`
	frm2_dat = `{"v": 2}`
)

type FlowImplTestSuite struct {
	mgo_cli  *mongo.Client
	opt      *FlowOption
	flow     *flow
	flow_fty *flowFactory
	push_at  time.Time
	suite.Suite
}

func (s *FlowImplTestSuite) SetupTest() {
	var opt FlowOption

	mgo_uri := test_helper.GetTestMongoUri()
	mgo_db := test_helper.GetTestMongoDatabase()
	rs_addr := test_helper.GetTestRedisAddr()
	rs_db, err := strconv.Atoi(test_helper.GetTestRedisDB())
	s.Nil(err)

	opt.FlowId = test_helper.GetenvWithDefault("MTT_FLOW_ID", "floooow")
	opt.DeviceId = test_helper.GetenvWithDefault("MTT_DEVICE_ID", "deeeev")

	logger, err := log_helper.NewLogger("test", "debug")
	s.Nil(err)

	flw_fty, err := new_default_flow_factory(
		"redis_stream_addr", rs_addr,
		"redis_stream_db", rs_db,
		"mongo_uri", mgo_uri,
		"mongo_database", mgo_db,
		"logger", logger,
	)
	s.flow_fty = flw_fty.(*flowFactory)

	flw, err := s.flow_fty.New(&opt)
	s.Nil(err)
	s.flow = flw.(*flow)

	// insert prepared data
	s.push_at = time.Now()
	var dat struct_.Struct
	s.Nil(grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(frm0_dat), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))

	time.Sleep(100 * time.Millisecond)
	s.Nil(grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(frm1_dat), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))

	time.Sleep(100 * time.Millisecond)
	s.Nil(grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(frm2_dat), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))
}

func (s *FlowImplTestSuite) TestPushFrame() {
	wg := new(sync.WaitGroup)
	push_at := time.Now()
	var dat struct_.Struct
	s.Nil(grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(`{"v": 4}`), &dat))

	wg.Add(1)
	go func() {
		frm_ch := s.flow.PullFrame()
		frm := <-frm_ch
		s.NotNil(frm)
		wg.Done()
	}()

	time.Sleep(500 * time.Millisecond)
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))
	wg.Wait()

	frms, err := s.flow.QueryFrame(&FlowFilter{BeginAt: push_at})
	s.Nil(err)
	s.Len(frms, 1)
}

func (s *FlowImplTestSuite) TestQueryFrame() {
	frms, err := s.flow.QueryFrame(&FlowFilter{
		BeginAt: s.push_at,
	})
	s.Nil(err)
	s.Len(frms, 3)

	frms, err = s.flow.QueryFrame(&FlowFilter{
		EndAt: s.push_at,
	})
	s.Nil(err)
	s.Len(frms, 0)

	frms, err = s.flow.QueryFrame(&FlowFilter{
		BeginAt: s.push_at.Add(50 * time.Millisecond),
	})
	s.Nil(err)
	s.Len(frms, 2)

	frms, err = s.flow.QueryFrame(&FlowFilter{
		BeginAt: s.push_at.Add(50 * time.Millisecond),
		EndAt:   s.push_at.Add(150 * time.Millisecond),
	})
	s.Nil(err)
	s.Len(frms, 1)
}

func TestFlowImplTestSuite(t *testing.T) {
	suite.Run(t, new(FlowImplTestSuite))
}
