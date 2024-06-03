package targetmanager

import "xtt/model/common"

type ReqId struct {
	Id string `json:"id"` //
}
type ReqPalletTargetSearch struct {
	ShopName  string   `json:"shop_name"`  // 店铺名称
	StartDate string   `json:"start_date"` // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate   string   `json:"end_date"`   // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	Pallet    []string `json:"pallet"`
	PageNum   int      `json:"pageNum"`  //分页，页号
	PageSize  int      `json:"pageSize"` //分页，每页数目
}

type ReqProductTargetSearch struct {
	ShopName    string `json:"shop_name"`    // 店铺名称
	StartDate   string `json:"start_date"`   // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate     string `json:"end_date"`     // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	ProductId   string `json:"product_id"`   // 产品id
	ProductName string `json:"product_name"` // 产品名称
	PageNum     int    `json:"pageNum"`      // 分页，页号
	PageSize    int    `json:"pageSize"`     // 分页，每页数目
}

func (r *ReqPalletTargetSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":shop_name", ValType: "string", ValString: r.ShopName})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pallet", ValType: "list", ValString: common.ListToInStringStr(r.Pallet, "pallet")}) // "('A','B')"

	if r.PageNum <= 0 || r.PageSize <= 0 {
		r.PageNum = 1
		r.PageSize = 9999
	}
	ofset := (r.PageNum - 1) * r.PageSize
	sqlp.SetKeyVal(common.KeyVal{Key: ":offset", ValType: "int", ValInt: int64(ofset)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pageSize", ValType: "int", ValInt: int64(r.PageSize)})
}

func (r *ReqProductTargetSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
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
