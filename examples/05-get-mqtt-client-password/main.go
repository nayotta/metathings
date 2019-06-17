package main

import (
	"fmt"

	"github.com/spf13/pflag"

	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
)

var (
	credential_id     string
	credential_secret string
)

func main() {
	pflag.StringVar(&credential_id, "id", "", "Device Credential ID")
	pflag.StringVar(&credential_secret, "secret", "", "Device Credential Secret")

	pflag.Parse()

	passwd := mosquitto_service.ParseMosquittoPluginPassword(credential_id, credential_secret)
	fmt.Printf(`Mosquitto Client Username: %v
Mosquitto Client Password: %v

Subscribe:
$ mqttcli sub --host <mosquitto-address> -p <mosquitto-port> -u %v -P %v

Publish:
$ mqttcli pub --host <mosquitto-address> -p <mosquitto-port> -u %v -P %v -s
`, credential_id, passwd, credential_id, passwd, credential_id, passwd)
}
