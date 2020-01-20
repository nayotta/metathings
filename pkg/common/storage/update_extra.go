package storage_helper

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	extra_helper "github.com/nayotta/metathings/contrib/extra"
)

func UpdateExtra(db *gorm.DB, selector interface{}, config map[string]string) error {
	type Result struct {
		Extra string `gorm:"column:extra"`
	}

	var result Result
	if err := db.First(selector).Scan(&result).Error; err != nil {
		return err
	}

	src := map[string]string{}
	if result.Extra != "" {
		if err := json.Unmarshal([]byte(result.Extra), &src); err != nil {
			return err
		}
	}

	src_extra := extra_helper.FromRaw(src)
	dst_extra, err := src_extra.Apply(config)
	if err != nil {
		return err
	}

	dst := dst_extra.Raw()
	dst_buf, err := json.Marshal(dst)
	if err != nil {
		return err
	}

	if err := db.Model(selector).Update("extra", string(dst_buf)).Error; err != nil {
		return err
	}

	return nil
}
