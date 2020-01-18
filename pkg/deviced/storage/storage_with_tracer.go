package metathings_deviced_storage

import (
	"context"

	opentracing_storage_helper "github.com/nayotta/metathings/pkg/common/opentracing/storage"
)

type TracedStorage struct {
	*opentracing_storage_helper.BaseTracedStorage
	*StorageImpl
}

func (s *TracedStorage) CreateDevice(ctx context.Context, dev *Device) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateDevice")
	defer span.Finish()

	return s.StorageImpl.CreateDevice(ctx, dev)
}

func (s *TracedStorage) DeleteDevice(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteDevice")
	defer span.Finish()

	return s.StorageImpl.DeleteDevice(ctx, id)
}
func (s *TracedStorage) PatchDevice(ctx context.Context, id string, device *Device) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "patchDevice")
	defer span.Finish()

	return s.StorageImpl.PatchDevice(ctx, id, device)
}
func (s *TracedStorage) GetDevice(ctx context.Context, id string) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "GetDevice")
	defer span.Finish()

	return s.StorageImpl.GetDevice(ctx, id)
}
func (s *TracedStorage) ListDevices(ctx context.Context, dev *Device) ([]*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "ListDevices")
	defer span.Finish()

	return s.StorageImpl.ListDevices(ctx, dev)
}
func (s *TracedStorage) GetDeviceByModuleId(ctx context.Context, id string) (*Device, error) {
	span, ctx := s.TraceWrapper(ctx, "GetDeviceByModuleId")
	defer span.Finish()

	return s.StorageImpl.GetDeviceByModuleId(ctx, id)
}

func (s *TracedStorage) CreateModule(ctx context.Context, dev *Module) (*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateModule")
	defer span.Finish()

	return s.StorageImpl.CreateModule(ctx, dev)
}
func (s *TracedStorage) DeleteModule(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteModule")
	defer span.Finish()

	return s.StorageImpl.DeleteModule(ctx, id)
}
func (s *TracedStorage) PatchModule(ctx context.Context, id string, module *Module) (*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchModule")
	defer span.Finish()

	return s.StorageImpl.PatchModule(ctx, id, module)
}
func (s *TracedStorage) GetModule(ctx context.Context, id string) (*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "GetModule")
	defer span.Finish()

	return s.StorageImpl.GetModule(ctx, id)
}
func (s *TracedStorage) ListModules(ctx context.Context, mdl *Module) ([]*Module, error) {
	span, ctx := s.TraceWrapper(ctx, "ListModules")
	defer span.Finish()

	return s.StorageImpl.ListModules(ctx, mdl)
}

func (s *TracedStorage) CreateFlow(ctx context.Context, flw *Flow) (*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateFlow")
	defer span.Finish()

	return s.StorageImpl.CreateFlow(ctx, flw)
}
func (s *TracedStorage) DeleteFlow(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteFlow")
	defer span.Finish()

	return s.StorageImpl.DeleteFlow(ctx, id)
}
func (s *TracedStorage) PatchFlow(ctx context.Context, id string, flow *Flow) (*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchFlow")
	defer span.Finish()

	return s.StorageImpl.PatchFlow(ctx, id, flow)
}
func (s *TracedStorage) GetFlow(ctx context.Context, id string) (*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "GetFlow")
	defer span.Finish()

	return s.StorageImpl.GetFlow(ctx, id)
}
func (s *TracedStorage) ListFlows(ctx context.Context, flw *Flow) ([]*Flow, error) {
	span, ctx := s.TraceWrapper(ctx, "ListFlows")
	defer span.Finish()

	return s.StorageImpl.ListFlows(ctx, flw)
}

func (s *TracedStorage) CreateFlowSet(ctx context.Context, flwst *FlowSet) (*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateFlowSet")
	defer span.Finish()

	return s.StorageImpl.CreateFlowSet(ctx, flwst)
}
func (s *TracedStorage) DeleteFlowSet(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteFlowSet")
	defer span.Finish()

	return s.StorageImpl.DeleteFlowSet(ctx, id)
}
func (s *TracedStorage) PatchFlowSet(ctx context.Context, id string, flwst *FlowSet) (*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchFlowSet")
	defer span.Finish()

	return s.StorageImpl.PatchFlowSet(ctx, id, flwst)
}
func (s *TracedStorage) GetFlowSet(ctx context.Context, id string) (*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "GetFlowSet")
	defer span.Finish()

	return s.StorageImpl.GetFlowSet(ctx, id)
}
func (s *TracedStorage) ListFlowSets(ctx context.Context, flwsts *FlowSet) ([]*FlowSet, error) {
	span, ctx := s.TraceWrapper(ctx, "ListFlows")
	defer span.Finish()

	return s.StorageImpl.ListFlowSets(ctx, flwsts)
}
func (s *TracedStorage) AddFlowToFlowSet(ctx context.Context, flwst_id, flw_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddFlowToFlowSet")
	defer span.Finish()

	return s.StorageImpl.AddFlowToFlowSet(ctx, flwst_id, flw_id)
}
func (s *TracedStorage) RemoveFlowFromFlowSet(ctx context.Context, flwst_id, flw_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveFlowFromFlowSet")
	defer span.Finish()

	return s.StorageImpl.RemoveFlowFromFlowSet(ctx, flwst_id, flw_id)
}

func NewTracedStorage(s *StorageImpl) (Storage, error) {
	return &TracedStorage{
		BaseTracedStorage: opentracing_storage_helper.NewBaseTracedStorage(s),
		StorageImpl:       s,
	}, nil
}
