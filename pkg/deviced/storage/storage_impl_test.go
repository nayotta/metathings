package metathings_deviced_storage

import (
	"context"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/suite"
)

var (
	test_firmware_hub_id                 = "test_firmware_hub_id"
	test_firmware_hub_alias              = "test_firmware_hub_alias"
	test_firmware_hub_description        = "test firmware hub description"
	test_firmware_hub2_id                = "test_firmware_hub2_id"
	test_firmware_hub2_alias             = "test_firmware_hub2_alias"
	test_firmware_hub2_description       = "test_firmware_hub2_description"
	test_device_id                       = "test_device_id"
	test_device2_id                      = "test_device2_id"
	test_firmware_descriptor_id          = "test_firmware_descriptor_id"
	test_firmware_descriptor_name        = "test_firmware_descriptor_name"
	test_firmware_descriptor_descriptor  = "{}"
	test_firmware_descriptor2_id         = "test_firmware_descriptor2_id"
	test_firmware_descriptor2_name       = "test_firmware_descriptor2_name"
	test_firmware_descriptor2_descriptor = "{}"

	test_firmware_hub = &FirmwareHub{
		Id:          &test_firmware_hub_id,
		Alias:       &test_firmware_hub_alias,
		Description: &test_firmware_hub_description,
	}
	test_firmware_hub2 = &FirmwareHub{
		Id:          &test_firmware_hub2_id,
		Alias:       &test_firmware_hub2_alias,
		Description: &test_firmware_hub2_description,
	}
	test_device              = &Device{Id: &test_device_id}
	test_device2             = &Device{Id: &test_device2_id}
	test_firmware_descriptor = &FirmwareDescriptor{
		Id:            &test_firmware_descriptor_id,
		Name:          &test_firmware_descriptor_name,
		FirmwareHubId: &test_firmware_hub_id,
		Descriptor:    &test_firmware_descriptor_descriptor,
	}
	test_firmware_descriptor2 = &FirmwareDescriptor{
		Id:            &test_firmware_descriptor2_id,
		Name:          &test_firmware_descriptor2_name,
		FirmwareHubId: &test_firmware_hub_id,
		Descriptor:    &test_firmware_descriptor2_descriptor,
	}
)

type StorageImplTestSuite struct {
	suite.Suite
	stor *StorageImpl
	ctx  context.Context
}

func (s *StorageImplTestSuite) SetupTest() {
	stor, err := NewStorageImpl(test_helper.GetTestGormDriver(), test_helper.GetTestGormUri(), "logger", log.New())
	s.Require().Nil(err)

	s.stor = stor.(*StorageImpl)
	s.ctx = context.TODO()

	_, err = s.stor.CreateFirmwareHub(s.ctx, test_firmware_hub)
	s.Require().Nil(err)

	err = s.stor.AddDeviceToFirmwareHub(s.ctx, test_firmware_hub_id, test_device_id)
	s.Require().Nil(err)

	err = s.stor.CreateFirmwareDescriptor(s.ctx, test_firmware_descriptor)
	s.Require().Nil(err)

	err = s.stor.SetDeviceFirmwareDescriptor(s.ctx, test_device_id, test_firmware_descriptor_id)
	s.Require().Nil(err)
}

func (s *StorageImplTestSuite) BeforeTest(suiteName, testName string) {
	if fn, ok := map[string]func(){}[testName]; ok {
		fn()
	}
}

func (s *StorageImplTestSuite) TestCreateFirmwareHub() {
	frm_hub, err := s.stor.CreateFirmwareHub(s.ctx, test_firmware_hub2)
	s.Require().Nil(err)

	s.Equal(test_firmware_hub2_id, *frm_hub.Id)
	s.Equal(test_firmware_hub2_alias, *frm_hub.Alias)
	s.Equal(test_firmware_hub2_description, *frm_hub.Description)
	s.Len(frm_hub.Devices, 0)
	s.Len(frm_hub.FirmwareDescriptors, 0)

	frm_hubs, err := s.stor.ListFirmwareHubs(s.ctx, &FirmwareHub{})
	s.Require().Nil(err)
	s.Len(frm_hubs, 2)
}

func (s *StorageImplTestSuite) TestDeleteFirmwareHub() {
	err := s.stor.DeleteFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)
	frm_hubs, err := s.stor.ListFirmwareHubs(s.ctx, &FirmwareHub{})
	s.Require().Nil(err)
	s.Len(frm_hubs, 0)
}

