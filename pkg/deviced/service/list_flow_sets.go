package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ListFlowSets(ctx context.Context, req *pb.ListFlowSetsRequest) (*pb.ListFlowSetsResponse, error) {
	var flwsts_s []*storage.FlowSet
	var err error

	flwst := req.GetFlowSet()
	flwst_s := &storage.FlowSet{}

	id := flwst.GetId()
	if id != nil {
		flwst_s.Id = &id.Value
	}

	name := flwst.GetName()
	if name != nil {
		flwst_s.Name = &name.Value
	}

	alias := flwst.GetAlias()
	if alias != nil {
		flwst_s.Alias = &alias.Value
	}

	if flwsts_s, err = self.storage.ListFlowSets(flwst_s); err != nil {
		self.logger.WithError(err).Errorf("failed to list flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListFlowSetsResponse{
		FlowSets: copy_flow_sets(flwsts_s),
	}

	self.logger.Debugf("list flow sets")

	return res, nil
}
