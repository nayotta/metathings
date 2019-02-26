package metathings_deviced_flow

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb"
	struct_ "github.com/golang/protobuf/ptypes/struct"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

const (
	frm0_dat = `{"v": 0}`
	frm1_dat = `{"v": 1}`
	frm2_dat = `{"v": 2}`
)

type FlowImplTestSuite struct {
	mgo_cli *mongo.Client
	opt     *FlowOption
	flow    *FlowImpl
	enc     *jsonpb.Marshaler
	dec     *jsonpb.Unmarshaler
	push_at time.Time
	suite.Suite
}

func (s *FlowImplTestSuite) SetupTest() {
	var opt FlowOption

	uri := test_helper.GetTestMongoUri()
	opt.MgoDb = test_helper.GetTestMongoDatabase()
	opt.Id = test_helper.GetenvWithDefault("MTT_FLOW_ID", "floooow")
	opt.DevId = test_helper.GetenvWithDefault("MTT_DEVICE_ID", "deeeev")

	logger, err := log_helper.NewLogger("test", "debug")
	s.Nil(err)
	mgo_cli, err := mongo.Connect(context.TODO(), uri)
	s.Nil(err)
	s.flow, err = new_flow_impl("option", &opt, "logger", logger, "mongo_client", mgo_cli)
	s.Nil(err)

	// clean up database
	s.Nil(s.flow.mgo_db.Drop(s.flow.context()))

	// insert prepared data
	s.push_at = time.Now()
	var dat struct_.Struct
	s.Nil(s.dec.Unmarshal(strings.NewReader(frm0_dat), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))

	time.Sleep(100 * time.Millisecond)
	s.Nil(s.dec.Unmarshal(strings.NewReader(frm1_dat), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))

	time.Sleep(100 * time.Millisecond)
	s.Nil(s.dec.Unmarshal(strings.NewReader(frm2_dat), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))
}

func (s *FlowImplTestSuite) TestPushFrame() {
	push_at := time.Now()
	var dat struct_.Struct
	s.Nil(s.dec.Unmarshal(strings.NewReader(`{"v": 4}`), &dat))
	s.Nil(s.flow.PushFrame(&pb.Frame{Data: &dat}))

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
