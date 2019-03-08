package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) DeleteAction(ctx context.Context, req *pb.DeleteActionRequest) (*empty.Empty, error) {
	var err error

	act := req.GetAction()
	act_id_str := act.GetId().GetValue()

	if err = self.storage.DeleteAction(act_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete action in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithField("id", act_id_str).Infof("delete action")

	return &empty.Empty{}, nil
}
