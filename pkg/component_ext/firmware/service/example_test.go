package metathings_component_ext_firmware_service_test

import (
	component "github.com/nayotta/metathings/pkg/component"
	component_ext_firmware_service "github.com/nayotta/metathings/pkg/component_ext/firmware/service"
)

type exampleService struct {
	*component_ext_firmware_service.ComponentExtFirmwareService
}

func (s *exampleService) InitModuleService(m *component.Module) error {
	var err error
	s.ComponentExtFirmwareService, err = component_ext_firmware_service.NewComponentExtFirmwareService(m)
	if err != nil {
		return err
	}

	// do other initialize things.

	return nil
}
