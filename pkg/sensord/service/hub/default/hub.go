package default_hub

import (
	"math/rand"
	"sync"

	log "github.com/sirupsen/logrus"

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

func path(opt opt_helper.Option) string {
	return opt.GetString("sensor_id")
}

func (self *defaultHub) Subscriber(opt opt_helper.Option) (hub.Subscriber, error) {
	var ok bool
	var id uint64
	var m map[uint64]chan *sensord_pb.SensorData
	path := path(opt)

	self.glock.Lock()
	defer self.glock.Unlock()

	if m, ok = self.subs[path]; !ok {
		m = make(map[uint64]chan *sensord_pb.SensorData)
		self.subs[path] = m
	}

	ch := make(chan *sensord_pb.SensorData)
	id = rand.Uint64()
	m[id] = ch

	sub := &subscriber{
		p:  path,
		id: id,
		ch: ch,
	}

	return sub, nil
}

func (self *defaultHub) Publisher(opt opt_helper.Option) (hub.Publisher, error) {
	var ok bool
	var ch chan *sensord_pb.SensorData
	path := path(opt)

	self.glock.Lock()
	defer self.glock.Unlock()

	if ch, ok = self.pubs[path]; !ok {
		ch = make(chan *sensord_pb.SensorData)
		self.pubs[path] = ch
		self.pub_counters[path] = 0
		go self.transfer(path, ch)
	}

	id := rand.Uint64()

	pub := &publisher{
		p:  path,
		id: id,
		ch: ch,
	}
	self.pub_counters[path]++

	return pub, nil
}

func (self *defaultHub) Close(sp hub.SubPub) error {
	switch sp.(type) {
	case hub.Subscriber:
		return self.closeSub(sp)
	case hub.Publisher:
		return self.closePub(sp)
	}
	return nil
}

func (self *defaultHub) closeSub(sp hub.SubPub) error {
	self.glock.Lock()
	defer self.glock.Unlock()

	p := sp.Path()
	id := sp.Id()

	subs, ok := self.subs[p]
	if !ok {
		self.logger.WithFields(log.Fields{"path": p, "id": id}).Warningf("subscriber not found")
		return hub.ErrSubPubNotFound
	}

	ch, ok := subs[id]
	if !ok {
		self.logger.WithFields(log.Fields{"path": p, "id": id}).Warningf("subscriber not found")
		return hub.ErrSubPubNotFound
	}

	close(ch)
	delete(self.subs[p], id)
	self.logger.WithFields(log.Fields{"path": p, "id": id}).Debugf("close subscriber")
	return nil
}

func (self *defaultHub) closePub(sp hub.SubPub) error {
	self.glock.Lock()
	defer self.glock.Unlock()

	p := sp.Path()

	ch, ok := self.pubs[p]
	if !ok {
		return hub.ErrSubPubNotFound
	}

	if _, ok := self.pub_counters[p]; !ok {
		self.pub_counters[p] = 0
		close(ch)
		self.logger.WithField("path", p).Warningf("close channel with unexpected situation")
		return hub.ErrUnexpected
	}

	self.pub_counters[p]--
	if self.pub_counters[p] < 0 {
		self.pub_counters[p] = 0
		self.logger.WithField("path", p).Warningf("reset counter to 0")
	}

	if self.pub_counters[p] == 0 {
		close(ch)
		delete(self.pubs, p)
		self.logger.WithField("path", p).Debugf("close publisher")
	}

	return nil
}

func (self *defaultHub) transfer(path string, ch chan *sensord_pb.SensorData) {
	for {
		dat, ok := <-ch
		if !ok {
			self.logger.WithField("path", path).Debugf("failed to recv data from channel, maybe closed")
			return
		}

		go func(dat *sensord_pb.SensorData) {
			subs := self.subs[path]
			for id := range subs {
				ch := subs[id]
				ch <- dat
			}
		}(dat)
	}
}

type subscriber struct {
	id uint64
	p  string
	ch chan *sensord_pb.SensorData
}

func (self *subscriber) Subscribe() (*sensord_pb.SensorData, error) {
	dat, ok := <-self.ch
	if !ok {
		return nil, hub.ErrUnsubscribable
	}
	return dat, nil
}

func (self *subscriber) Id() uint64 {
	return self.id
}

func (self *subscriber) Path() string {
	return self.p
}

type publisher struct {
	id uint64
	p  string
	ch chan *sensord_pb.SensorData
}

func (self *publisher) Publish(dat *sensord_pb.SensorData) error {
	self.ch <- dat
	return nil
}

func (self *publisher) Id() uint64 {
	return self.id
}

func (self *publisher) Path() string {
	return self.p
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
