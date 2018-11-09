package metathings_deviced_connection

type Storage interface {
	AddBridgeToDevice(dev_id, br_id string) error
	RemoveBridgeFromDevice(dev_id, br_id string) error
	ListBridgesFromDevice(dev_id string) ([]string, error)
}

type storageImpl struct {
}

func (self *storageImpl) AddBridgeToDevice(dev_id, br_id string) error {
	panic("unimplemented")
}

func (self *storageImpl) RemoveBridgeFromDevice(dev_id, br_id string) error {
	panic("unimplemented")
}

func (self *storageImpl) ListBridgesFromDevice(dev_id string) ([]string, error) {
	panic("unimplemented")
}

func NewStorage() (Storage, error) {
	return &storageImpl{}, nil
}
