package hmbiz

import (
	"context"

	hmmdl "bds/module/home/model"
)

type CreateHomeStorage interface {
	Create(ctx context.Context, data *hmmdl.HomeCreate) error
}

type CreateHomeBiz struct {
	storage CreateHomeStorage
}

func NewCreateHomeBiz(storage CreateHomeStorage) *CreateHomeBiz {
	return &CreateHomeBiz{storage: storage}
}

func (biz *CreateHomeBiz) CreateHome(ctx context.Context, data *hmmdl.HomeCreate) error {
	if err := biz.storage.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
