package mqtt_helper

import passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"

func ParseMqttPassword(id, secret string) string {
	return passwd_helper.MustParseHmac(secret, id, passwd_helper.DEFAULT_HMAC_TIMESTAMP, passwd_helper.DEFAULT_HMAC_NONCE)
}
