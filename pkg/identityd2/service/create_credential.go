package metathings_identityd2_service

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) CreateCredential(ctx context.Context, req *pb.CreateCredentialRequest) (*pb.CreateCredentialResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id_str := id_helper.NewId()
	if req.GetId() != nil && req.GetId().GetValue() != "" {
		id_str = req.GetId().GetValue()
	}

	dom := req.GetDomain()
	if dom.GetId() == nil || dom.GetId().GetValue() == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom.GetId().GetValue()

	roles := []*storage.Role{}
	for _, r := range req.GetRoles() {
		if r.GetId() == nil || r.GetId().GetValue() == "" {
			err = errors.New("role.id is empty")
			self.logger.WithError(err).Warningf("failed to validate request data")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		roles = append(roles, &storage.Role{
			Id: &r.Id.Value,
		})
	}

	ent := req.GetEntity()
	if ent.GetId() == nil || ent.GetId().GetValue() == "" {
		err = errors.New("enity.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_id_str := ent.GetId().GetValue()

	alias_str := req.Name.Value
	if req.GetAlias() != nil {
		alias_str = req.GetAlias().GetValue()
	}

	srt_str := generate_secret()
	if req.GetSecret() != nil {
		srt_str = req.GetSecret().GetValue()
	}
	srt_str = passwd_helper.MustParsePassword(srt_str)

	desc_str := ""
	if req.GetDescription() != nil {
		desc_str = req.GetDescription().GetValue()
	}

	var expires time.Time
	if req.GetExpiresAt() != nil {
		expires = pb_helper.ToTime(*req.GetExpiresAt())
	}

	cred := &storage.Credential{
		Id:          &id_str,
		DomainId:    &dom_id_str,
		EntityId:    &ent_id_str,
		Name:        &req.Name.Value,
		Alias:       &alias_str,
		Secret:      &srt_str,
		Description: &desc_str,
		ExpiresAt:   &expires,
		Roles:       roles,
	}

	if cred, err = self.storage.CreateCredential(cred); err != nil {
		self.logger.WithError(err).Errorf("failed to create credential in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateCredentialResponse{
		Credential: copy_credential(cred),
	}

	self.logger.WithField("id", id_str).Infof("create credential")

	return res, nil
}
