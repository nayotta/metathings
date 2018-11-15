package policy_helper

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func ensure_helper(interface{}) error {
	return nil
}

type validatorTestSuite struct {
	suite.Suite
}

func (self *validatorTestSuite) TestFalse() {
	err := ValidateChain([]interface{}{
		func() (interface{}, error) {
			return nil, errors.New("should not pass")
		},
	}, []interface{}{ensure_helper})

	self.Error(err)
}

func (self *validatorTestSuite) TestLongRequest() {
	err := ValidateChain([]interface{}{
		func() (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return nil, nil
		},
	}, []interface{}{ensure_helper})

	self.Nil(err)
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(validatorTestSuite))
}
