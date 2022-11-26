package hmstrg

import (
	"context"

	"bds/common"
	hmmdl "bds/module/home/model"
)

func (s *SqlStore) Update(ctx context.Context, id int, data *hmmdl.HomeUpdate) error {
	if err := s.db.
		Table(hmmdl.Home{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
