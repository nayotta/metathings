package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListCredentials(ctx context.Context, req *pb.ListCredentialsRequest) (*pb.ListCredentialsResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := &storage.Credential{}

	if id := req.GetId(); id != nil {
		cred.Id = id.Value
	}
	if domain := req.GetDomain(); domain != nil {
		if domainID := domain.GetId(); domainID != nil {
			cred.DomainId = domainID.Value
		}
	}
	if entity := req.GetEntity(); entity != nil {
		if entityID := entity.GetId(); entityID != nil {
			cred.entityId = entityID.Value
		}
	}
	if name := req.GetName(); name != nil {
		cred.Name = name.Value
	}
	if alias := req.GetAlias(); alias != nil {
		cred.Alias = alias.Value
	}

	if creds, err := self.storage.ListCredentials(cred); err != nil {
		self.logger.WithError(err).Errorf("failed to list credentials in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListCredentialsResponse{}

	for cred = range creds {
		res.creds = append(res.creds, copy_credential(cred)),
	}

	return res, nil
}
