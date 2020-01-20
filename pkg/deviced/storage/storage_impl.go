package metathings_deviced_storage

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	otgorm "github.com/smacker/opentracing-gorm"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage_helper "github.com/nayotta/metathings/pkg/common/storage"
)

type StorageImplOption struct {
	IsTraced bool
}

type StorageImpl struct {
	opt     *StorageImplOption
	root_db *gorm.DB
	logger  log.FieldLogger
}

func (self *StorageImpl) GetRootDBConn() *gorm.DB {
	return self.root_db
}

func (self *StorageImpl) GetDBConn(ctx context.Context) *gorm.DB {
	if db := ctx.Value("dbconn"); db != nil {
		return db.(*gorm.DB)
	}

	return self.GetRootDBConn()
}

func (self *StorageImpl) get_flow(ctx context.Context, id string) (*Flow, error) {
	var err error
	flw := &Flow{}

	if err = self.GetDBConn(ctx).First(flw, "id = ?", id).Error; err != nil {
		return nil, err
	}

	flw.Device = &Device{Id: flw.DeviceId}

	return flw, nil
}

func (self *StorageImpl) list_flow_sets(ctx context.Context, flwst *FlowSet) ([]*FlowSet, error) {
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

	if err = self.GetDBConn(ctx).Find(&flwsts_t, fs).Error; err != nil {
		return nil, err
	}

	for _, fs = range flwsts_t {
		if flwst, err = self.get_flow_set(ctx, *fs.Id); err != nil {
			return nil, err
		}
		flwsts = append(flwsts, flwst)
	}

	return flwsts, nil
}

func (self *StorageImpl) list_view_flow_sets_by_flow_id(ctx context.Context, id string) ([]*FlowSet, error) {
	var err error
	var flw_flwst_maps []*FlowFlowSetMapping
	var flwsts []*FlowSet

	if err = self.GetDBConn(ctx).Find(&flw_flwst_maps, "flow_id = ?", id).Error; err != nil {
		return nil, err
	}

	for _, m := range flw_flwst_maps {
		flwst := &FlowSet{
			Id: m.FlowSetId,
		}
		flwsts = append(flwsts, flwst)
	}

	return flwsts, nil
}

func (self *StorageImpl) internal_list_view_flows(ctx context.Context, flwst *FlowSet) ([]*Flow, error) {
	var err error
	var flws []*Flow
	var flw_flwst_maps []*FlowFlowSetMapping

	if err = self.GetDBConn(ctx).Find(&flw_flwst_maps, "flow_set_id = ?", *flwst.Id).Error; err != nil {
		return nil, err
	}

	for _, ffm := range flw_flwst_maps {
		flws = append(flws, &Flow{
			Id: ffm.FlowId,
		})
	}

	return flws, nil
}

