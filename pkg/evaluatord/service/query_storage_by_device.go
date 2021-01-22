package metathings_evaluatord_service

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"
	structpb "github.com/golang/protobuf/ptypes/struct"
	log "github.com/sirupsen/logrus"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	"github.com/nayotta/metathings/proto/common/option/pagination"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
	pb "github.com/nayotta/metathings/proto/evaluatord"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
)

func (srv *MetathingsEvaluatordService) ValidateQueryStorageByDevice(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.QueryStorageByDeviceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizeQueryStorageByDevice(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.QueryStorageByDeviceRequest).GetDevice().GetId().GetValue(), "evaluatord:query_storage_by_device")
}

func (srv *MetathingsEvaluatordService) QueryStorageByDevice(ctx context.Context, req *pb.QueryStorageByDeviceRequest) (*pb.QueryStorageByDeviceResponse, error) {
	var err error

	dev := req.GetDevice()
	dev_id := dev.GetId().GetValue()

	src := req.GetSource()
	src_id := src.GetId().GetValue()
	src_typ := src.GetType().GetValue()

	meas := req.GetMeasurement().GetValue()

	logger := srv.get_logger().WithFields(log.Fields{
		"device":      dev_id,
		"source_id":   src_id,
		"source_type": src_typ,
	})

	tags := map[string]string{
		"$device_id":   dev_id,
		"$source_id":   src_id,
		"$source_type": src_typ,
	}

	var opts []dssdk.QueryOption

	query_string := req.GetQueryString().GetValue()
	opts = append(opts,
		dssdk.QueryString(query_string),
		dssdk.PageSize(srv.opt.Methods.QueryStorageByDevice.DefaultPageSize),
	)

	var range_from_ts time.Time
	if range_from := req.GetRangeFrom(); range_from != nil {
		range_from_ts, err = ptypes.Timestamp(range_from)
		if err != nil {
			logger.WithError(err).Errorf("range_from is invalid value")
			return nil, srv.ParseError(err)
		}
	} else {
		range_from_ts = time.Now().Add(srv.opt.Methods.QueryStorageByDevice.DefaultRangeFromDuration)
	}
	opts = append(opts, dssdk.RangeFrom(range_from_ts))
	logger = logger.WithField("range_from", range_from_ts)

	if range_to := req.GetRangeTo(); range_to != nil {
		range_to_ts, err := ptypes.Timestamp(range_to)
		if err != nil {
			logger.WithError(err).Errorf("range_to is invalid value")
			return nil, srv.ParseError(err)
		}
		opts = append(opts, dssdk.RangeTo(range_to_ts))
		logger = logger.WithField("range_to", range_to_ts)
	}

	if pagination := req.GetPagination(); pagination != nil {
		if page_size := pagination.GetPageSize(); page_size != nil {
			page_size_i32 := page_size.GetValue()
			opts = append(opts, dssdk.PageSize(page_size_i32))
			logger = logger.WithField("page_size", page_size_i32)
		}

		if page_token := pagination.GetPageToken(); page_token != nil {
			page_token_str := page_token.GetValue()
			opts = append(opts, dssdk.PageToken(page_token_str))
			logger = logger.WithField("page_token", page_token_str)
		}
	}

	ret, err := srv.data_storage.Query(ctx, meas, tags, opts...)
	if err != nil {
		logger.WithError(err).Errorf("failed to query data in storage")
		return nil, srv.ParseError(err)
	}

	var frms []*deviced_pb.Frame
	for _, rec := range ret.Records() {
		pbts, err := ptypes.TimestampProto(rec.Time())
		if err != nil {
			logger.WithError(err).Warningf("failed to parse record.time to protobuf timestamp")
			continue
		}

		dat := rec.Data()
		dat_buf, err := json.Marshal(dat)
		if err != nil {
			logger.WithError(err).Warningf("failed to parse record.data to json string")
			continue
		}
		var dat_st structpb.Struct
		err = grpc_helper.JSONPBUnmarshaler.Unmarshal(bytes.NewReader(dat_buf), &dat_st)
		if err != nil {
			logger.WithError(err).Warningf("failed to parse record.data json string to protobuf struct")
			continue
		}

		frms = append(frms, &deviced_pb.Frame{
			Ts:   pbts,
			Data: &dat_st,
		})
	}

	res := &pb.QueryStorageByDeviceResponse{
		Frames: frms,
		Pagination: &pagination.PaginationOption{
			NextPageToken: ret.NextPageToken(),
		},
	}

	logger.Debugf("query storage by device")

	return res, nil
}
