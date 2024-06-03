package traffic

import (
	"xtt/model/common"
)

type ReqTrafficAllSearch struct {
	// 店铺名称
	ShopName       string   `json:"shop_name"`       // 店铺名称
	StartDate      string   `json:"start_date"`      // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate        string   `json:"end_date"`        // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	TrafficBelong  string   `json:"traffic_belong"`  // 流量归属原则
	ProductManager []string `json:"product_manager"` // 商品负责人姓名或ID，对应于JSON中的"商品负责人"
	Channel        []string `json:"channel"`         // 渠道
	ProductId      string   `json:"productid"`       // 商品ID
	PageNum        int      `json:"pageNum"`         // 分页，页号
	PageSize       int      `json:"pageSize"`        // 分页，每页数目
}

func (r *ReqTrafficAllSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":belong", ValType: "string", ValString: r.TrafficBelong})
	sqlp.SetKeyVal(common.KeyVal{Key: ":shop_name", ValType: "string", ValString: r.ShopName})
	sqlp.SetKeyVal(common.KeyVal{Key: ":productid", ValType: "string", ValString: r.ProductId})

	sqlp.SetKeyVal(common.KeyVal{Key: ":channel", ValType: "list", ValString: common.ListToInStringStr(r.Channel, "channel")})
	if len(r.ProductManager) == 0 {
		sqlp.SetKeyVal(common.KeyVal{Key: ":resperson", ValType: "stringsrc", ValString: GetAllProductManager()})
	} else {
		sqlp.SetKeyVal(common.KeyVal{Key: ":resperson", ValType: "list", ValString: common.ListToInStringStr(r.ProductManager, "resperson")})
	}

	if r.PageNum <= 0 || r.PageSize <= 0 {
		r.PageNum = 1
		r.PageSize = 9999
	}
	ofset := (r.PageNum - 1) * r.PageSize
	sqlp.SetKeyVal(common.KeyVal{Key: ":offset", ValType: "int", ValInt: int64(ofset)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pageSize", ValType: "int", ValInt: int64(r.PageSize)})
}

func GetAllProductManager() string {
	return "(select DISTINCT responsible  from biz_product  ) "
}
