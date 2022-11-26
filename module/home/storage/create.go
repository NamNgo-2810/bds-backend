package hmstrg

import (
	"context"

	hmmdl "bds/module/home/model"
)

func (s *SqlStore) Create(ctx context.Context, data *hmmdl.HomeCreate) error {
	if err := s.db.Create(&data); err != nil {
		return err.Error
	}

	return nil
}
