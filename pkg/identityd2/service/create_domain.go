package metathings_identityd2_service

import (
	"context"
	"errors"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) CreateDomain(ctx context.Context, req *pb.CreateDomainRequest) (*pb.CreateDomainResponse, error) {
	var dom *storage.Domain
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	parent_id := req.GetParent().GetId().GetValue()
	if parent_id == "" {
		err = errors.New("parent.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := req.GetId().GetValue()
	if id == "" {
		id = id_helper.NewId()
	}

	extra_str := must_parse_extra(req.GetExtra())
	name_str := req.Name.Value
	alias_str := req.Alias.Value

	if err = self.enforcer.AddObjectToKind(id, KIND_DOMAIN); err != nil {
		self.logger.WithError(err).Errorf("failed to add domain to kind in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	dom = &storage.Domain{
		Id:       &id,
		Name:     &name_str,
		Alias:    &alias_str,
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
