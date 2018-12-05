package cmd_contrib

type HeartbeatOptioner interface {
	GetIntervalP() *int
	GetInterval() int
	SetInterval(int)
}

type HeartbeatOption struct {
	Interval int
}

func (self *HeartbeatOption) GetIntervalP() *int {
	return &self.Interval
}

func (self *HeartbeatOption) GetInterval() int {
	return self.Interval
}

func (self *HeartbeatOption) SetInterval(ivt int) {
	self.Interval = ivt
}
