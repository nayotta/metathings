package webhook_helper

import (
	"bytes"

	"github.com/nayotta/viper"
)

type Event struct {
	*viper.Viper
}

func UnmarshalEvent(buf []byte) (*Event, error) {
	v := viper.New()
	v.SetConfigType("json")
	err := v.ReadConfig(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	return &Event{v}, nil
}
