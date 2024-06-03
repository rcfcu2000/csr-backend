package targetmanager

import (
	"fmt"
	"time"
	"xtt/global"
	"xtt/utils"
)

type PalletTarget struct {
	global.GVA_BASE
	ShopID        string  `gorm:"column:shop_id;size:32;comment:店铺id" json:"shop_id,omitempty"`              // ShopID 店铺ID（假设与商品有关联，但非本表主键）
	ShopName      string  `gorm:"column:shop_name;size:255;comment:店铺名" json:"shop_name,omitempty"`          //
	StatisticDate string  `gorm:"column:statistic_date;size:32;comment:月份" json:"statistic_date,omitempty"`  //
	Pallet        string  `gorm:"column:pallet;size:32;comment:货盘" json:"pallet"`                            // 货盘
	ProfitTarget  float64 `gorm:"column:profit_target;type:double(12,2);comment:目标毛利润" json:"profit_target"` // 利润目标
	GmvTarget     float64 `gorm:"column:gmv_target;type:double(12,2);comment:gmv目标" json:"gmv_target"`       //   gmv目标
	MonthlyBudget float64 `gorm:"column:monthly_budget;type:double(12,2);comment:月费用" json:"monthly_budget"` // 月费用
}

func (t *PalletTarget) TableName() string {
	return "biz_shop_pallet_targets"
}

func (t *PalletTarget) Save() error {
	if t.ID == "" {
		t.ID = "target" + utils.GenId().String()
	}
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
		t.UpdatedAt = time.Now()
	}
	return global.GVA_DB.Save(t).Error
}

func (t *PalletTarget) GetById(id string) error {
	return global.GVA_DB.First(t, "id = ?", id).Error
}

func (t *PalletTarget) DelById(id string) error {
	fmt.Println("id", 1000)
	return global.GVA_DB.Delete(t, "id = ?", id).Debug().Error
}

func GetPalletTargetList(req ReqPalletTargetSearch) (RespPalletTargetListData, error) {
	resp := &RespPalletTargetListData{Records: []PalletTarget{}}
	resp.GetData(req)
	return *resp, nil
}

// ///////////////////////////////////////////////////
// 产品目标
type ProductTarget struct {
	global.GVA_BASE
	ProductID     string  `gorm:"column:product_id;size:32;comment:店铺id" json:"shop_id,omitempty"`           // ShopID 店铺ID（假设与商品有关联，但非本表主键）
	ProductName   string  `gorm:"column:product_name;size:255;comment:店铺名" json:"shop_name,omitempty"`       //
	StatisticDate string  `gorm:"column:statistic_date;size:32;comment:月份" json:"statistic_date,omitempty"`  //
	ProfitTarget  float64 `gorm:"column:profit_target;type:double(12,2);comment:目标毛利润" json:"profit_target"` // 利润目标
	GmvTarget     float64 `gorm:"column:gmv_target;type:double(12,2);comment:gmv目标" json:"gmv_target"`       //   gmv目标
	MonthlyBudget float64 `gorm:"column:monthly_budget;type:double(12,2);comment:月费用" json:"monthly_budget"` // 月费用
}

func (t *ProductTarget) TableName() string {
	return "biz_product_targets"
}

func (t *ProductTarget) Save() error {
	if t.ID == "" {
		t.ID = "target" + utils.GenId().String()
	}
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
		t.UpdatedAt = time.Now()
	}
	return global.GVA_DB.Save(t).Error
}

func (t *ProductTarget) GetById(id string) error {
	return global.GVA_DB.First(t, "id = ?", id).Error
}

func (t *ProductTarget) DelById(id string) error {
	fmt.Println("id", 1000)
	return global.GVA_DB.Delete(t, "id = ?", id).Debug().Error
}

func GetProductTargetList(req ReqProductTargetSearch) (RespProductTargetListData, error) {
	resp := &RespProductTargetListData{Records: []ProductTarget{}}
	resp.GetData(req)
	return *resp, nil
}
