package product

import (
	"xtt/model/common"
)

type ReqProductListSearch struct {
	// key 模糊查询， 返回 product_id  和  product_name
	Key string `json:"key"`
}

type ReqProductAllSearch struct {
	// 商品编号
	ProductId string `json:"product_id"`
	StartDate string `json:"start_date"` // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate   string `json:"end_date"`   // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	PageNum   int    `json:"pageNum"`    //分页，页号
	PageSize  int    `json:"pageSize"`   //分页，每页数目
}

type ReqProductThendSearch struct {
	// 商品编号
	ProductId string `json:"product_id"`
	StartDate string `json:"start_date"` // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate   string `json:"end_date"`   // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	PageNum   int    `json:"pageNum"`    //分页，页号
	PageSize  int    `json:"pageSize"`   //分页，每页数目
}

func (r *ReqProductAllSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {

	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":productid", ValType: "string", ValString: r.ProductId})

	if r.PageNum <= 0 || r.PageSize <= 0 {
		r.PageNum = 1
		r.PageSize = 9999
	}
	ofset := (r.PageNum - 1) * r.PageSize
	sqlp.SetKeyVal(common.KeyVal{Key: ":offset", ValType: "int", ValInt: int64(ofset)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pageSize", ValType: "int", ValInt: int64(r.PageSize)})
}
