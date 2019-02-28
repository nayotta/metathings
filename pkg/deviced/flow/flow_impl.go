package metathings_deviced_flow

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/protobuf/jsonpb"
	struct_ "github.com/golang/protobuf/ptypes/struct"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	json_encoder = jsonpb.Marshaler{}
	json_decoder = jsonpb.Unmarshaler{}
)

type FlowOption struct {
	Id         string
	DevId      string
	KfkBrokers []string
}

type FlowImpl struct {
	opt           *FlowOption
	mgo_db        *mongo.Database
	kfk_prod      sarama.SyncProducer
	kfk_prod_once *sync.Once
	logger        log.FieldLogger
}

func (f *FlowImpl) mongo_collection() *mongo.Collection {
	return f.mgo_db.Collection("metathings.flow." + f.Id())
}

func (f *FlowImpl) kafka_topic() string {
	return "metathings.flow." + f.Id()
}

func (f *FlowImpl) init_sarama_producer() {
	var err error
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	f.kfk_prod, err = sarama.NewSyncProducer(f.opt.KfkBrokers, cfg)
	if err != nil {
		// TODO(Peer): handle error
		panic(err)
	}
	f.logger.WithField("topic", f.kafka_topic()).Debugf("init sarama producer")
}

func (f *FlowImpl) new_sarama_cluster_consumer(grp string, topics []string) (*cluster.Consumer, error) {
	cfg := cluster.NewConfig()
	cfg.Consumer.Return.Errors = true
	cfg.Group.Return.Notifications = true
	cfg.Group.Heartbeat.Interval = 100 * time.Millisecond
	consumer, err := cluster.NewConsumer(f.opt.KfkBrokers, grp, topics, cfg)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (f *FlowImpl) get_sarama_producer() sarama.SyncProducer {
	f.kfk_prod_once.Do(f.init_sarama_producer)
	return f.kfk_prod
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

func (f *FlowImpl) Close() error {
	var err error
	if f.kfk_prod != nil {
		if err = f.kfk_prod.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (f *FlowImpl) PushFrame(frm *pb.Frame) error {
	ts := pb_helper.Now()
	frm.Ts = &ts

	err := f.push_frame_to_mgo(frm)
	if err != nil {
		return err
	}

	// TODO(Peer): dont push frame to kafka when noone pull frame.
	err = f.push_frame_to_kafka(frm)
	if err != nil {
		return err
	}

	return nil
}

func (f *FlowImpl) push_frame_to_kafka(frm *pb.Frame) error {
	frm_txt, err := json_encoder.MarshalToString(frm)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     f.kafka_topic(),
		Value:     sarama.ByteEncoder(frm_txt),
		Partition: -1,
	}

	_, _, err = f.get_sarama_producer().SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (f *FlowImpl) push_frame_to_mgo(frm *pb.Frame) error {
	frm_dat := frm.GetData()
	frm_dat_txt, err := json_encoder.MarshalToString(frm_dat)
	if err != nil {
		return err
	}

	frm_dat_buf := bson.M{}
	err = bson.UnmarshalExtJSON([]byte(frm_dat_txt), true, &frm_dat_buf)
	if err != nil {
		return err
	}
	frm_dat_buf["#ts"] = pb_helper.ToTime(*frm.GetTs()).UnixNano()

	coll := f.mongo_collection()
	_, err = coll.InsertOne(f.context(), frm_dat_buf)
	if err != nil {
		return err
	}

	return nil
}

func (f *FlowImpl) PullFrame() (<-chan *pb.Frame, <-chan struct{}) {
	frm_chan := make(chan *pb.Frame)
	quit_chan := make(chan struct{})

	go func() {
		defer close(frm_chan)

		grp := id_helper.NewId()
		consumer, err := f.new_sarama_cluster_consumer(grp, []string{f.kafka_topic()})
		defer consumer.Close()
		if err != nil {
			f.logger.WithError(err).Debugf("failed to new sarama cluster consumer")
			return
		}

		for {
			select {
			case <-quit_chan:
				f.logger.Debugf("receive quit signal from outside")
				return
			case msg, ok := <-consumer.Messages():
				if ok {
					consumer.MarkOffset(msg, "")
					var frm pb.Frame
					err = json_decoder.Unmarshal(strings.NewReader(string(msg.Value)), &frm)
					if err != nil {
						f.logger.WithError(err).Debugf("failed to decode frame from message")
						return
					}
					frm_chan <- &frm
				}
			case err := <-consumer.Errors():
				f.logger.WithError(err).Debugf("failed to receive message from kafka")
				return
			case ntf := <-consumer.Notifications():
				f.logger.WithField("notification", ntf).Debugf("receive notification from kafka")
			case prt := <-consumer.Partitions():
				f.logger.WithField("partition", prt).Debugf("receive partition change from kafka")
			}
		}
	}()

	return frm_chan, quit_chan
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

	cur, err := coll.Find(f.context(), bson.M{"#ts": flr_buf})
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

		ts, ok := res_buf["#ts"]
		if !ok {
			ts = nil
		}
		delete(res_buf, "#ts")

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
	var mgo_db *mongo.Database

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
		"mongo_database": func(key string, val interface{}) error {
			mgo_db, ok = val.(*mongo.Database)
			if !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
	})(args...)
	if err != nil {
		return nil, err
	}

	return &FlowImpl{
		opt:           opt,
		mgo_db:        mgo_db,
		logger:        logger,
		kfk_prod_once: new(sync.Once),
	}, nil
}
