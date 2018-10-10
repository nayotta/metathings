package metathings_identityd2_service

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) CreateDomain(ctx context.Context, req *pb.CreateDomainRequest) (*pb.CreateDomainResponse, error) {
	var dom *storage.Domain
	var buf []byte
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	parent_id := req.GetParent().GetId().GetValue()
	if parent_id == "" {
		err = errors.New("parent.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := req.GetId().GetValue()
	if id == "" {
		id = id_helper.NewId()
	}

	extra_map := pb_helper.ExtractStringMap(req.GetExtra())
	if buf, err = json.Marshal(extra_map); err != nil {
		err = errors.New("extra is bad argument")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	extra_str := string(buf)

	dom = &storage.Domain{
		Id:       &id,
		Name:     &req.Name.Value,
		Alias:    &req.Alias.Value,
		ParentId: &parent_id,
		Extra:    &extra_str,
	}

	if dom, err = self.storage.CreateDomain(dom); err != nil {
		self.logger.WithError(err).Errorf("failed to create domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateDomainResponse{
		Domain: copy_domain(dom),
	}

	self.logger.WithField("id", *dom.Id).Infof("create domain")

	return res, nil
}