func (self *StorageImpl) get_flow_set(ctx context.Context, id string) (*FlowSet, error) {
	var err error
	flwst := &FlowSet{}

	if err = self.GetDBConn(ctx).First(flwst, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if flwst.Flows, err = self.internal_list_view_flows(ctx, flwst); err != nil {
		return nil, err
	}

	return flwst, nil
}

func (self *StorageImpl) list_flows_by_device_id(ctx context.Context, id string) ([]*Flow, error) {
	var err error
	var flws_t []*Flow

	if err = self.GetDBConn(ctx).Select("id").Find(&flws_t, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	flws := []*Flow{}
	for _, f := range flws_t {
		if f, err = self.get_flow(ctx, *f.Id); err != nil {
			return nil, err
		}

		flws = append(flws, f)
	}

	return flws, nil
}

func (self *StorageImpl) list_flows(ctx context.Context, flw *Flow) ([]*Flow, error) {
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

	if err = self.GetDBConn(ctx).Select("id").Find(&flws_t, f).Error; err != nil {
		return nil, err
	}

	var flws []*Flow
	for _, f = range flws_t {
		if f, err = self.get_flow(ctx, *f.Id); err != nil {
			return nil, err
		}

		flws = append(flws, f)
	}

	return flws, nil
}

func (self *StorageImpl) get_module(ctx context.Context, id string) (*Module, error) {
	var err error
	mdl := &Module{}

	if err = self.GetDBConn(ctx).First(mdl, "id = ?", id).Error; err != nil {
		return nil, err
	}

	mdl.Device = &Device{Id: mdl.DeviceId}

	return mdl, nil
}

func (self *StorageImpl) list_modules_by_device_id(ctx context.Context, id string) ([]*Module, error) {
	var err error
	var mdls_t []*Module

	if err = self.GetDBConn(ctx).Select("id").Find(&mdls_t, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	mdls := []*Module{}
	for _, m := range mdls_t {
		if m, err = self.get_module(ctx, *m.Id); err != nil {
			return nil, err
		}

		mdls = append(mdls, m)
	}

	return mdls, nil
}

func (self *StorageImpl) list_modules(ctx context.Context, mdl *Module) ([]*Module, error) {
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

	if err = self.GetDBConn(ctx).Select("id").Find(&mdls_t, m).Error; err != nil {
		return nil, err
	}

	var mdls []*Module
	for _, m = range mdls_t {
		if m, err = self.get_module(ctx, *m.Id); err != nil {
			return nil, err
		}

		mdls = append(mdls, m)
	}

	return mdls, nil
}

func (self *StorageImpl) internal_get_device(ctx context.Context, dev *Device) (*Device, error) {
	var err error

	if dev.Modules, err = self.list_modules_by_device_id(ctx, *dev.Id); err != nil {
		return nil, err
	}

	if dev.Flows, err = self.list_flows_by_device_id(ctx, *dev.Id); err != nil {
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) get_device(ctx context.Context, id string) (*Device, error) {
	var err error
	dev := &Device{}

	if err = self.GetDBConn(ctx).First(dev, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if dev, err = self.internal_get_device(ctx, dev); err != nil {
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) get_device_by_module_id(ctx context.Context, id string) (*Device, error) {
	var err error
	var mdl Module

	if err = self.GetDBConn(ctx).Select("device_id").First(&mdl, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return self.get_device(ctx, *mdl.DeviceId)
}

func (self *StorageImpl) list_devices(ctx context.Context, dev *Device) ([]*Device, error) {
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

	if err = self.GetDBConn(ctx).Select("id").Find(&devs_t, d).Error; err != nil {
		return nil, err
	}

	var devs []*Device
	for _, d = range devs_t {
		if d, err = self.get_device(ctx, *d.Id); err != nil {
			return nil, err
		}

		devs = append(devs, d)
	}

	return devs, nil
}

func (self *StorageImpl) CreateDevice(ctx context.Context, dev *Device) (*Device, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(dev).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create device")
		return nil, err
	}

	if dev.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Device{Id: dev.Id}, dev.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra field")
		}
	}

	if dev, err = self.get_device(ctx, *dev.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", *dev.Id).Debugf("create device")

	return dev, nil
}

func (self *StorageImpl) DeleteDevice(ctx context.Context, id string) error {
	if err := self.GetDBConn(ctx).Delete(&Device{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete device")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete device")

	return nil
}

func (self *StorageImpl) PatchDevice(ctx context.Context, id string, device *Device) (*Device, error) {
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

	if err = self.GetDBConn(ctx).Model(&Device{Id: &id}).Update(dev).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch device")
		return nil, err
	}

	if device.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Device{Id: &id}, device.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra field")
		}
	}

	if device, err = self.get_device(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch device")

	return device, nil
}

func (self *StorageImpl) GetDevice(ctx context.Context, id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = self.get_device(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get device")

	return dev, nil
}

func (self *StorageImpl) ListDevices(ctx context.Context, dev *Device) ([]*Device, error) {
	var devs []*Device
	var err error

	if devs, err = self.list_devices(ctx, dev); err != nil {
		self.logger.WithError(err).Debugf("failed to list devices")
		return nil, err
	}

	self.logger.Debugf("list devices")

	return devs, nil
}

func (self *StorageImpl) GetDeviceByModuleId(ctx context.Context, id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = self.get_device_by_module_id(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device by module id")
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) CreateModule(ctx context.Context, mdl *Module) (*Module, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(mdl).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create module")
		return nil, err
	}

	if mdl, err = self.get_module(ctx, *mdl.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", *mdl.Id).Debugf("create module")

	return mdl, nil
}

func (self *StorageImpl) DeleteModule(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Module{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete module")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete module")

	return nil
}

func (self *StorageImpl) PatchModule(ctx context.Context, id string, mdl *Module) (*Module, error) {
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

	if err = self.GetDBConn(ctx).Model(&Module{Id: &id}).Update(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch module")
		return nil, err
	}

	if mdl, err = self.get_module(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch module")

	return mdl, nil
}

func (self *StorageImpl) GetModule(ctx context.Context, id string) (*Module, error) {
	var mdl *Module
	var err error

	if mdl, err = self.get_module(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get module")

	return mdl, nil
}

func (self *StorageImpl) ListModules(ctx context.Context, mdl *Module) ([]*Module, error) {
	var mdls []*Module
	var err error

	if mdls, err = self.list_modules(ctx, mdl); err != nil {
		self.logger.WithError(err).Debugf("failed to list modules")
		return nil, err
	}

	self.logger.Debugf("list modules")

	return mdls, nil
}

func (self *StorageImpl) CreateFlow(ctx context.Context, flw *Flow) (*Flow, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(flw).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create flow")
		return nil, err
	}

	if flw, err = self.get_flow(ctx, *flw.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", *flw.Id).Debugf("create flow")

	return flw, nil
}

func (self *StorageImpl) DeleteFlow(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Flow{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete flow")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete flow")

	return nil
}

func (self *StorageImpl) PatchFlow(ctx context.Context, id string, flw *Flow) (*Flow, error) {
	var err error
	var f Flow

	if flw.Alias != nil {
		f.Alias = flw.Alias
	}

	if err = self.GetDBConn(ctx).Model(&Flow{Id: &id}).Update(f).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch flow")
		return nil, err
	}

	if flw, err = self.get_flow(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch flow")

	return flw, nil
}

func (self *StorageImpl) GetFlow(ctx context.Context, id string) (*Flow, error) {
	var flw *Flow
	var err error

	if flw, err = self.get_flow(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get flow")

	return flw, nil
}

func (self *StorageImpl) ListFlows(ctx context.Context, flw *Flow) ([]*Flow, error) {
	var flws []*Flow
	var err error

	if flws, err = self.list_flows(ctx, flw); err != nil {
		self.logger.WithError(err).Debugf("failed to list flows")
		return nil, err
	}

	self.logger.Debugf("list flows")

	return flws, nil
}

func (self *StorageImpl) CreateFlowSet(ctx context.Context, flwst *FlowSet) (*FlowSet, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(flwst).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create flow set")
		return nil, err
	}

	if flwst, err = self.get_flow_set(ctx, *flwst.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow set")
		return nil, err
	}

	self.logger.WithField("id", *flwst.Id).Debugf("create flow set")

	return flwst, nil
}

func (self *StorageImpl) DeleteFlowSet(ctx context.Context, id string) error {
	if err := self.GetDBConn(ctx).Delete(&FlowSet{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete flow set")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete flow set")

	return nil
}

func (self *StorageImpl) PatchFlowSet(ctx context.Context, id string, flwst *FlowSet) (*FlowSet, error) {
	var err error
	var fs FlowSet

	if flwst.Alias != nil {
		fs.Alias = flwst.Alias
	}

	if err = self.GetDBConn(ctx).Model(&FlowSet{Id: &id}).Update(fs).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch flow set")
		return nil, err
	}

	if flwst, err = self.get_flow_set(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow set")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch flow set")

	return flwst, nil
}

func (self *StorageImpl) GetFlowSet(ctx context.Context, id string) (*FlowSet, error) {
	var flwst *FlowSet
	var err error

	if flwst, err = self.get_flow_set(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow set")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get flow set")

	return flwst, nil
}

func (self *StorageImpl) ListFlowSets(ctx context.Context, flwst *FlowSet) ([]*FlowSet, error) {
	var flwsts []*FlowSet
	var err error

	if flwsts, err = self.list_flow_sets(ctx, flwst); err != nil {
		self.logger.WithError(err).Debugf("failed to list flow sets")
		return nil, err
	}

	self.logger.Debugf("list flow sets")

	return flwsts, nil
}

func (self *StorageImpl) ListViewFlowSetsByFlowId(ctx context.Context, id string) ([]*FlowSet, error) {
	var flwsts []*FlowSet
	var err error

	if flwsts, err = self.list_view_flow_sets_by_flow_id(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to list flow sets by flow id")
		return nil, err
	}

	self.logger.Debugf("list flow sets by flow id")

	return flwsts, nil
}

func (self *StorageImpl) AddFlowToFlowSet(ctx context.Context, flwst_id, flw_id string) error {
	m := &FlowFlowSetMapping{
		FlowSetId: &flwst_id,
		FlowId:    &flw_id,
	}

	if err := self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add flow to flow set")
		return err
	}

	self.logger.WithFields(log.Fields{
		"flow_id":     flw_id,
		"flow_set_id": flwst_id,
	})

	return nil
}

func (self *StorageImpl) RemoveFlowFromFlowSet(ctx context.Context, flwst_id, flw_id string) error {
	if err := self.GetDBConn(ctx).Delete(&FlowFlowSetMapping{}, "flow_set_id = ? and flow_id = ?", flwst_id, flw_id).Error; err != nil {
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

	if s.opt.IsTraced {
		otgorm.AddGormCallbacks(db)
	}

	s.root_db = db

	return nil
}

func init_db(s *StorageImpl) error {
	s.GetRootDBConn().AutoMigrate(
		&Device{},
		&Module{},
		&Flow{},
		&FlowSet{},
		&FlowFlowSetMapping{},
	)

	return nil
}

func NewStorageImpl(driver, uri string, args ...interface{}) (Storage, error) {
	var err error
	var ok bool
	var opt StorageImplOption
	var logger log.FieldLogger

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
		"tracer": func(key string, val interface{}) error {
			var tracer opentracing.Tracer

			if tracer, ok = val.(opentracing.Tracer); !ok {
				return opt_helper.InvalidArgument(key)
			}

			if tracer != nil {
				opt.IsTraced = true
			}

			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	s := &StorageImpl{
		opt:    &opt,
		logger: logger,
	}

	if err = new_db(s, driver, uri); err != nil {
		return nil, err
	}

	if err = init_db(s); err != nil {
		return nil, err
	}

	if s.opt.IsTraced {
		return NewTracedStorage(s)
	}

	return s, nil
}
