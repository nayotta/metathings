package metathings_deviced_storage

import (
	"context"

	opentracing_storage_helper "github.com/nayotta/metathings/pkg/common/opentracing/storage"
)

type TracedStorage struct {
	*opentracing_storage_helper.BaseTracedStorage
	Storage
}

func (s *TracedStorage) CreateDevice(ctx context.Context, dev *Device) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateDevice")
	defer span.Finish()

	return s.Storage.CreateDevice(ctx, dev)
}

func (s *TracedStorage) DeleteDevice(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteDevice")
	defer span.Finish()

	return s.Storage.DeleteDevice(ctx, id)
}
func (s *TracedStorage) PatchDevice(ctx context.Context, id string, device *Device) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "patchDevice")
	defer span.Finish()

	return s.Storage.PatchDevice(ctx, id, device)
}
func (s *TracedStorage) GetDevice(ctx context.Context, id string) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "GetDevice")
	defer span.Finish()

	return s.Storage.GetDevice(ctx, id)
}
func (s *TracedStorage) ListDevices(ctx context.Context, dev *Device) ([]*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "ListDevices")
	defer span.Finish()

	return s.Storage.ListDevices(ctx, dev)
}
func (s *TracedStorage) GetDeviceByModuleId(ctx context.Context, id string) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "GetDeviceByModuleId")
	defer span.Finish()

	return s.Storage.GetDeviceByModuleId(ctx, id)
}

func (s *TracedStorage) CreateModule(ctx context.Context, dev *Module) (*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateModule")
	defer span.Finish()

	return s.Storage.CreateModule(ctx, dev)
}
func (s *TracedStorage) DeleteModule(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteModule")
	defer span.Finish()

	return s.Storage.DeleteModule(ctx, id)
}
func (s *TracedStorage) PatchModule(ctx context.Context, id string, module *Module) (*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchModule")
	defer span.Finish()

	return s.Storage.PatchModule(ctx, id, module)
}
func (s *TracedStorage) GetModule(ctx context.Context, id string) (*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "GetModule")
	defer span.Finish()

	return s.Storage.GetModule(ctx, id)
}
func (s *TracedStorage) ListModules(ctx context.Context, mdl *Module) ([]*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "ListModules")
	defer span.Finish()

	return s.Storage.ListModules(ctx, mdl)
}

func (s *TracedStorage) CreateFlow(ctx context.Context, flw *Flow) (*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateFlow")
	defer span.Finish()

	return s.Storage.CreateFlow(ctx, flw)
}
func (s *TracedStorage) DeleteFlow(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteFlow")
	defer span.Finish()

	return s.Storage.DeleteFlow(ctx, id)
}
func (s *TracedStorage) PatchFlow(ctx context.Context, id string, flow *Flow) (*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchFlow")
	defer span.Finish()

	return s.Storage.PatchFlow(ctx, id, flow)
}
func (s *TracedStorage) GetFlow(ctx context.Context, id string) (*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "GetFlow")
	defer span.Finish()

	return s.Storage.GetFlow(ctx, id)
}
func (s *TracedStorage) ListFlows(ctx context.Context, flw *Flow) ([]*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "ListFlows")
	defer span.Finish()

	return s.Storage.ListFlows(ctx, flw)
}

func (s *TracedStorage) CreateFlowSet(ctx context.Context, flwst *FlowSet) (*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateFlowSet")
	defer span.Finish()

	return s.Storage.CreateFlowSet(ctx, flwst)
}

func (s *TracedStorage) DeleteFlowSet(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteFlowSet")
	defer span.Finish()

	return s.Storage.DeleteFlowSet(ctx, id)
}

func (s *TracedStorage) PatchFlowSet(ctx context.Context, id string, flwst *FlowSet) (*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchFlowSet")
	defer span.Finish()

	return s.Storage.PatchFlowSet(ctx, id, flwst)
}

func (s *TracedStorage) GetFlowSet(ctx context.Context, id string) (*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "GetFlowSet")
	defer span.Finish()

	return s.Storage.GetFlowSet(ctx, id)
}

func (s *TracedStorage) ListFlowSets(ctx context.Context, flwsts *FlowSet) ([]*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "ListFlows")
	defer span.Finish()

	return s.Storage.ListFlowSets(ctx, flwsts)
}

func (s *TracedStorage) AddFlowToFlowSet(ctx context.Context, flwst_id, flw_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddFlowToFlowSet")
	defer span.Finish()

	return s.Storage.AddFlowToFlowSet(ctx, flwst_id, flw_id)
}

func (s *TracedStorage) RemoveFlowFromFlowSet(ctx context.Context, flwst_id, flw_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveFlowFromFlowSet")
	defer span.Finish()

	return s.Storage.RemoveFlowFromFlowSet(ctx, flwst_id, flw_id)
}

func (s *TracedStorage) CreateConfig(ctx context.Context, cfg *Config) (*Config, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateConfig")
	defer span.Finish()

	return s.Storage.CreateConfig(ctx, cfg)
}

func (s *TracedStorage) DeleteConfig(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteConfig")
	defer span.Finish()

	return s.Storage.DeleteConfig(ctx, id)
}

func (s *TracedStorage) PatchConfig(ctx context.Context, id string, cfg *Config) (*Config, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchConfig")
	defer span.Finish()

	return s.Storage.PatchConfig(ctx, id, cfg)
}

func (s *TracedStorage) GetConfig(ctx context.Context, id string) (*Config, error) {
	span, ctx := s.TraceWrapper(ctx, "GetConfig")
	defer span.Finish()

	return s.Storage.GetConfig(ctx, id)
}

func (s *TracedStorage) ListConfigs(ctx context.Context, cfg *Config) ([]*Config, error) {
	span, ctx := s.TraceWrapper(ctx, "ListConfig")
	defer span.Finish()

	return s.Storage.ListConfigs(ctx, cfg)
}

func (s *TracedStorage) AddConfigToDevice(ctx context.Context, dev_id, cfg_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddConfigToDevice")
	defer span.Finish()

	return s.Storage.AddConfigToDevice(ctx, dev_id, cfg_id)
}

func (s *TracedStorage) RemoveConfigFromDevice(ctx context.Context, dev_id, cfg_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveConfigFromDevice")
	defer span.Finish()

	return s.Storage.RemoveConfigFromDevice(ctx, dev_id, cfg_id)
}

func (s *TracedStorage) RemoveConfigFromDeviceByConfigId(ctx context.Context, cfg_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveConfigFromDeviceByConfigId")
	defer span.Finish()

	return s.Storage.RemoveConfigFromDeviceByConfigId(ctx, cfg_id)
}

func (s *TracedStorage) ListConfigsByDeviceId(ctx context.Context, dev_id string) ([]*Config, error) {
	span, ctx := s.TraceWrapper(ctx, "ListConfigsByDeviceId")
	defer span.Finish()

	return s.Storage.ListConfigsByDeviceId(ctx, dev_id)
}

func NewTracedStorage(s Storage, getter opentracing_storage_helper.RootDBConnGetter) (Storage, error) {
	return &TracedStorage{
		BaseTracedStorage: opentracing_storage_helper.NewBaseTracedStorage(getter),
		Storage:           s,
	}, nil
}
