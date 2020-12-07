package metathings_identityd2_service

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateCreateCredential(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, credential_getter) {
				req := in.(*pb.CreateCredentialRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			func(x credential_getter) error {
				cred := x.GetCredential()

				if cred.GetDomain() == nil || cred.GetDomain().GetId() == nil {
					return errors.New("credential.domain.id is empty")
				}

				if cred.GetEntity() == nil || cred.GetEntity().GetId() == nil {
					return errors.New("credential.entity.id is empty")
				}

				if cred.GetName() == nil {
					return errors.New("credential.name is empty")
				}

				for _, r := range cred.GetRoles() {
					if r.GetId() == nil {
						return errors.New("credential.roles.id is empty")
					}
				}

				return nil
			},
		},
	)
}

func (self *MetathingsIdentitydService) CreateCredential(ctx context.Context, req *pb.CreateCredentialRequest) (*pb.CreateCredentialResponse, error) {
	var err error

	cred := req.GetCredential()

	id_str := id_helper.NewId()
	if cred.GetId() != nil {
		id_str = cred.GetId().GetValue()
	}

	dom_id_str := cred.GetDomain().GetId().GetValue()

	roles := []*storage.Role{}
	for _, r := range cred.GetRoles() {
		roles = append(roles, &storage.Role{
			Id: &r.Id.Value,
		})
	}

	ent_id_str := cred.GetEntity().GetId().GetValue()
	name_str := cred.GetName().GetValue()
	alias_str := name_str
	if cred.GetAlias() != nil {
		alias_str = cred.GetAlias().GetValue()
	}

	var siz int32 = 32
	if req.GetSecretSize() != nil {
		siz = req.GetSecretSize().GetValue()
	}

	srt_str := generate_secret(siz)
	if cred.GetSecret() != nil {
		srt_str = cred.GetSecret().GetValue()
	}

	desc_str := ""
	if cred.GetDescription() != nil {
		desc_str = cred.GetDescription().GetValue()
	}

	now := time.Now()
	var expires time.Time
	if cred.GetExpiresAt() != nil {
		expires = pb_helper.ToTime(*cred.GetExpiresAt())
	} else {
		expires = now.Add(self.opt.CredentialExpire)
	}

	cred_s := &storage.Credential{
		Id:          &id_str,
		DomainId:    &dom_id_str,
		EntityId:    &ent_id_str,
		Name:        &name_str,
		Alias:       &alias_str,
		Secret:      &srt_str,
		Description: &desc_str,
		ExpiresAt:   &expires,
		Roles:       roles,
	}

	if cred_s, err = self.storage.CreateCredential(ctx, cred_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create credential in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	cred_s.Secret = &srt_str

	cred_r := copy_credential_with_secret(cred_s)
	if err = self.webhook.Trigger(map[string]interface{}{
		"action":     "create_credential",
		"credential": cred_r,
	}); err != nil {
		self.logger.WithError(err).Errorf("failed to trigger create credential webhook")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateCredentialResponse{
		Credential: cred_r,
	}

	self.logger.WithField("id", id_str).Infof("create credential")

	return res, nil
}
