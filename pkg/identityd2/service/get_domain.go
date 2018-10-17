package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) GetDomain(ctx context.Context, req *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	var dom_s *storage.Domain
	var err error

	err = req.Validate()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	dom := req.GetDomain()
	if dom.GetId() == nil {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := dom.GetId().GetValue()
	dom_s, err = self.storage.GetDomain(id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get domain in storage")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res := &pb.GetDomainResponse{
		Domain: copy_domain(dom_s),
	}

	self.logger.WithField("id", id).Debugf("get domain")

	return res, nil
}
