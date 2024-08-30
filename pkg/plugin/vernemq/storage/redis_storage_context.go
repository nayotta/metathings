package metathings_plugin_vernemq_storage

import "context"

func (s *RedisStorage) getContext() context.Context {
	return context.Background()
}
