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

func (h *defaultHub) Subscriber(path string) hub.Subscriber {
	var ok bool
	var id uint64
	var m map[uint64]chan *sensord_pb.SensorData

	h.glock.Lock()
	defer h.glock.Unlock()

	if m, ok = h.subs[path]; !ok {
		m = make(map[uint64]chan *sensord_pb.SensorData)
		h.subs[path] = m
	}

	ch := make(chan *sensord_pb.SensorData)
	id = rand.Uint64()
	m[id] = ch

	sub := &subscriber{
		id: id,
		ch: ch,
	}

	return sub
}

func (h *defaultHub) Publisher(path string) hub.Publisher {
	var ok bool
	var ch chan *sensord_pb.SensorData

	h.glock.Lock()
	defer h.glock.Unlock()

	if ch, ok = h.pubs[path]; !ok {
		ch = make(chan *sensord_pb.SensorData)
		h.pubs[path] = ch
		h.pub_counters[path] = 0
		go h.transfer(path, ch)
	}

	id := rand.Uint64()

	pub := &publisher{
		id: id,
		ch: ch,
	}
	h.pub_counters[path]++

	return pub
}

func (h *defaultHub) Close(sp hub.SubPub) {
	switch sp.(type) {
	case hub.Subscriber:
		h.closeSub(sp)
	case hub.Publisher:
		h.closePub(sp)
	}
}

func (h defaultHub) closeSub(sp hub.SubPub) {
	h.glock.Lock()
	defer h.glock.Unlock()

	p := sp.Path()
	id := sp.Id()

	subs, ok := h.subs[p]
	if !ok {
		return
	}

	ch, ok := subs[id]
	if !ok {
		return
	}

	close(ch)
	delete(h.subs[p], id)
}

func (h defaultHub) closePub(sp hub.SubPub) {
	h.glock.Lock()
	defer h.glock.Unlock()

	p := sp.Path()

	ch, ok := h.pubs[p]
	if !ok {
		return
	}

	if _, ok := h.pub_counters[p]; !ok {
		h.pub_counters[p] = 0
		close(ch)
		h.logger.WithField("path", p).Warningf("SHOULD NOT HAPPEND!!! FIXME PLS!!! close channel with unexpected situation")
		return
	}

	h.pub_counters[p]--
	if h.pub_counters[p] < 0 {
		h.pub_counters[p] = 0
	}

	if h.pub_counters[p] == 0 {
		close(ch)
		delete(h.pubs, p)
	}

	return
}

func (h *defaultHub) transfer(path string, ch chan *sensord_pb.SensorData) {
	for {
		dat, ok := <-ch
		if !ok {
			h.glock.Lock()
			defer h.glock.Unlock()

			close(ch)
			delete(h.pubs, path)

			h.logger.WithField("path", path).Debugf("failed to recv data from channel, maybe closed")
			return
		}

		go func(dat *sensord_pb.SensorData) {
			subs := h.subs[path]
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

func (s *subscriber) Subscribe() (*sensord_pb.SensorData, error) {
	dat, ok := <-s.ch
	if !ok {
		return nil, hub.ErrUnsubscribable
	}
	return dat, nil
}

func (s *subscriber) Id() uint64 {
	return s.id
}

func (s *subscriber) Path() string {
	return s.p
}

type publisher struct {
	id uint64
	p  string
	ch chan *sensord_pb.SensorData
}

func (p *publisher) Publish(dat *sensord_pb.SensorData) error {
	p.ch <- dat
	return nil
}

func (p *publisher) Id() uint64 {
	return p.id
}

func (p *publisher) Path() string {
	return p.p
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
