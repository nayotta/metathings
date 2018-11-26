package metathingsmqttdconnection

//	log "github.com/sirupsen/logrus"
//	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
//	pb "github.com/nayotta/metathings/pkg/proto/mqttd"

// Connection Connection
type Connection interface {
	Err(err ...error) error
	Wait() chan bool
}

type connection struct {
	err error
	c   chan bool
}

// MqttBridge MqttBridge interface
type MqttBridge interface {
	Pub(msg []byte, path string) error
	InitMqttBridge() error
	CloseBridge()
}

// GpioType GpioType
var GpioType = "gpio"

// I2cType I2cType
var I2cType = "i2c"
