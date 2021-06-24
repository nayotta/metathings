package metathings_deviced_service

func (self *MetathingsDevicedService) doTerminating() error {
	if err := self.cc.Shutdown(); err != nil {
		return err
	}

	return nil
}
