package shophome

import (
	"xtt/model/common"
)

type ReqShopHomeAllSearch struct {
	// 店铺名称
	ShopName  string `json:"shop_name"`  // 店铺名称
	StartDate string `json:"start_date"` // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate   string `json:"end_date"`   // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	PageNum   int    `json:"pageNum"`    //分页，页号
	PageSize  int    `json:"pageSize"`   //分页，每页数目
}

func (r *ReqShopHomeAllSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":shop_name", ValType: "string", ValString: r.ShopName})
	if r.PageNum <= 0 || r.PageSize <= 0 {
		r.PageNum = 1
		r.PageSize = 9999
	}
	ofset := (r.PageNum - 1) * r.PageSize
	sqlp.SetKeyVal(common.KeyVal{Key: ":offset", ValType: "int", ValInt: int64(ofset)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pageSize", ValType: "int", ValInt: int64(r.PageSize)})
}
