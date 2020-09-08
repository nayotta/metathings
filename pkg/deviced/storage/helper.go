package metathings_deviced_storage

import "github.com/jinzhu/gorm"

func parse_error(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return RecordNotFound
	default:
		return err
	}
}

func wrap_module(mdl *Module) *Module {
	if mdl.DeviceId != nil {
		mdl.Device = &Device{Id: mdl.DeviceId}
	}

	return mdl
}

func wrap_flow(flw *Flow) *Flow {
	if flw.DeviceId != nil {
		flw.Device = &Device{Id: flw.DeviceId}
	}

	return flw
}
