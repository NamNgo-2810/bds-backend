package appctx

import (
	upldprv "bds/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() upldprv.UploadProvider
	SecretKey() string
	// GetPubSub() pubsub.PubSub
}

type AppCtx struct {
	db             *gorm.DB
	uploadProvider upldprv.UploadProvider
	secretKey      string
	// ps             pubsub.PubSub
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider upldprv.UploadProvider,
	secretKey string,
	// ps pubsub.PubSub,
) *AppCtx {
	return &AppCtx{
		db:             db,
		uploadProvider: uploadProvider,
		secretKey:      secretKey,
		// ps: ps,
	}
}

func (ctx *AppCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *AppCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *AppCtx) UploadProvider() upldprv.UploadProvider {
	return ctx.uploadProvider
}
