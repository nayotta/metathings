package metathings_identityd2_service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateAuthorizeToken(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (object_getter, action_getter) {
				req := in.(*pb.AuthorizeTokenRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_id,
			ensure_get_action_name,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeToken(ctx context.Context, req *pb.AuthorizeTokenRequest) (*empty.Empty, error) {
	var err error

	tkn := ctx.Value("token").(*pb.Token)
	sub := tkn.GetEntity().GetId()
	obj := req.GetObject().GetId().GetValue()
	act := req.GetAction().GetName().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"subject": sub,
		"object":  obj,
		"action":  act,
	})

	if err = self.authorize(ctx, obj, act); err != nil {
		if err == policy.ErrPermissionDenied {
			logger.Warningf("permission denied")
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		} else {
			logger.WithError(err).Errorf("failed to authorize in backend")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	logger.Infof("permission authorized")

	return &empty.Empty{}, nil
}

func (self *MetathingsIdentitydService) authorize(ctx context.Context, object, action string) error {
	var err error
	var grps []*storage.Group

	tkn := ctx.Value("token").(*pb.Token)
	dom := &storage.Domain{Id: &tkn.Domain.Id}
	for _, g := range tkn.GetGroups() {
		grps = append(grps, &storage.Group{
			Id:       &g.Id,
			DomainId: &tkn.Domain.Id,
			Domain:   dom,
		})
	}

	sub := &storage.Entity{
		Id:     &tkn.Entity.Id,
		Groups: grps,
	}
	obj := &storage.Entity{
		Id: &object,
	}
	act := &storage.Action{
		Name: &action,
	}

	start := time.Now()
	err = self.backend.Enforce(ctx, sub, obj, act)
	self.logger.WithFields(log.Fields{
		"elapse":  time.Since(start),
		"subject": tkn.Entity.Id,
		"object":  object,
		"action":  action,
	}).Debugf("enforce")

	if err != nil {
		return err
	}
	return nil
}
