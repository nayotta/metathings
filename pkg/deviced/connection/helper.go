package metathings_deviced_connection

import (
	"fmt"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
)

func generate_session() int32 {
	return rand_helper.Int31()
}

func parse_bridge_id(device string, session int32) string {
	return id_helper.NewNamedId(fmt.Sprintf("device.%v.session.%v", device, session))
}
