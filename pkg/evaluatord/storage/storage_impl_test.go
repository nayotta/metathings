package metathings_evaluatord_storage

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	test_helper "github.com/nayotta/metathings/pkg/common/test"
)

var (
	test_evaluator_id          = "test_evaluator_id"
	test_evaluator_alias       = "test_evaluator_alias"
	test_evaluator_description = "test evaluator description"
	test_evaluator_config      = "{}"
	test_source_id             = "test_source_id"
	test_source_type           = "test_source_type"
	test_operator_id           = "test_operator_id"
	test_operator_alias        = "test_operator_alias"
	test_operator_description  = "test operator description"
	test_operator_driver       = "lua"
	test_lua_descriptor_code   = `return "hello, world"`

	test_evaluator = &Evaluator{
		Id:          &test_evaluator_id,
		Alias:       &test_evaluator_alias,
		Description: &test_evaluator_description,
		Config:      &test_evaluator_config,
		Sources: []*Resource{
			&Resource{
				Id:   &test_source_id,
				Type: &test_source_type,
			},
		},
		Operator: &Operator{
			Id:          &test_operator_id,
			Alias:       &test_operator_alias,
			Description: &test_operator_description,
			Driver:      &test_operator_driver,
			LuaDescriptor: &LuaDescriptor{
				Code: &test_lua_descriptor_code,
			},
		},
	}
)

type StorageImplTestSuite struct {
	suite.Suite
	stor *StorageImpl
	ctx  context.Context
}

func (s *StorageImplTestSuite) SetupTest() {
	stor, err := NewStorageImpl(test_helper.GetTestGormDriver(), test_helper.GetTestGormUri(), "logger", log.New())
	s.Nil(err)

	s.stor = stor.(*StorageImpl)
	s.ctx = context.Background()

	_, err = s.stor.CreateEvaluator(s.ctx, test_evaluator)
	s.Nil(err)
}

func (s *StorageImplTestSuite) TestGetEvaluator() {
	evltr, err := s.stor.GetEvaluator(s.ctx, test_evaluator_id)
	s.Nil(err)

	s.Equal(test_evaluator_id, *evltr.Id)
	s.Equal(test_evaluator_alias, *evltr.Alias)
	s.Equal(test_evaluator_description, *evltr.Description)
	s.Equal(test_evaluator_config, *evltr.Config)
	s.Len(evltr.Sources, 1)
	s.Equal(test_source_id, *evltr.Sources[0].Id)
	s.Equal(test_source_type, *evltr.Sources[0].Type)
	s.Equal(test_operator_id, *evltr.Operator.Id)
	s.Equal(test_operator_alias, *evltr.Operator.Alias)
	s.Equal(test_operator_description, *evltr.Operator.Description)
	s.Equal(test_operator_driver, *evltr.Operator.Driver)
	s.Equal(test_evaluator_id, *evltr.Operator.EvaluatorId)
	s.Equal(test_lua_descriptor_code, *evltr.Operator.LuaDescriptor.Code)
}

func (s *StorageImplTestSuite) TestDeleteEvaluator() {
	err := s.stor.DeleteEvaluator(s.ctx, test_evaluator_id)
	s.Nil(err)

	_, err = s.stor.GetEvaluator(s.ctx, test_evaluator_id)
	s.NotNil(err)
}

func (s *StorageImplTestSuite) TestPatchEvaluator() {
	new_evaluator_alias := "new_evaluator_alias"
	new_evaluator_description := "new evaluator description"
	new_operator_alias := "new_operator_alias"
	new_operator_description := "new operator description"
	new_lua_descriptor_code := `return "yes, my lord"`
	evltr, err := s.stor.PatchEvaluator(s.ctx, test_evaluator_id, &Evaluator{
		Alias:       &new_evaluator_alias,
		Description: &new_evaluator_description,
		Operator: &Operator{
			Alias:       &new_operator_alias,
			Description: &new_operator_description,
			LuaDescriptor: &LuaDescriptor{
				Code: &new_lua_descriptor_code,
			},
		},
	})
	s.Nil(err)

	s.Equal(new_evaluator_alias, *evltr.Alias)
	s.Equal(new_evaluator_description, *evltr.Description)
	s.Equal(new_operator_alias, *evltr.Operator.Alias)
	s.Equal(new_operator_description, *evltr.Operator.Description)
	s.Equal(new_lua_descriptor_code, *evltr.Operator.LuaDescriptor.Code)
}

func (s *StorageImplTestSuite) TestAddSourcesToEvaluator() {
	new_source1_id := "new_source1_id"
	new_source1_type := "new_source1_type"
	new_source2_id := "new_source2_id"
	new_source2_type := "new_source2_type"

	err := s.stor.AddSourcesToEvaluator(s.ctx, test_evaluator_id, []*Resource{
		&Resource{
			Id:   &new_source1_id,
			Type: &new_source1_type,
		},
		&Resource{
			Id:   &new_source2_id,
			Type: &new_source2_type,
		},
	})
	s.Nil(err)

	evltr, err := s.stor.GetEvaluator(s.ctx, test_evaluator_id)
	s.Nil(err)

	s.Len(evltr.Sources, 3)
	s.Equal(new_source1_id, *evltr.Sources[1].Id)
	s.Equal(new_source1_type, *evltr.Sources[1].Type)
	s.Equal(new_source2_id, *evltr.Sources[2].Id)
	s.Equal(new_source2_type, *evltr.Sources[2].Type)
}

func (s *StorageImplTestSuite) TestRemoveSourcesFromEvaluator() {
	err := s.stor.RemoveSourcesFromEvaluator(s.ctx, test_evaluator_id, []*Resource{
		&Resource{
			Id:   &test_source_id,
			Type: &test_source_type,
		},
	})
	s.Nil(err)

	evltr, err := s.stor.GetEvaluator(s.ctx, test_evaluator_id)
	s.Nil(err)
	s.Len(evltr.Sources, 0)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(StorageImplTestSuite))
}
