package client_helper

import (
	"github.com/minio/minio-go/v7"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

func ToMinioClient(v **minio.Client) func(string, any) error {
	return func(key string, val any) error {
		var ok bool
		if *v, ok = val.(*minio.Client); !ok {
			return opt_helper.InvalidArgument(key)
		}
		return nil
	}
}
