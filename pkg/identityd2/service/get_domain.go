package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) GetDomain(ctx context.Context, req *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	var dom *storage.Domain
	var err error

	err = req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := req.GetId().GetValue()
	dom, err = self.storage.GetDomain(id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get domain in storage")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res := &pb.GetDomainResponse{
		Domain: copy_domain(dom),
	}

	self.logger.WithField("id", id).Debugf("get domain")

	return res, nil
}
