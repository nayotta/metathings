package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/spf13/pflag"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
)

var (
	domain    string
	id        string
	secret    string
	timestamp string
	nonce     int64
)

const (
	DEFAULT_TIMESTAMP = "1970-01-01T00:00:00Z"
)

func main() {
	pflag.StringVar(&domain, "domain", "default", "Credential Domain")
	pflag.StringVar(&id, "id", "", "Credential ID")
	pflag.StringVar(&secret, "secret", "", "Credential Secret")
	pflag.StringVar(&timestamp, "timestamp", DEFAULT_TIMESTAMP, "Timestamp for hmac")
	pflag.Int64Var(&nonce, "nonce", 0, "Nonce for hmac")

	pflag.Parse()

	var ts time.Time

	if timestamp == DEFAULT_TIMESTAMP {
		ts = time.Now()
	} else {
		ts, _ = time.Parse(time.RFC3339Nano, timestamp)
	}

	if nonce == 0 {
		nonce = rand.Int63()
	}

	hmac := passwd_helper.MustParseHmac(secret, id, ts, nonce)

	req := map[string]interface{}{
		"credential": map[string]interface{}{
			"id": id,
			"domain": map[string]interface{}{
				"id": domain,
			},
		},
		"timestamp": ts.Format(time.RFC3339Nano),
		"nonce":     nonce,
		"hmac":      hmac,
	}
	buf, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}
