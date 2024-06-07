package biz

import (
	"context"
	bizModel "xtt/model/biz"
	"xtt/service/system"

	"gorm.io/gorm"
)

const initBiz = system.InitOrderSystem + 1

type initBizModel struct{}

// auto run
func init() {
	system.RegisterInit(initBiz, &initBizModel{})
}

func (i initBizModel) InitializerName() string {
	return bizModel.BizMerchant{}.TableName()
}

func (i *initBizModel) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&bizModel.BizMerchant{},
		&bizModel.BizLinks{},
		&bizModel.BizMerchantLinks{},
	)
}

func (i *initBizModel) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&bizModel.BizMerchant{}) &&
		m.HasTable(&bizModel.BizLinks{}) &&
		m.HasTable(&bizModel.BizMerchantLinks{})
}

func (i *initBizModel) InitializeData(ctx context.Context) (next context.Context, err error) {
	_, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return next, err
}

func (i *initBizModel) DataInserted(ctx context.Context) bool {
	_, ok := ctx.Value("db").(*gorm.DB)
	return ok
}
