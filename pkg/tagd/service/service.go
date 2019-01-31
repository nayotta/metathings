package metathings_tagd_service

import (
	pb "github.com/nayotta/metathings/pkg/proto/tagd"
	tagtk "github.com/nayotta/metathings/pkg/toolkit/tag"
)

type MetathingsTagdService struct {
	tagtk tagtk.TagToolkit
}

func NewMetathingsTagdService() (pb.TagdServiceServer, error) {
	return &MetathingsTagdService{}, nil
}
