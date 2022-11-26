package hmstrg

import (
	"context"

	"bds/common"
	hmmdl "bds/module/home/model"

	"gorm.io/gorm"
)

func (s *SqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*hmmdl.Home, error) {
	var data hmmdl.Home
	err := s.db.Where(condition).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
