package metathings_deviced_flow

import (
	"context"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	struct_ "github.com/golang/protobuf/ptypes/struct"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type FlowOption struct {
	Id    string
	DevId string
	MgoDb string
}

type FlowImpl struct {
	opt    *FlowOption
	mgo_db *mongo.Database
	logger log.FieldLogger
}

func (f *FlowImpl) mongo_collection() *mongo.Collection {
	return f.mgo_db.Collection("metathings.flow." + f.Id())
}

func (f *FlowImpl) context() context.Context {
	return context.TODO()
}

func (f *FlowImpl) Id() string {
	return f.opt.Id
}

func (f *FlowImpl) Device() string {
	return f.opt.DevId
}

func (f *FlowImpl) PushFrame(frm *pb.Frame) error {
	frm_dat := frm.GetData()
	enc := jsonpb.Marshaler{}
	frm_dat_txt, err := enc.MarshalToString(frm_dat)
	if err != nil {
		return err
	}

	frm_dat_buf := bson.M{}
	err = bson.UnmarshalExtJSON([]byte(frm_dat_txt), true, &frm_dat_buf)
	if err != nil {
		return err
	}
	frm_dat_buf["ts"] = time.Now().UnixNano()

	coll := f.mongo_collection()
	_, err = coll.InsertOne(f.context(), frm_dat_buf)
	if err != nil {
		return err
	}

	return nil
}

func (f *FlowImpl) PullFrame(flr *FlowFilter) <-chan *pb.Frame {
	panic("unimplemented")
}

func (f *FlowImpl) QueryFrame(flrs ...*FlowFilter) ([]*pb.Frame, error) {
	var ret []*pb.Frame

	coll := f.mongo_collection()

	for _, flr := range flrs {
		frms, err := f.query_frame(coll, flr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, frms...)
	}

	return ret, nil
}

func (f *FlowImpl) query_frame(coll *mongo.Collection, flr *FlowFilter) ([]*pb.Frame, error) {
	flr_buf := bson.M{}
	if !flr.BeginAt.Equal(time.Time{}) {
		flr_buf["$gte"] = flr.BeginAt.UnixNano()
	}

	if !flr.EndAt.Equal(time.Time{}) {
		flr_buf["$lte"] = flr.EndAt.UnixNano()
	}

	cur, err := coll.Find(f.context(), bson.M{"ts": flr_buf})
	if err != nil {
		return nil, err
	}

	dec := jsonpb.Unmarshaler{}
	var frms []*pb.Frame
	for cur.Next(f.context()) {
		var res_buf bson.M
		if err = cur.Decode(&res_buf); err != nil {
			return nil, err
		}

		ts, ok := res_buf["ts"]
		if !ok {
			ts = nil
		}
		delete(res_buf, "ts")

		ts_int64, ok := ts.(int64)
		if !ok {
			ts_int64 = 0
		}

		frm_dat_txt, err := bson.MarshalExtJSON(res_buf, true, false)
		if err != nil {
			return nil, err
		}

		var frm_dat struct_.Struct
		err = dec.Unmarshal(strings.NewReader(string(frm_dat_txt)), &frm_dat)
		if err != nil {
			return nil, err
		}

		frm_ts := pb_helper.FromTimestamp(ts_int64)
		frm := &pb.Frame{
			Ts:   &frm_ts,
			Data: &frm_dat,
		}

		frms = append(frms, frm)
	}

	return frms, nil
}

func new_flow_impl(args ...interface{}) (*FlowImpl, error) {
	var ok bool
	var logger log.FieldLogger
	var opt *FlowOption
	var mgo_cli *mongo.Client

	err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"option": func(key string, val interface{}) error {
			opt, ok = val.(*FlowOption)
			if !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
		"logger": func(key string, val interface{}) error {
			logger, ok = val.(log.FieldLogger)
			if !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
		"mongo_client": func(key string, val interface{}) error {
			mgo_cli, ok = val.(*mongo.Client)
			if !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
	})(args...)
	if err != nil {
		return nil, err
	}

	mgo_db := mgo_cli.Database(opt.MgoDb)

	return &FlowImpl{
		opt:    opt,
		mgo_db: mgo_db,
		logger: logger,
	}, nil
}
