package hmbiz

import (
	"context"
	"errors"

	"bds/common"
	hmmdl "bds/module/home/model"
)

type UpdateHomeStorage interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*hmmdl.Home, error)
	Update(ctx context.Context, id int, data *hmmdl.HomeUpdate) error
}

type UpdateHomeBiz struct {
	storage UpdateHomeStorage
}

func NewUpdateHomeBiz(storage UpdateHomeStorage) *UpdateHomeBiz {
	return &UpdateHomeBiz{storage: storage}
}

func (biz *UpdateHomeBiz) UpdateHome(ctx context.Context, id int, data *hmmdl.HomeUpdate) error {
	oldData, err := biz.storage.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(hmmdl.EntityName, err)
	}

	if oldData.Status == 0 {
		return errors.New("Data has been deleted")
	}
	if err := biz.storage.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
