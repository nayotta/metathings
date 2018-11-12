package metathings_identityd2_service

import (
	"context"
	"errors"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) PatchDomain(ctx context.Context, req *pb.PatchDomainRequest) (*pb.PatchDomainResponse, error) {
	var dom = &storage.Domain{}
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if req.GetId() == nil || req.GetId().GetValue() == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	idStr := req.GetId().GetValue()

	if req.GetAlias() != nil {
		aliasStr := req.GetAlias().GetValue()
		dom.Alias = &aliasStr
	}
	if req.GetExtra() != nil {
		extraStr := must_parse_extra(req.GetExtra())
		dom.Extra = &extraStr
	}

	if dom, err = self.storage.PatchDomain(idStr, dom); err != nil {
		self.logger.WithError(err).Errorf("failed to patch domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchDomainResponse{
		Domain: copy_domain(dom),
	}

	self.logger.WithField("id", idStr).Infof("patch domain")

	return res, nil
}
