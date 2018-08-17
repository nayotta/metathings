package default_hub

import (
	"fmt"
	"math/rand"
	"sync"

	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
	"github.com/nayotta/metathings/pkg/sensord/service/hub"
)

type defaultHub struct {
	logger log.FieldLogger
	glock  *sync.Mutex

	pubs         map[string]chan *sensord_pb.SensorData
	subs         map[string]map[uint64]chan *sensord_pb.SensorData
	pub_counters map[string]uint64
}

func symbol(opt opt_helper.Option) string {
	return fmt.Sprintf("sensor.%v", opt.GetString("sensor_id"))
}

func (self *defaultHub) NewSubscriber(opt opt_helper.Option) (hub.Subscriber, error) {
	var ok bool
	var id uint64
	var m map[uint64]chan *sensord_pb.SensorData
	sym := symbol(opt)

	self.glock.Lock()
	defer self.glock.Unlock()

	if m, ok = self.subs[sym]; !ok {
		m = make(map[uint64]chan *sensord_pb.SensorData)
		self.subs[sym] = m
	}

	ch := make(chan *sensord_pb.SensorData)
	id = id_helper.NewUint64Id()
	m[id] = ch

	sub := &subscriber{
		sym: sym,
		id:  id,
		ch:  ch,
		q:   make(chan interface{}, 1),
	}
	sub.cls_cb = func() error {
		return self.closeSub(sub)
	}

	return sub, nil
}

func (self *defaultHub) NewPublisher(opt opt_helper.Option) (hub.Publisher, error) {
	var ok bool
	var ch chan *sensord_pb.SensorData
	sym := symbol(opt)

	self.glock.Lock()
	defer self.glock.Unlock()

	if ch, ok = self.pubs[sym]; !ok {
		ch = make(chan *sensord_pb.SensorData)
		self.pubs[sym] = ch
		self.pub_counters[sym] = 0
		go self.transfer(sym, ch)
	}

	id := rand.Uint64()

	pub := &publisher{
		sym: sym,
		id:  id,
		ch:  ch,
	}
	pub.cls_cb = func() error {
		return self.closePub(pub)
	}
	self.pub_counters[sym]++

	return pub, nil
}

func (self *defaultHub) GetSubscriber(opt opt_helper.Option) (hub.Subscriber, error) {
	panic("unimplemented")
}

func (self *defaultHub) GetPublisher(opt opt_helper.Option) (hub.Publisher, error) {
	panic("unimplemented")
}

func (self *defaultHub) closeSub(sp hub.SubPub) error {
	self.glock.Lock()
	defer self.glock.Unlock()

	sym := sp.Symbol()
	id := sp.Id()

	subs, ok := self.subs[sym]
	if !ok {
		self.logger.WithFields(log.Fields{"sym": sym, "id": id}).Warningf("subscriber not found")
		return hub.ErrSubPubNotFound
	}

	ch, ok := subs[id]
	if !ok {
		self.logger.WithFields(log.Fields{"sym": sym, "id": id}).Warningf("subscriber not found")
		return hub.ErrSubPubNotFound
	}

	close(ch)
	delete(self.subs[sym], id)
	self.logger.WithFields(log.Fields{"sym": sym, "id": id}).Debugf("close subscriber")
	return nil
}

func (self *defaultHub) closePub(sp hub.SubPub) error {
	self.glock.Lock()
	defer self.glock.Unlock()

	sym := sp.Symbol()

	ch, ok := self.pubs[sym]
	if !ok {
		return hub.ErrSubPubNotFound
	}

	if _, ok := self.pub_counters[sym]; !ok {
		self.pub_counters[sym] = 0
		close(ch)
		self.logger.WithField("sym", sym).Warningf("close channel with unexpected situation")
		return hub.ErrUnexpected
	}

	self.pub_counters[sym]--
	if self.pub_counters[sym] < 0 {
		self.pub_counters[sym] = 0
		self.logger.WithField("sym", sym).Warningf("reset counter to 0")
	}

	if self.pub_counters[sym] == 0 {
		close(ch)
		delete(self.pubs, sym)
		self.logger.WithField("sym", sym).Debugf("close publisher")
	}

	return nil
}

func (self *defaultHub) transfer(sym string, ch chan *sensord_pb.SensorData) {
	for {
		dat, ok := <-ch
		if !ok {
			self.logger.WithField("sym", sym).Debugf("failed to recv data from channel, maybe closed")
			return
		}

		go func(dat *sensord_pb.SensorData) {
			subs := self.subs[sym]
			for id := range subs {
				ch := subs[id]
				ch <- dat
			}
		}(dat)
	}
}

type subscriber struct {
	id     uint64
	sym    string
	ch     chan *sensord_pb.SensorData
	cls_cb func() error
	q      chan interface{}
}

func (self *subscriber) Subscribe() (*sensord_pb.SensorData, error) {
	var dat *sensord_pb.SensorData
	var ok bool
	select {
	case dat, ok = <-self.ch:
		if !ok {
			return nil, hub.ErrUnsubscribable
		}
		return dat, nil
	case <-self.q:
		return nil, hub.Terminated
	}

}

func (self *subscriber) Id() uint64 {
	return self.id
}

func (self *subscriber) Symbol() string {
	return self.sym
}

func (self *subscriber) Close() error {
	self.q <- nil
	return self.cls_cb()
}

type publisher struct {
	id     uint64
	sym    string
	ch     chan *sensord_pb.SensorData
	cls_cb func() error
}

func (self *publisher) Publish(dat *sensord_pb.SensorData) error {
	self.ch <- dat
	return nil
}

func (self *publisher) Id() uint64 {
	return self.id
}

func (self *publisher) Symbol() string {
	return self.sym
}

func (self *publisher) Close() error {
	return self.cls_cb()
}

func NewHub(opt opt_helper.Option) (hub.Hub, error) {
	return &defaultHub{
		glock:  new(sync.Mutex),
		logger: opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{"#module": "hub", "#driver": "default"}),

		pubs:         make(map[string]chan *sensord_pb.SensorData),
		subs:         make(map[string]map[uint64]chan *sensord_pb.SensorData),
		pub_counters: make(map[string]uint64),
	}, nil
}

func init() {
	hub.XXX_RegisterHub("default", NewHub)
}
