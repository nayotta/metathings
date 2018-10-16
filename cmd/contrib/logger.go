package cmd_contrib

import (
	log "github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
)

type LoggerOptioner interface {
	GetLevelP() *string
	GetLevel() string
	SetLevel(string)
}

type LoggerOption struct {
	Log struct {
		Level string
	}
}

func (self *LoggerOption) GetLevelP() *string {
	return &self.Log.Level
}

func (self *LoggerOption) GetLevel() string {
	return self.Log.Level
}

func (self *LoggerOption) SetLevel(lvl string) {
	self.Log.Level = lvl
}

func NewLogger(service string) func(LoggerOptioner) (log.FieldLogger, error) {
	return func(opt LoggerOptioner) (log.FieldLogger, error) {
		return log_helper.NewLogger(service, opt.GetLevel())
	}
}
