package metathings_tagd_service

import (
	log "github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	pb "github.com/nayotta/metathings/pkg/proto/tagd"
	storage "github.com/nayotta/metathings/pkg/tagd/storage"
)

type MetathingsTagdService struct {
	*log_helper.GetLoggerer

	logger log.FieldLogger
	stor   storage.Storage
}

func NewMetathingsTagdService(logger log.FieldLogger, stor storage.Storage) (pb.TagdServiceServer, error) {
	ts := &MetathingsTagdService{
		GetLoggerer: log_helper.NewGetLoggerer(logger),
		logger:      logger,
		stor:        stor,
	}
	return ts, nil
}
