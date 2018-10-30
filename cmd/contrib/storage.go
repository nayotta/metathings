package cmd_contrib

type GetStorageOptioner interface {
	GetStorage() StorageOptioner
}

type StorageOptioner interface {
	GetDriverP() *string
	GetDriver() string
	SetDriver(string)

	GetUriP() *string
	GetUri() string
	SetUri(string)
}

type StorageOption struct {
	Storage struct {
		Driver string
		Uri    string
	}
}

func (self *StorageOption) GetDriverP() *string {
	return &self.Storage.Driver
}

func (self *StorageOption) GetDriver() string {
	return self.Storage.Driver
}

func (self *StorageOption) SetDriver(drv string) {
	self.Storage.Driver = drv
}

func (self *StorageOption) GetUriP() *string {
	return &self.Storage.Uri
}

func (self *StorageOption) GetUri() string {
	return self.Storage.Uri
}

func (self *StorageOption) SetUri(uri string) {
	self.Storage.Uri = uri
}
