package metathings_deviced_sdk

func SetBool(key string) func(bool) func(map[string]interface{}) {
	return func(val bool) func(map[string]interface{}) {
		return func(opts map[string]interface{}) {
			opts[key] = val
		}
	}
}

func SetInt(key string) func(int) func(map[string]interface{}) {
	return func(val int) func(map[string]interface{}) {
		return func(opts map[string]interface{}) {
			opts[key] = val
		}
	}
}
