package metathings_tagd_service

import (
	log "github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	pb "github.com/nayotta/metathings/pkg/proto/tagd"
	tagtk "github.com/nayotta/metathings/pkg/toolkit/tag"
)

type MetathingsTagdService struct {
	*log_helper.GetLoggerer

	logger log.FieldLogger
	tagtk  tagtk.TagToolkit
}

func NewMetathingsTagdService(logger log.FieldLogger) (pb.TagdServiceServer, error) {
	ts := &MetathingsTagdService{
		GetLoggerer: log_helper.NewGetLoggerer(logger),
		logger:      logger,
	}
	return ts, nil
}
