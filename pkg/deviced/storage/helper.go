package metathings_deviced_storage

import "github.com/jinzhu/gorm"

func parse_error(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return RecordNotFound
	default:
		return err
	}
}
