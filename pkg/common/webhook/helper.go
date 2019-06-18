package webhook_helper

import (
	"net/http"
	"strconv"
	"time"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
)

func deepcopy_webhook(wh *Webhook) *Webhook {
	return &Webhook{
		Id:          wh.Id,
		Url:         wh.Url,
		ContentType: wh.ContentType,
		Secret:      wh.Secret,
	}
}

func ValidateHmac(srt string, req *http.Request) bool {
	id := req.Header.Get("MT-Webhook-Id")
	ts_s := req.Header.Get("MT-Webhook-Timestamp")
	nonce_s := req.Header.Get("MT-Webhook-Nonce")
	hmac := req.Header.Get("MT-Webhook-HMAC")

	nsec, err := strconv.ParseInt(ts_s, 10, 64)
	if err != nil {
		return false
	}
	ts := time.Unix(0, nsec)
	nonce, err := strconv.ParseInt(nonce_s, 10, 64)
	if err != nil {
		return false
	}

	return passwd_helper.ValidateHmac(hmac, srt, id, ts, nonce)
}
