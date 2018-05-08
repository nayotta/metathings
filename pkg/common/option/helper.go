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

func (o Option) GetBool(k string) bool {
	v := o.Get(k)
	if v == nil {
		return false
	}
	return v.(bool)
}
