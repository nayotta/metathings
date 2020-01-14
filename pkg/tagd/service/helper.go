package metathings_tagd_service

import (
	"errors"

	"github.com/golang/protobuf/ptypes/wrappers"
)

type tags_getter interface {
	GetTags() []*wrappers.StringValue
}

func ensure_tags_size_gt0(x tags_getter) error {
	if len(x.GetTags()) <= 0 {
		return errors.New("tags size should great than 0")
	}
	return nil
}
