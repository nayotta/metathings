package helper

import (
	"strings"

	"github.com/spf13/viper"
)

type WithPrefix func(string) string
type GetArgument func(string) string

type argumentHelper struct {
	prefix string
}

func (h *argumentHelper) Get(name string) interface{} {
	return viper.Get(h.PrefixWith(name))
}

func (h *argumentHelper) GetString(name string) string {
	v := h.Get(name)
	if v == nil {
		return ""
	}
	return v.(string)
}

func (h *argumentHelper) PrefixWith(name string) string {
	return h.prefix + strings.Replace(strings.ToUpper(name), "-", "_", -1)
}

func NewArgumentHelper(prefix string) *argumentHelper {
	return &argumentHelper{prefix}
}
