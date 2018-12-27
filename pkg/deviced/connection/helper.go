package metathings_deviced_connection

import (
	"fmt"

	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
)

func generate_session() int32 {
	return rand_helper.Int31()
}

func bridge_id_to_symbol(id string) string {
	return fmt.Sprintf("metathings.deviced.connection.bridge.%v", id)
}
