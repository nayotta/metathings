package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/pflag"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
)

var (
	webhookUrl       string
	webhookSecret    string
	action           string
	credentialId     string
	credentialSecret string
)

func main() {
	pflag.StringVar(&webhookUrl, "url", "http://127.0.0.1:21884/webhook", "webhook url")
	pflag.StringVar(&webhookSecret, "secret", "", "webhook secret")
	pflag.StringVar(&action, "action", "", "action[created_credential, delete_credential]")
	pflag.StringVar(&credentialId, "credential-id", "", "credential id")
	pflag.StringVar(&credentialSecret, "credential-secret", "", "credential secret")

	pflag.Parse()

	fmt.Printf(`[ARGUMENTS]
     URL : %s
  SECRET : %s
  ACTION : %s
 CRED-ID : %s
CRED-SEC : %s

`, webhookUrl, webhookSecret, action, credentialId, credentialSecret)

	switch action {
	case "create_credential":
		if credentialId == "" {
			fmt.Println("require credential-id")
			os.Exit(1)
		}

		if credentialSecret == "" {
			fmt.Println("require credential-secret")
			os.Exit(1)
		}
	case "delete_credential":
		if credentialId == "" {
			fmt.Println("require credential-id")
			os.Exit(1)
		}

	default:
		fmt.Println("unsupported action:", action)
		os.Exit(1)
	}

	ts := time.Now()
	nonce := rand.Int63()
	webhookId := id_helper.NewId()
	b64secret := base64.StdEncoding.EncodeToString([]byte(webhookSecret))
	hmac := passwd_helper.MustParseHmac(b64secret, webhookId, ts, nonce)

	fmt.Printf(`[HEADERS]
ts(rfc3339nano): %s
ts(unixnano): %v
nonce: %v
webhook-id: %s
hmac: %s
`, ts.Format(time.RFC3339Nano), ts.UnixNano(), nonce, webhookId, hmac)

	cred := map[string]any{
		"action": action,
		"credential": map[string]any{
			"id":     credentialId,
			"secret": credentialSecret,
		},
	}
	credStr, err := json.Marshal(cred)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewReader(credStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("MT-Webhook-Id", webhookId)
	req.Header.Set("MT-Webhook-Timestamp", cast.ToString(ts.UnixNano()))
	req.Header.Set("MT-Webhook-Nonce", cast.ToString(nonce))
	req.Header.Set("MT-Webhook-HMAC", hmac)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("ok")
}
