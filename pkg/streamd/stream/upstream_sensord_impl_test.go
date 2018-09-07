package stream_manager

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testFilterUpstreamDataSuite struct {
	suite.Suite
	filter_bad          string
	upstream_data_empty *UpstreamData

	upstream_data0 *UpstreamData
	filter0        string
	filter1        string
	filter2        string
}

func (self *testFilterUpstreamDataSuite) SetupTest() {
	self.filter_bad = "x x x"
	self.upstream_data_empty = NewUpstreamData(
		map[string]interface{}{},
		map[string]interface{}{},
	)

	self.upstream_data0 = NewUpstreamData(
		map[string]interface{}{
			"value": "1",
		},
		map[string]interface{}{
			"sensor_id":   "id",
			"sensor_name": "name",
			"created_at":  "0",
			"arrvied_at":  "0",
		},
	)
	self.filter0 = `metadata.sensor_name=="name"`
	self.filter1 = `tonumber(value)==1`
	self.filter2 = `tonumber(value)<1`
}

func (self *testFilterUpstreamDataSuite) TestBadFilter() {
	_, err := filter_upstream_data(self.filter_bad, self.upstream_data_empty)
	self.Error(err)
}

func (self *testFilterUpstreamDataSuite) TestPassFilter0() {
	pass, err := filter_upstream_data(self.filter0, self.upstream_data0)
	self.Nil(err)
	self.True(pass)
}

func (self *testFilterUpstreamDataSuite) TestPassFilter1() {
	pass, err := filter_upstream_data(self.filter1, self.upstream_data0)
	self.Nil(err)
	self.True(pass)
}

func (self *testFilterUpstreamDataSuite) TestNotPassFilter2() {
	pass, err := filter_upstream_data(self.filter2, self.upstream_data0)
	self.Nil(err)
	self.False(pass)
}

func TestFilterUpstreamData(t *testing.T) {
	suite.Run(t, new(testFilterUpstreamDataSuite))
}
