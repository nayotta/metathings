package metathings_device_cloud_profile_storage_redis

import (
	"encoding/json"
	"fmt"

	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
)

func (ps *RedisProfileStorage) profileToRedisKey(p intf.Profile) (string, error) {
	return ps.profileNameToRedisKey(p.Name)
}

func (ps *RedisProfileStorage) profileNameToRedisKey(name string) (string, error) {
	return fmt.Sprintf("ps:profiles:%s", name), nil
}

func (ps *RedisProfileStorage) profileToRedisVal(p intf.Profile) (string, error) {
	buf, err := json.Marshal(p)
	return string(buf), err
}

func (ps *RedisProfileStorage) redisValToProfile(s string) (intf.Profile, error) {
	var p intf.Profile
	err := json.Unmarshal([]byte(s), &p)
	return p, err
}

func (ps *RedisProfileStorage) deviceIdToBindingProfileRedisKey(name string) (string, error) {
	return fmt.Sprintf("ps:devices:%s:profile", name), nil
}
