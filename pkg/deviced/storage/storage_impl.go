package metathings_deviced_storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *StorageImpl) get_flow(id string) (*Flow, error) {
	var err error
	flw := &Flow{}

	if err = self.db.First(flw, "id = ?", id).Error; err != nil {
		return nil, err
	}

	flw.Device = &Device{Id: flw.DeviceId}

	return flw, nil
}

func (self *StorageImpl) list_flow_sets(flwst *FlowSet) ([]*FlowSet, error) {
	var err error
	var flwsts_t []*FlowSet
	var flwsts []*FlowSet

	fs := &FlowSet{}
	if flwst.Id != nil {
		fs.Id = flwst.Id
	}

	if flwst.Alias != nil {
		fs.Alias = flwst.Alias
	}

	if flwst.Name != nil {
		fs.Name = flwst.Name
	}

	if err = self.db.Find(&flwsts_t, fs).Error; err != nil {
		return nil, err
	}

	for _, fs = range flwsts_t {
		if flwst, err = self.get_flow_set(*fs.Id); err != nil {
			return nil, err
		}
		flwsts = append(flwsts, flwst)
	}

	return flwsts, nil
}

func (self *StorageImpl) internal_list_view_flows(flwst *FlowSet) ([]*Flow, error) {
	var err error
	var flws []*Flow
	var flw_flwst_maps []*FlowFlowSetMapping

	if err = self.db.Find(&flw_flwst_maps, "flow_set_id = ?", *flwst.Id).Error; err != nil {
		return nil, err
	}

	for _, ffm := range flw_flwst_maps {
		flws = append(flws, &Flow{
			Id: ffm.FlowId,
		})
	}

	return flws, nil
}

