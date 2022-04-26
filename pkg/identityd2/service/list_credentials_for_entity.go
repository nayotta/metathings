package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateListCredentialsForEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() entity_getter {
				req := in.(*pb.ListCredentialsForEntityRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_entity_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeListCredentialsForEntity(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.ListCredentialsForEntityRequest).GetEntity().GetId().GetValue(), "identityd2:list_credentials_for_entity")
}

func (self *MetathingsIdentitydService) ListCredentialsForEntity(ctx context.Context, req *pb.ListCredentialsForEntityRequest) (*pb.ListCredentialsForEntityResponse, error) {
	var err error

	cred := &storage.Credential{}
	ent_id := req.GetEntity().GetId().GetValue()
	cred.EntityId = &ent_id

	creds, err := self.storage.ListCredentials(ctx, cred)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list credentials for entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListCredentialsForEntityResponse{}
	for _, cred = range creds {
		res.Credentials = append(res.Credentials, copy_credential(cred))
	}

	self.logger.WithField("entity", ent_id).Debugf("list credentials for entity")

	return res, nil
}
