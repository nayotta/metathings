package metathings_plugin_vernemq_storage

type PublishAcl struct {
	Pattern string `json:"pattern"`
	MaxQos  int    `json:"max_qos,omitempty"`
}

type SubscribeAcl struct {
	Pattern        string `json:"pattern"`
	MaxQos         int    `json:"max_qos,omitempty"`
	MaxPayloadSize int    `json:"max_payload_size,omitempty"`
	AllowedRetain  bool   `json:"allowed_retain,omitempty"`
}

type User struct {
	MountPoint    string         `json:"mountpoint"`
	ClientID      string         `json:"client_id"`
	Username      string         `json:"username"`
	Passhash      string         `json:"passhash"`
	PublishAcls   []PublishAcl   `json:"publish_acl"`
	SubscribeAcls []SubscribeAcl `json:"subscribe_acl"`
}

type Storage interface {
	CreateOrUpdateUser(user *User) error
	RemoveUser(clientid, username string) error
}