func (self *StorageImpl) get_flow_set(id string) (*FlowSet, error) {
	var err error
	flwst := &FlowSet{}

	if err = self.db.First(flwst, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if flwst.Flows, err = self.internal_list_view_flows(flwst); err != nil {
		return nil, err
	}

	return flwst, nil
}

func (self *StorageImpl) list_flows_by_device_id(id string) ([]*Flow, error) {
	var err error
	var flws_t []*Flow

	if err = self.db.Select("id").Find(&flws_t, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	flws := []*Flow{}
	for _, f := range flws_t {
		if f, err = self.get_flow(*f.Id); err != nil {
			return nil, err
		}

		flws = append(flws, f)
	}

	return flws, nil
}

func (self *StorageImpl) list_flows(flw *Flow) ([]*Flow, error) {
	var err error
	var flws_t []*Flow

	f := &Flow{}
	if flw.Id != nil {
		f.Id = flw.Id
	}

	if flw.DeviceId != nil {
		f.DeviceId = flw.DeviceId
	}

	if flw.Name != nil {
		f.Name = flw.Name
	}

	if flw.Alias != nil {
		f.Alias = flw.Alias
	}

	if err = self.db.Select("id").Find(&flws_t, f).Error; err != nil {
		return nil, err
	}

	var flws []*Flow
	for _, f = range flws_t {
		if f, err = self.get_flow(*f.Id); err != nil {
			return nil, err
		}

		flws = append(flws, f)
	}

	return flws, nil
}

func (self *StorageImpl) get_module(id string) (*Module, error) {
	var err error
	mdl := &Module{}

	if err = self.db.First(mdl, "id = ?", id).Error; err != nil {
		return nil, err
	}

	mdl.Device = &Device{Id: mdl.DeviceId}

	return mdl, nil
}

func (self *StorageImpl) list_modules_by_device_id(id string) ([]*Module, error) {
	var err error
	var mdls_t []*Module

	if err = self.db.Select("id").Find(&mdls_t, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	mdls := []*Module{}
	for _, m := range mdls_t {
		if m, err = self.get_module(*m.Id); err != nil {
			return nil, err
		}

		mdls = append(mdls, m)
	}

	return mdls, nil
}

func (self *StorageImpl) list_modules(mdl *Module) ([]*Module, error) {
	var err error
	var mdls_t []*Module

	m := &Module{}
	if mdl.Id != nil {
		m.Id = mdl.Id
	}

	if mdl.State != nil {
		m.State = mdl.State
	}

	if mdl.DeviceId != nil {
		m.DeviceId = mdl.DeviceId
	}

	if mdl.Endpoint != nil {
		m.Endpoint = mdl.Endpoint
	}

	if mdl.Component != nil {
		m.Component = mdl.Component
	}

	if mdl.Name != nil {
		m.Name = mdl.Name
	}

	if mdl.Alias != nil {
		m.Alias = mdl.Alias
	}

	if err = self.db.Select("id").Find(&mdls_t, m).Error; err != nil {
		return nil, err
	}

	var mdls []*Module
	for _, m = range mdls_t {
		if m, err = self.get_module(*m.Id); err != nil {
			return nil, err
		}

		mdls = append(mdls, m)
	}

	return mdls, nil
}

func (self *StorageImpl) internal_get_device(dev *Device) (*Device, error) {
	var err error

	if dev.Modules, err = self.list_modules_by_device_id(*dev.Id); err != nil {
		return nil, err
	}

	if dev.Flows, err = self.list_flows_by_device_id(*dev.Id); err != nil {
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) get_device(id string) (*Device, error) {
	var err error
	dev := &Device{}

	if err = self.db.First(dev, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if dev, err = self.internal_get_device(dev); err != nil {
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) get_device_by_module_id(id string) (*Device, error) {
	var err error
	var mdl Module

	if err = self.db.Select("device_id").First(&mdl, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return self.get_device(*mdl.DeviceId)
}

func (self *StorageImpl) list_devices(dev *Device) ([]*Device, error) {
	var err error
	var devs_t []*Device

	d := &Device{}
	if dev.Id != nil {
		d.Id = dev.Id
	}

	if dev.Kind != nil {
		d.Kind = dev.Kind
	}

	if dev.State != nil {
		d.State = dev.State
	}

	if dev.Name != nil {
		d.Name = dev.Name
	}

	if dev.Alias != nil {
		d.Alias = dev.Alias
	}

	if err = self.db.Select("id").Find(&devs_t, d).Error; err != nil {
		return nil, err
	}

	var devs []*Device
	for _, d = range devs_t {
		if d, err = self.get_device(*d.Id); err != nil {
			return nil, err
		}

		devs = append(devs, d)
	}

	return devs, nil
}

func (self *StorageImpl) CreateDevice(dev *Device) (*Device, error) {
	var err error

	if err = self.db.Create(dev).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create device")
		return nil, err
	}

	if dev, err = self.get_device(*dev.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", *dev.Id).Debugf("create device")

	return dev, nil
}

func (self *StorageImpl) DeleteDevice(id string) error {
	if err := self.db.Delete(&Device{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete device")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete device")

	return nil
}

func (self *StorageImpl) PatchDevice(id string, device *Device) (*Device, error) {
	var err error
	var dev Device

	if device.Alias != nil {
		dev.Alias = device.Alias
	}

	if device.State != nil {
		dev.State = device.State
	}

	if device.HeartbeatAt != nil {
		dev.HeartbeatAt = device.HeartbeatAt
	}

	if err = self.db.Model(&Device{Id: &id}).Update(dev).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch device")
		return nil, err
	}

	if device, err = self.get_device(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch device")

	return device, nil
}

func (self *StorageImpl) GetDevice(id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = self.get_device(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get device")

	return dev, nil
}

func (self *StorageImpl) ListDevices(dev *Device) ([]*Device, error) {
	var devs []*Device
	var err error

	if devs, err = self.list_devices(dev); err != nil {
		self.logger.WithError(err).Debugf("failed to list devices")
		return nil, err
	}

	self.logger.Debugf("list devices")

	return devs, nil
}

func (self *StorageImpl) GetDeviceByModuleId(id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = self.get_device_by_module_id(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device by module id")
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) CreateModule(mdl *Module) (*Module, error) {
	var err error

	if err = self.db.Create(mdl).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create module")
		return nil, err
	}

	if mdl, err = self.get_module(*mdl.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", *mdl.Id).Debugf("create module")

	return mdl, nil
}

func (self *StorageImpl) DeleteModule(id string) error {
	var err error

	if err = self.db.Delete(&Module{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete module")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete module")

	return nil
}

func (self *StorageImpl) PatchModule(id string, mdl *Module) (*Module, error) {
	var err error
	var m Module

	if mdl.Alias != nil {
		m.Alias = mdl.Alias
	}

	if mdl.State != nil {
		m.State = mdl.State
	}

	if mdl.HeartbeatAt != nil {
		m.HeartbeatAt = mdl.HeartbeatAt
	}

	if err = self.db.Model(&Module{Id: &id}).Update(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch module")
		return nil, err
	}

	if mdl, err = self.get_module(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch module")

	return mdl, nil
}

func (self *StorageImpl) GetModule(id string) (*Module, error) {
	var mdl *Module
	var err error

	if mdl, err = self.get_module(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get module")

	return mdl, nil
}

func (self *StorageImpl) ListModules(mdl *Module) ([]*Module, error) {
	var mdls []*Module
	var err error

	if mdls, err = self.list_modules(mdl); err != nil {
		self.logger.WithError(err).Debugf("failed to list modules")
		return nil, err
	}

	self.logger.Debugf("list modules")

	return mdls, nil
}

func (self *StorageImpl) CreateFlow(flw *Flow) (*Flow, error) {
	var err error

	if err = self.db.Create(flw).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create flow")
		return nil, err
	}

	if flw, err = self.get_flow(*flw.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", *flw.Id).Debugf("create flow")

	return flw, nil
}

func (self *StorageImpl) DeleteFlow(id string) error {
	var err error

	if err = self.db.Delete(&Flow{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete flow")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete flow")

	return nil
}

func (self *StorageImpl) PatchFlow(id string, flw *Flow) (*Flow, error) {
	var err error
	var f Flow

	if flw.Alias != nil {
		f.Alias = flw.Alias
	}

	if err = self.db.Model(&Flow{Id: &id}).Update(f).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch flow")
		return nil, err
	}

	if flw, err = self.get_flow(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch flow")

	return flw, nil
}

func (self *StorageImpl) GetFlow(id string) (*Flow, error) {
	var flw *Flow
	var err error

	if flw, err = self.get_flow(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get flow")

	return flw, nil
}

func (self *StorageImpl) ListFlows(flw *Flow) ([]*Flow, error) {
	var flws []*Flow
	var err error

	if flws, err = self.list_flows(flw); err != nil {
		self.logger.WithError(err).Debugf("failed to list flows")
		return nil, err
	}

	self.logger.Debugf("list flows")

	return flws, nil
}

func (self *StorageImpl) CreateFlowSet(flwst *FlowSet) (*FlowSet, error) {
	var err error

	if err = self.db.Create(flwst).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create flow set")
		return nil, err
	}

	if flwst, err = self.get_flow_set(*flwst.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow set")
		return nil, err
	}

	self.logger.WithField("id", *flwst.Id).Debugf("create flow set")

	return flwst, nil
}

func (self *StorageImpl) DeleteFlowSet(id string) error {
	if err := self.db.Delete(&FlowSet{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete flow set")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete flow set")

	return nil
}

func (self *StorageImpl) PatchFlowSet(id string, flwst *FlowSet) (*FlowSet, error) {
	var err error
	var fs FlowSet

	if flwst.Alias != nil {
		fs.Alias = flwst.Alias
	}

	if err = self.db.Model(&FlowSet{Id: &id}).Update(fs).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch flow set")
		return nil, err
	}

	if flwst, err = self.get_flow_set(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow set")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch flow set")

	return flwst, nil
}

func (self *StorageImpl) GetFlowSet(id string) (*FlowSet, error) {
	var flwst *FlowSet
	var err error

	if flwst, err = self.get_flow_set(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow set")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get flow set")

	return flwst, nil
}

func (self *StorageImpl) ListFlowSets(flwst *FlowSet) ([]*FlowSet, error) {
	var flwsts []*FlowSet
	var err error

	if flwsts, err = self.list_flow_sets(flwst); err != nil {
		self.logger.WithError(err).Debugf("failed to list flow sets")
		return nil, err
	}

	self.logger.Debugf("list flow sets")

	return flwsts, nil
}

func (self *StorageImpl) AddFlowToFlowSet(flwst_id, flw_id string) error {
	m := &FlowFlowSetMapping{
		FlowSetId: &flwst_id,
		FlowId:    &flw_id,
	}

	if err := self.db.Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add flow to flow set")
		return err
	}

	self.logger.WithFields(log.Fields{
		"flow_id":     flw_id,
		"flow_set_id": flwst_id,
	})

	return nil
}

func (self *StorageImpl) RemoveFlowFromFlowSet(flwst_id, flw_id string) error {
	if err := self.db.Delete(&FlowFlowSetMapping{}, "flow_set_id = ? and flow_id = ?", flwst_id, flw_id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to remove flow from flow set")
		return err
	}

	self.logger.WithFields(log.Fields{
		"flow_set_id": flwst_id,
		"flow_id":     flw_id,
	})

	return nil
}

func init_args(s *StorageImpl, args ...interface{}) error {
	var key string
	var ok bool

	if len(args)%2 != 0 {
		return InvalidArgument
	}

	for i := 0; i < len(args); i += 2 {
		key, ok = args[i].(string)
		if !ok {
			return InvalidArgument
		}

		switch key {
		case "logger":
			s.logger, ok = args[i+1].(log.FieldLogger)
			if !ok {
				return InvalidArgument
			}
		}
	}

	return nil
}

func new_db(s *StorageImpl, driver, uri string) error {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(driver, uri); err != nil {
		return err
	}
	s.db = db

	return nil
}

func init_db(s *StorageImpl) error {
	s.db.AutoMigrate(
		&Device{},
		&Module{},
		&Flow{},
		&FlowSet{},
	)

	return nil
}

func NewStorageImpl(driver, uri string, args ...interface{}) (*StorageImpl, error) {
	var err error

	s := &StorageImpl{}
	if err = init_args(s, args...); err != nil {
		return nil, err
	}

	if err = new_db(s, driver, uri); err != nil {
		return nil, err
	}

	if err = init_db(s); err != nil {
		return nil, err
	}

	return s, nil
}
