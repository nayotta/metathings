package metathings_streamd_storage

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

var (
	test_stream0_id       = "test-stream0-id"
	test_stream0_name     = "test-stream0-name"
	test_stream0_owner_id = "test-stream0-owner-id"
	test_stream0_state    = "stop"

	test_stream1_id       = "test-stream1-id"
	test_stream1_name     = "test-stream1-name"
	test_stream1_owner_id = "test-stream1-owner-id"
	test_stream1_state    = "stop"

	test_source0_id       = "test-source0-id"
	test_upstream0_id     = "test-upstream0-id"
	test_upstream0_name   = "test-upstream0-name"
	test_upstream0_alias  = "test-upstream0-alias"
	test_upstream0_state  = "stop"
	test_upstream0_config = "{}"

	test_group0_id      = "test-group0-id"
	test_input0_id      = "test-input0-id"
	test_input0_name    = "test-input0-name"
	test_input0_alias   = "test-input0-alias"
	test_input0_state   = "stop"
	test_input0_config  = "{}"
	test_output0_id     = "test-output0-id"
	test_output0_name   = "test-output0-name"
	test_output0_alias  = "test-output0-alias"
	test_output0_state  = "stop"
	test_output0_config = "{}"
)

type storageImplTestSuite struct {
	suite.Suite

	storage *storageImpl
	stream0 Stream
	stream1 Stream
}

func (self *storageImplTestSuite) SetupTest() {
	self.storage, _ = newStorageImpl("sqlite3", ":memory:", logrus.New())

	self.stream0 = Stream{
		Id:      &test_stream0_id,
		Name:    &test_stream0_name,
		OwnerId: &test_stream0_owner_id,
		State:   &test_stream0_state,
	}

	self.stream1 = Stream{
		Id:      &test_stream1_id,
		Name:    &test_stream1_name,
		OwnerId: &test_stream1_owner_id,
		State:   &test_stream1_state,
		Sources: []Source{
			Source{
				Id: &test_source0_id,
				Upstream: Upstream{
					Id:     &test_upstream0_id,
					Name:   &test_upstream0_name,
					Alias:  &test_upstream0_alias,
					State:  &test_upstream0_state,
					Config: &test_upstream0_config,
				},
			},
		},
		Groups: []Group{
			Group{
				Id: &test_group0_id,
				Inputs: []Input{
					Input{
						Id:     &test_input0_id,
						Name:   &test_input0_name,
						Alias:  &test_input0_alias,
						State:  &test_input0_state,
						Config: &test_input0_config,
					},
				},
				Outputs: []Output{
					Output{
						Id:     &test_output0_id,
						Name:   &test_output0_name,
						Alias:  &test_output0_alias,
						State:  &test_output0_state,
						Config: &test_output0_config,
					},
				},
			},
		},
	}

	self.storage.CreateStream(self.stream0)
}

func (self *storageImplTestSuite) TestInit() {
	var count int
	err := self.storage.db.Model(&Stream{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)
}

func (self *storageImplTestSuite) TestGetStream() {
	stm, err := self.storage.GetStream(*self.stream0.Id)
	self.Nil(err)
	self.Equal(*self.stream0.Id, *stm.Id)
}

func (self *storageImplTestSuite) SkipCreateStream() {
	stm, err := self.storage.CreateStream(self.stream1)
	self.Nil(err)
	self.Equal(*self.stream1.Id, *stm.Id)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
