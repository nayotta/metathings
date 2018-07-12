package option_helper

type Option map[string]interface{}

func (o Option) Set(k string, v interface{}) {
	o[k] = v
}

func (o Option) Get(k string) interface{} {
	v, ok := o[k]
	if !ok {
		return nil
	}
	return v
}

func (o Option) Keys() []string {
	ks := make([]string, 0, len(o))
	for k, _ := range o {
		ks = append(ks, k)
	}
	return ks
}

func (o Option) Contains(k string) bool {
	_, ok := o[k]
	return ok
}

func (o Option) GetString(k string) string {
	v := o.Get(k)
	if v == nil {
		return ""
	}
	return v.(string)
}

func (o Option) GetStrings(k string) []string {
	v := o.Get(k)
	if v == nil {
		return nil
	}
	return v.([]string)
}

func (o Option) GetInt(k string) int {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(int)
}

func (o Option) GetUInt32(k string) uint32 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(uint32)
}

func (o Option) GetUInt64(k string) uint64 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(uint64)
}

func (o Option) GetInt32(k string) int32 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(int32)
}

func (o Option) GetInt64(k string) int64 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(int64)
}

func (o Option) GetFloat32(k string) float32 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(float32)
}

func (o Option) GetFloat64(k string) float64 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(float64)
}

func (o Option) GetBool(k string) bool {
	v := o.Get(k)
	if v == nil {
		return false
	}
	return v.(bool)
}

func Copy(opt Option) Option {
	o := Option{}
	for k := range opt {
		o[k] = opt[k]
	}
	return o
}
