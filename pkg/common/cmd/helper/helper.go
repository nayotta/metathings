package helper

import (
	"strings"

	"github.com/spf13/viper"
)

type argumentHelper struct {
	prefix string
}

func (h *argumentHelper) Get(name string) interface{} {
	return viper.Get(h.prefix + strings.Replace(strings.ToUpper(name), "-", "_", -1))
}

func (h *argumentHelper) GetString(name string) string {
	v := h.Get(name)
	if v == nil {
		return ""
	}
	return v.(string)
}

func NewArgumentHelper(prefix string) *argumentHelper {
	return &argumentHelper{prefix}
}
