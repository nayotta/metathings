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
	test_input0_config  = "{}"
	test_output0_id     = "test-output0-id"
	test_output0_name   = "test-output0-name"
	test_output0_alias  = "test-output0-alias"
	test_output0_config = "{}"

	test_stream2_id       = "test-stream2-id"
	test_stream2_name     = "test-stream2-name"
	test_stream2_owner_id = "test-stream2-owner-id"
	test_stream2_state    = "stop"

	test_source1_id       = "test-soruce1-id"
	test_upstream1_id     = "test-upstream1-id"
	test_upstream1_name   = "test-upstream1-name"
	test_upstream1_alias  = "test-upstream1-alias"
	test_upstream1_state  = "stop"
	test_upstream1_config = "{}"

	test_group1_id      = "test-group1-id"
	test_input1_id      = "test-input1-id"
	test_input1_name    = "test-intpu1-name"
	test_input1_alias   = "test-input1-alias"
	test_input1_config  = "{}"
	test_output1_id     = "test-output1-id"
	test_output1_name   = "test-output1-name"
	test_output1_alias  = "test-output1-alias"
	test_output1_config = "{}"
)

type storageImplTestSuite struct {
	suite.Suite

	storage *storageImpl
	stream0 Stream
	stream1 Stream
	stream2 Stream
}

func (self *storageImplTestSuite) assert_equal_stream(x, y Stream) {
	self.Equal(*x.Id, *y.Id)
	self.Equal(*x.Name, *y.Name)
	self.Equal(*x.OwnerId, *y.OwnerId)
	self.Equal(*x.State, *y.State)

	self.assert_equal_sources(x.Sources, y.Sources)
	self.assert_equal_groups(x.Groups, y.Groups)
}

func (self *storageImplTestSuite) assert_equal_sources(xs, ys []Source) {
	for _, x := range xs {
		ok := false
	assert_equal_sources_loop:
		for _, y := range ys {
			if *x.Id == *y.Id {
				self.assert_equal_source(x, y)
				ok = true
				break assert_equal_sources_loop
			}
		}
		self.True(ok)
	}
}

func (self *storageImplTestSuite) assert_equal_groups(xs, ys []Group) {
	self.Equal(len(xs), len(ys))
	for _, x := range xs {
		ok := false
	assert_equal_groups_loop:
		for _, y := range ys {
			if *x.Id == *y.Id {
				self.assert_equal_group(x, y)
				ok = true
				break assert_equal_groups_loop
			}
		}
		self.True(ok)
	}
}

func (self *storageImplTestSuite) assert_equal_source(x, y Source) {
	self.Equal(*x.Id, *y.Id)
	self.Equal(*x.StreamId, *y.StreamId)

	self.assert_equal_upstream(x.Upstream, y.Upstream)
}

func (self *storageImplTestSuite) assert_equal_group(x, y Group) {
	self.Equal(*x.Id, *y.Id)
	self.Equal(*x.StreamId, *y.StreamId)

	self.assert_equal_inputs(x.Inputs, y.Inputs)
	self.assert_equal_outputs(x.Outputs, y.Outputs)
}

func (self *storageImplTestSuite) assert_equal_upstream(x, y Upstream) {
	self.Equal(*x.Id, *y.Id)
	self.Equal(*x.SourceId, *y.SourceId)
	self.Equal(*x.Name, *y.Name)
	self.Equal(*x.Alias, *y.Alias)
	self.Equal(*x.Config, *y.Config)
}

func (self *storageImplTestSuite) assert_equal_inputs(xs, ys []Input) {
	for _, x := range xs {
		ok := false
	assert_equal_inputs_loop:
		for _, y := range ys {
			if *x.Id == *y.Id {
				self.assert_equal_input(x, y)
				ok = true
				break assert_equal_inputs_loop
			}
		}
		self.True(ok)
	}
}

func (self *storageImplTestSuite) assert_equal_outputs(xs, ys []Output) {
	for _, x := range xs {
		ok := false
	assert_equal_outputs_loop:
		for _, y := range ys {
			if *x.Id == *y.Id {
				self.assert_equal_output(x, y)
				ok = true
				break assert_equal_outputs_loop
			}
		}
		self.True(ok)
	}
}

func (self *storageImplTestSuite) assert_equal_input(x, y Input) {
	self.Equal(*x.Id, *y.Id)
	self.Equal(*x.GroupId, *y.GroupId)
	self.Equal(*x.Name, *y.Name)
	self.Equal(*x.Alias, *y.Alias)
	self.Equal(*y.Config, *y.Config)
}

