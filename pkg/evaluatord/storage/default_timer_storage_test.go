package metathings_evaluatord_storage

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	test_helper "github.com/nayotta/metathings/pkg/common/test"
)

var (
	test_timer_id          = "test-timer-id"
	test_timer_alias       = "test-timer-alias"
	test_timer_description = "test-timer-description"
	test_timer_schedule    = "test-timer-schedule"
	test_timer_timezone    = "test-timer-timezone"
	test_timer_enabled     = true
	test_timer_configs     = []string{"config-id1", "config-id2"}
	test_timer             = &Timer{
		Id:          &test_timer_id,
		Alias:       &test_timer_alias,
		Description: &test_timer_description,
		Schedule:    &test_timer_schedule,
		Timezone:    &test_timer_timezone,
		Enabled:     &test_timer_enabled,
	}

	new_test_timer_id          = "new-test-timer-id"
	new_test_timer_alias       = "new-test-timer-alias"
	new_test_timer_description = "new-test-timer-description"
	new_test_timer_schedule    = "new-test-timer-schedule"
	new_test_timer_timezone    = "new-test-timer-timezone"
	new_test_timer_enabled     = false
	new_test_timer_configs     = []string{"config-id1", "config-id3", "config-id4"}
	new_test_timer             = &Timer{
		Id:          &new_test_timer_id,
		Alias:       &new_test_timer_alias,
		Description: &new_test_timer_description,
		Schedule:    &new_test_timer_schedule,
		Timezone:    &new_test_timer_timezone,
		Enabled:     &new_test_timer_enabled,
	}
)

type DefaultTimerStorageTestSuite struct {
	suite.Suite
	s   *DefaultTimerStorage
	ctx context.Context
}

func (s *DefaultTimerStorageTestSuite) SetupTest() {
	s.ctx = context.TODO()
	stor, err := NewDefaultTimerStorage(
		"gorm_driver", test_helper.GetTestGormDriver(),
		"gorm_uri", test_helper.GetTestGormUri(),
		"logger", logrus.New(),
	)
	s.Require().Nil(err)
	s.s = stor.(*DefaultTimerStorage)

	_, err = s.s.CreateTimer(s.ctx, test_timer)
	s.Require().Nil(err)

	err = s.s.AddConfigsToTimer(s.ctx, test_timer_id, test_timer_configs)
	s.Require().Nil(err)
}

func (s *DefaultTimerStorageTestSuite) TestGetTimer() {
	t, err := s.s.GetTimer(s.ctx, test_timer_id)
	s.Require().Nil(err)

	s.Equal(test_timer_id, *t.Id)
	s.Equal(test_timer_alias, *t.Alias)
	s.Equal(test_timer_description, *t.Description)
	s.Equal(test_timer_schedule, *t.Schedule)
	s.Equal(test_timer_timezone, *t.Timezone)
	s.Equal(test_timer_enabled, *t.Enabled)
	s.Len(t.Configs, len(test_timer_configs))
	for _, cfg_id := range test_timer_configs {
		s.Contains(t.Configs, cfg_id)
	}
}

func (s *DefaultTimerStorageTestSuite) TestDeleteTimer() {
	err := s.s.DeleteTimer(s.ctx, test_timer_id)
	s.Require().Nil(err)

	_, err = s.s.GetTimer(s.ctx, test_timer_id)
	s.NotNil(err)
}

func (s *DefaultTimerStorageTestSuite) TestListTimers() {
	ts, err := s.s.ListTimers(s.ctx, &Timer{
		Id: &test_timer_id,
	})
	s.Require().Nil(err)
	s.Len(ts, 1)
}

func (s *DefaultTimerStorageTestSuite) TestCreateTimer() {
	t, err := s.s.CreateTimer(s.ctx, new_test_timer)
	s.Require().Nil(err)

	s.Equal(new_test_timer_id, *t.Id)
	s.Equal(new_test_timer_alias, *t.Alias)
	s.Equal(new_test_timer_description, *t.Description)
	s.Equal(new_test_timer_schedule, *t.Schedule)
	s.Equal(new_test_timer_timezone, *t.Timezone)
	s.Equal(new_test_timer_enabled, *t.Enabled)
	s.Len(t.Configs, 0)

	ts, err := s.s.ListTimers(s.ctx, &Timer{})
	s.Require().Nil(err)
	s.Len(ts, 2)

	err = s.s.AddConfigsToTimer(s.ctx, new_test_timer_id, new_test_timer_configs)
	s.Require().Nil(err)

	t, err = s.s.GetTimer(s.ctx, new_test_timer_id)
	s.Len(t.Configs, len(new_test_timer_configs))
	for _, cfg_id := range new_test_timer_configs {
		s.Contains(t.Configs, cfg_id)
	}
}

func (s *DefaultTimerStorageTestSuite) TestPatchTimer() {
	t, err := s.s.PatchTimer(s.ctx, test_timer_id, &Timer{
		Alias:       &new_test_timer_alias,
		Description: &new_test_timer_description,
		Schedule:    &new_test_timer_schedule,
		Timezone:    &new_test_timer_timezone,
		Enabled:     &new_test_timer_enabled,
	})
	s.Require().Nil(err)

	s.Equal(new_test_timer_alias, *t.Alias)
	s.Equal(new_test_timer_description, *t.Description)
	s.Equal(new_test_timer_schedule, *t.Schedule)
	s.Equal(new_test_timer_timezone, *t.Timezone)
	s.Equal(new_test_timer_enabled, *t.Enabled)
}

func (s *DefaultTimerStorageTestSuite) TestRemoveConfigsFromTimer() {
	err := s.s.RemoveConfigsFromTimer(s.ctx, test_timer_id, []string{"config-id1"})
	s.Require().Nil(err)

	t, err := s.s.GetTimer(s.ctx, test_timer_id)
	s.Require().Nil(err)
	s.Require().Len(t.Configs, 1)
	s.Contains(t.Configs, "config-id2")
	s.NotContains(t.Configs, "config-id1")
}

func (s *DefaultTimerStorageTestSuite) TestAddConfigsToTimer() {
	err := s.s.AddConfigsToTimer(s.ctx, test_timer_id, []string{"config-id3"})
	s.Require().Nil(err)

	t, err := s.s.GetTimer(s.ctx, test_timer_id)
	s.Require().Nil(err)
	s.Require().Len(t.Configs, 3)
	s.Contains(t.Configs, "config-id3")
}

func TestDefaultTimerStorageTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultTimerStorageTestSuite))
}
