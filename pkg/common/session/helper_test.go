package session_helper

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SessionHelperTestSuite struct {
	suite.Suite
}

func (s *SessionHelperTestSuite) TestMajorSession() {
	sess := NewSession(GenerateStartupSession(), GenerateMajorSession())
	s.True(IsMajorSession(sess))
	s.False(IsMinorSession(sess))
	s.False(IsTempSession(sess))
}

func (s *SessionHelperTestSuite) TestMinorSession() {
	sess := NewSession(GenerateStartupSession(), GenerateMinorSession())
	s.False(IsMajorSession(sess))
	s.True(IsMinorSession(sess))
	s.False(IsTempSession(sess))
}

func (s *SessionHelperTestSuite) TestTempSession() {
	sess := NewSession(GenerateStartupSession(), GenerateTempSession())
	s.False(IsMajorSession(sess))
	s.True(IsMinorSession(sess))
	s.True(IsTempSession(sess))
}

func TestSessionHelperTestSuite(t *testing.T) {
	suite.Run(t, new(SessionHelperTestSuite))
}
