package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.GetPolicyResponse, error) {
	var plc_s *storage.Policy
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	plc := req.GetPolicy()
	if plc.GetId() == nil {
		err = errors.New("policy.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	id_str := plc.GetId().GetValue()

	if plc_s, err = self.storage.GetPolicy(id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get policy in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetPolicyResponse{
		Policy: copy_policy(plc_s),
	}

	self.logger.WithField("id", id_str).Debugf("get policy")

	return res, nil
}
