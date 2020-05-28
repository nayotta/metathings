package metathings_evaluatord_storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
)

var (
	testdata = []struct {
		TaskId     string
		SourceId   string
		SourceType string
	}{
		{"task_id_0", "source_id_0", "test"},
		{"task_id_1", "source_id_0", "test"},
		{"task_id_2", "source_id_1", "test"},
	}
)

type Influxdb2TaskStorageTestSuite struct {
	suite.Suite
	tstor *Influxdb2TaskStorage
	ctx   context.Context
}

func (s *Influxdb2TaskStorageTestSuite) SetupTest() {
	logger, err := log_helper.NewLogger("test", "debug")
	s.Require().Nil(err)

	ts, err := NewInfluxdb2TaskStorage(
		"logger", logger,
		"address", test_helper.GetTestInfluxdb2Address(),
		"token", test_helper.GetTestInfluxdb2Token(),
		"org", test_helper.GetTestInfluxdb2Org(),
		"bucket", test_helper.GetTestInfluxdb2Bucket(),
	)
	s.Require().Nil(err)

	s.tstor = ts.(*Influxdb2TaskStorage)
	s.ctx = context.TODO()
}

func (s *Influxdb2TaskStorageTestSuite) BeforeTest(suiteName, testName string) {
	map[string]func(){
		"TestAllInOne": s.setupTestAllInOne,
	}[testName]()
}

func (s *Influxdb2TaskStorageTestSuite) setupTestAllInOne() {
	st := "created"
	err := s.tstor.PatchTask(s.ctx, &Task{
		Id: &testdata[0].TaskId,
		Source: &Resource{
			Id:   &testdata[0].SourceId,
			Type: &testdata[0].SourceType,
		},
	}, &TaskState{
		State: &st,
	})
	s.Require().Nil(err)

	st = "running"
	err = s.tstor.PatchTask(s.ctx, &Task{
		Id: &testdata[0].TaskId,
		Source: &Resource{
			Id:   &testdata[0].SourceId,
			Type: &testdata[0].SourceType,
		},
	}, &TaskState{
		State: &st,
	})
	s.Require().Nil(err)

	st = "done"
	err = s.tstor.PatchTask(s.ctx, &Task{
		Id: &testdata[0].TaskId,
		Source: &Resource{
			Id:   &testdata[0].SourceId,
			Type: &testdata[0].SourceType,
		},
	}, &TaskState{
		State: &st,
	})
	s.Require().Nil(err)
}

func (s *Influxdb2TaskStorageTestSuite) TestAllInOne() {
	tsk, err := s.tstor.GetTask(s.ctx, testdata[0].TaskId)
	s.Nil(err)
	s.Require().Len(tsk.States, 3)
	s.Equal(testdata[0].TaskId, *tsk.Id)
	s.Equal(testdata[0].SourceId, *tsk.Source.Id)
	s.Equal(testdata[0].SourceType, *tsk.Source.Type)
	s.Equal("done", *tsk.CurrentState.State)
	s.Equal("created", *tsk.States[0].State)
	s.Equal("running", *tsk.States[1].State)
	s.Equal("done", *tsk.States[2].State)

	tsks, err := s.tstor.ListTasksBySource(s.ctx, &Resource{
		Id:   &testdata[0].SourceId,
		Type: &testdata[0].SourceType,
	})
	s.Nil(err)
	s.Len(tsks, 1)

	tsk, err = s.tstor.GetTask(s.ctx, testdata[1].TaskId)
	s.Equal(err, ErrTaskNotFound)
	s.Nil(tsk)

	tsks, err = s.tstor.ListTasksBySource(s.ctx, &Resource{
		Id:   &testdata[2].SourceId,
		Type: &testdata[2].SourceType,
	})
	s.Nil(err)
	s.Len(tsks, 0)

	st := "created"
	err = s.tstor.PatchTask(s.ctx, &Task{
		Id: &testdata[1].TaskId,
		Source: &Resource{
			Id:   &testdata[1].SourceId,
			Type: &testdata[1].SourceType,
		},
	}, &TaskState{
		State: &st,
	})
	s.Require().Nil(err)

	tsk, err = s.tstor.GetTask(s.ctx, testdata[1].TaskId)
	s.Nil(err)
	s.Require().Len(tsk.States, 1)
	s.Equal(testdata[1].TaskId, *tsk.Id)
	s.Equal(testdata[1].SourceId, *tsk.Source.Id)
	s.Equal(testdata[1].SourceType, *tsk.Source.Type)
	s.Equal("created", *tsk.CurrentState.State)
	s.Equal("created", *tsk.States[0].State)

	st = "running"
	err = s.tstor.PatchTask(s.ctx, &Task{
		Id: &testdata[1].TaskId,
		Source: &Resource{
			Id:   &testdata[1].SourceId,
			Type: &testdata[1].SourceType,
		},
	}, &TaskState{
		State: &st,
	})
	s.Require().Nil(err)

	st = "error"
	err = s.tstor.PatchTask(s.ctx, &Task{
		Id: &testdata[1].TaskId,
		Source: &Resource{
			Id:   &testdata[1].SourceId,
			Type: &testdata[1].SourceType,
		},
	}, &TaskState{
		State: &st,
	})
	s.Require().Nil(err)

	tsk, err = s.tstor.GetTask(s.ctx, testdata[1].TaskId)
	s.Nil(err)
	s.Require().Len(tsk.States, 3)
	s.Equal(testdata[1].TaskId, *tsk.Id)
	s.Equal(testdata[1].SourceId, *tsk.Source.Id)
	s.Equal(testdata[1].SourceType, *tsk.Source.Type)
	s.Equal("error", *tsk.CurrentState.State)
	s.Equal("created", *tsk.States[0].State)
	s.Equal("running", *tsk.States[1].State)
	s.Equal("error", *tsk.States[2].State)

	tsks, err = s.tstor.ListTasksBySource(s.ctx, &Resource{
		Id:   &testdata[1].SourceId,
		Type: &testdata[1].SourceType,
	})
	s.Require().Nil(err)
	s.Len(tsks, 2)

	tsks, err = s.tstor.ListTasksBySource(s.ctx, &Resource{
		Id:   &testdata[2].SourceId,
		Type: &testdata[2].SourceType,
	})
	s.Nil(err)
	s.Len(tsks, 0)
}

func TestInfluxdb2TaskStorageTestSuite(t *testing.T) {
	suite.Run(t, new(Influxdb2TaskStorageTestSuite))
}