func (self *storageImplTestSuite) assert_equal_output(x, y Output) {
	self.Equal(*x.Id, *y.Id)
	self.Equal(*x.GroupId, *y.GroupId)
	self.Equal(*x.Name, *y.Name)
	self.Equal(*x.Alias, *y.Alias)
	self.Equal(*y.Config, *y.Config)
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
				Id:       &test_source0_id,
				StreamId: &test_stream1_id,
				Upstream: Upstream{
					Id:       &test_upstream0_id,
					SourceId: &test_source0_id,
					Name:     &test_upstream0_name,
					Alias:    &test_upstream0_alias,
					Config:   &test_upstream0_config,
				},
			},
		},
		Groups: []Group{
			Group{
				Id:       &test_group0_id,
				StreamId: &test_stream1_id,
				Inputs: []Input{
					Input{
						Id:      &test_input0_id,
						GroupId: &test_group0_id,
						Name:    &test_input0_name,
						Alias:   &test_input0_alias,
						Config:  &test_input0_config,
					},
				},
				Outputs: []Output{
					Output{
						Id:      &test_output0_id,
						GroupId: &test_group0_id,
						Name:    &test_output0_name,
						Alias:   &test_output0_alias,
						Config:  &test_output0_config,
					},
				},
			},
		},
	}

	self.stream2 = Stream{
		Id:      &test_stream2_id,
		Name:    &test_stream2_name,
		OwnerId: &test_stream2_owner_id,
		State:   &test_stream2_state,
		Sources: []Source{
			Source{
				Id:       &test_source1_id,
				StreamId: &test_stream2_id,
				Upstream: Upstream{
					Id:       &test_upstream1_id,
					SourceId: &test_source1_id,
					Name:     &test_upstream1_name,
					Alias:    &test_upstream1_alias,
					Config:   &test_upstream1_config,
				},
			},
		},
		Groups: []Group{
			Group{
				Id:       &test_group1_id,
				StreamId: &test_stream2_id,
				Inputs: []Input{
					Input{
						Id:      &test_input1_id,
						GroupId: &test_group1_id,
						Name:    &test_input1_name,
						Alias:   &test_input1_alias,
						Config:  &test_input1_config,
					},
				},
				Outputs: []Output{
					Output{
						Id:      &test_output1_id,
						GroupId: &test_group1_id,
						Name:    &test_output1_name,
						Alias:   &test_output1_alias,
						Config:  &test_output1_config,
					},
				},
			},
		},
	}

	self.storage.CreateStream(self.stream0)
	self.storage.CreateStream(self.stream1)
}

func (self *storageImplTestSuite) TestInit() {
	var count int
	err := self.storage.db.Model(&Stream{}).Count(&count).Error
	self.Nil(err)
	self.Equal(2, count)
}

func (self *storageImplTestSuite) TestGetStream() {
	stm, err := self.storage.GetStream(*self.stream0.Id)
	self.Nil(err)
	self.assert_equal_stream(self.stream0, stm)

	stm, err = self.storage.GetStream(*self.stream1.Id)
	self.Nil(err)
	self.assert_equal_stream(self.stream1, stm)
}

func (self *storageImplTestSuite) TestCreateStream() {
	stm, err := self.storage.CreateStream(self.stream2)
	self.Nil(err)
	self.assert_equal_stream(self.stream2, stm)
}

func (self *storageImplTestSuite) TestDeleteStream() {
	err := self.storage.DeleteStream(*self.stream0.Id)
	self.Nil(err)

	var count int
	err = self.storage.db.Model(&Stream{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)

	err = self.storage.db.Model(&Source{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)

	err = self.storage.db.Model(&Group{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)

	err = self.storage.db.Model(&Upstream{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)

	err = self.storage.db.Model(&Input{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)

	err = self.storage.db.Model(&Output{}).Count(&count).Error
	self.Nil(err)
	self.Equal(1, count)

	err = self.storage.DeleteStream(*self.stream1.Id)
	self.Nil(err)

	err = self.storage.db.Model(&Stream{}).Count(&count).Error
	self.Nil(err)
	self.Equal(0, count)

	err = self.storage.db.Model(&Source{}).Count(&count).Error
	self.Nil(err)
	self.Equal(0, count)

	err = self.storage.db.Model(&Group{}).Count(&count).Error
	self.Nil(err)
	self.Equal(0, count)

	err = self.storage.db.Model(&Upstream{}).Count(&count).Error
	self.Nil(err)
	self.Equal(0, count)

	err = self.storage.db.Model(&Input{}).Count(&count).Error
	self.Nil(err)
	self.Equal(0, count)

	err = self.storage.db.Model(&Output{}).Count(&count).Error
	self.Nil(err)
	self.Equal(0, count)

}

func (self *storageImplTestSuite) TestPatchStream() {
	new_name := "new-name"
	new_state := "running"
	stm := Stream{
		Name:  &new_name,
		State: &new_state,
	}
	stm, err := self.storage.PatchStream(*self.stream0.Id, stm)
	self.Nil(err)
	self.Equal(new_name, *stm.Name)
	self.Equal(new_state, *stm.State)
}

func (self *storageImplTestSuite) TestListStreams() {
	stms, err := self.storage.ListStreams(Stream{})
	self.Nil(err)
	self.Len(stms, 2)

	streams := []Stream{self.stream0, self.stream1}
	for _, x := range stms {
		ok := false
		for _, y := range streams {
			if *x.Id == *y.Id {
				self.assert_equal_stream(x, y)
				ok = true
			}
		}
		self.True(ok)
	}
}

func (self *storageImplTestSuite) TestListStreamsWithFilter() {
	stms, err := self.storage.ListStreams(Stream{
		Name: &test_stream0_name,
	})
	self.Nil(err)
	self.Len(stms, 1)
	self.assert_equal_stream(stms[0], self.stream0)
}

func (self *storageImplTestSuite) TestListStreamsForUser() {
	stms, err := self.storage.ListStreamsForUser(test_stream0_owner_id, Stream{})
	self.Nil(err)
	self.Len(stms, 1)
	self.assert_equal_stream(stms[0], self.stream0)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
