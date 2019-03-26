package metathings_device_service

import (
	session_helper "github.com/nayotta/metathings/pkg/common/session"
)

func (self *MetathingsDeviceServiceImpl) generator_major_session() int64 {
	return int64(self.startup_session)<<32 | int64(session_helper.GenerateMajorSession())
}

func (self *MetathingsDeviceServiceImpl) generator_minor_session() int64 {
	return int64(self.startup_session)<<32 | int64(session_helper.GenerateMinorSession())
}
