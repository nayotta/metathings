package metathings_tagd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/tagd"
)

func (ts *MetathingsTagdService) AuthorizeUntag(ctx context.Context, in interface{}) error {
	return ts.authorizer.Authorize(ctx, in.(*pb.UntagRequest).GetId().GetValue(), TAGD_UNTAG)
}

func (ts *MetathingsTagdService) ValidateUntag(ctx context.Context, in interface{}) error {
	return ts.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, tags_getter) {
				req := in.(*pb.UntagRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_tags_size_gt0,
		},
	)
}

func (ts *MetathingsTagdService) Untag(ctx context.Context, req *pb.UntagRequest) (*empty.Empty, error) {
	logger := ts.GetLogger()

	id := req.GetId().GetValue()
	ns := req.GetNamespace().GetValue()
	var tags []string

	for _, tag := range req.GetTags() {
		tags = append(tags, tag.GetValue())
	}

	err := ts.stor.Untag(ns, id, tags)
	if err != nil {
		logger.WithError(err).Errorf("failed to untag")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.WithFields(log.Fields{
		"id":        id,
		"namespace": ns,
		"tags":      tags,
	}).Debugf("untag")

	return &empty.Empty{}, nil
}
