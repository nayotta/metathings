package metathings_deviced_connection

import (
	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
)

func generate_session() int32 {
	return rand_helper.Int31()
}