func (s *StorageImplTestSuite) TestPatchFirmwareHub() {
	frm_hub, err := s.stor.PatchFirmwareHub(s.ctx, test_firmware_hub_id, &FirmwareHub{
		Alias:       &test_firmware_hub2_alias,
		Description: &test_firmware_hub2_description,
	})
	s.Require().Nil(err)

	s.Equal(test_firmware_hub_id, *frm_hub.Id)
	s.Equal(test_firmware_hub2_alias, *frm_hub.Alias)
	s.Equal(test_firmware_hub2_description, *frm_hub.Description)
	s.Len(frm_hub.Devices, 1)
	s.Len(frm_hub.FirmwareDescriptors, 1)
}

func (s *StorageImplTestSuite) TestGetFirmwareHub() {
	frm_hub, err := s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Equal(test_firmware_hub_id, *frm_hub.Id)
	s.Equal(test_firmware_hub_alias, *frm_hub.Alias)
	s.Equal(test_firmware_hub_description, *frm_hub.Description)
	s.Len(frm_hub.Devices, 1)
	s.Len(frm_hub.FirmwareDescriptors, 1)
}

func (s *StorageImplTestSuite) TestListFirmwareHub() {
	frm_hubs, err := s.stor.ListFirmwareHubs(s.ctx, &FirmwareHub{Id: &test_firmware_hub_id})
	s.Require().Nil(err)

	s.Require().Len(frm_hubs, 1)
	frm_hub := frm_hubs[0]

	s.Equal(test_firmware_hub_id, *frm_hub.Id)
	s.Equal(test_firmware_hub_alias, *frm_hub.Alias)
	s.Equal(test_firmware_hub_description, *frm_hub.Description)
	s.Len(frm_hub.Devices, 1)
	s.Len(frm_hub.FirmwareDescriptors, 1)
}

func (s *StorageImplTestSuite) TestAddDeviceFirmwareHub() {
	err := s.stor.AddDeviceToFirmwareHub(s.ctx, test_firmware_hub_id, test_device2_id)
	s.Require().Nil(err)

	frm_hub, err := s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Len(frm_hub.Devices, 2)
}

func (s *StorageImplTestSuite) TestRemoveDeviceFirmwareHub() {
	err := s.stor.RemoveDeviceFromFirmwareHub(s.ctx, test_firmware_hub_id, test_device_id)
	s.Require().Nil(err)

	frm_hub, err := s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Len(frm_hub.Devices, 0)
}

func (s *StorageImplTestSuite) TestRemoveAllDevicesInFirmwareHub() {
	s.stor.AddDeviceToFirmwareHub(s.ctx, test_firmware_hub_id, test_device2_id)

	frm_hub, err := s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Len(frm_hub.Devices, 2)

	err = s.stor.RemoveAllDevicesInFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	frm_hub, err = s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Len(frm_hub.Devices, 0)
}

func (s *StorageImplTestSuite) TestCreateFirmwareDescriptor() {
	err := s.stor.CreateFirmwareDescriptor(s.ctx, test_firmware_descriptor2)
	s.Require().Nil(err)

	frm_hub, err := s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Len(frm_hub.FirmwareDescriptors, 2)
}

func (s *StorageImplTestSuite) TestDeleteFirmwareDescriptor() {
	err := s.stor.DeleteFirmwareDescriptor(s.ctx, test_firmware_descriptor_id)
	s.Require().Nil(err)

	frm_hub, err := s.stor.GetFirmwareHub(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)

	s.Len(frm_hub.FirmwareDescriptors, 0)
}

func (s *StorageImplTestSuite) TestListViewDevicesByFirmwareHubId() {
	err := s.stor.AddDeviceToFirmwareHub(s.ctx, test_firmware_hub_id, test_device2_id)
	s.Require().Nil(err)

	devs, err := s.stor.ListViewDevicesByFirmwareHubId(s.ctx, test_firmware_hub_id)
	s.Require().Nil(err)
	s.Len(devs, 2)

	s.Contains(devs, test_device)
	s.Contains(devs, test_device2)
}

func (s *StorageImplTestSuite) TestGetDeviceFirmwareDescriptor() {
	fd, err := s.stor.GetDeviceFirmwareDescriptor(s.ctx, test_device_id)
	s.Require().Nil(err)

	s.Equal(test_firmware_descriptor_id, *fd.Id)
	s.Equal(test_firmware_descriptor_name, *fd.Name)
	s.Equal(test_firmware_descriptor_descriptor, *fd.Descriptor)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(StorageImplTestSuite))
}
