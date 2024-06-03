package crowdsa

import (
	"fmt"
	"xtt/model/common"
)

type ReqCrowdSaAllSearch struct {
	// 店铺名称
	ShopName        string   `json:"shop_name"`        // 店铺名称
	StartDate       string   `json:"start_date"`       // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate         string   `json:"end_date"`         // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	SecondarySource string   `json:"secondary_source"` // 二级来源
	TertiarySource  string   `json:"tertiary_source"`  // 三级来源
	CrowdType       string   `json:"crowd_type"`       // 人群类型
	Ids             []string `json:"ids"`              //输入 product_id
	PageNum         int      `json:"pageNum"`          // 分页，页号
	PageSize        int      `json:"pageSize"`         // 分页，每页数目
}

func (r *ReqCrowdSaAllSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":shop_name", ValType: "string", ValString: r.ShopName})
	sqlp.SetKeyVal(common.KeyVal{Key: ":crowd_type", ValType: "string", ValString: r.CrowdType})
	sqlp.SetKeyVal(common.KeyVal{Key: ":secondary_source", ValType: "string", ValString: r.SecondarySource})
	sqlp.SetKeyVal(common.KeyVal{Key: ":tertiary_source", ValType: "string", ValString: r.TertiarySource})
	sqlp.SetKeyVal(common.KeyVal{Key: ":ids", ValType: "list", ValString: PidToInStringStr(r.Ids)})

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

func PidToInStringStr(i []string) string {
	if len(i) == 0 {
		sql := fmt.Sprintf(" (select DISTINCT product_id  from biz_shop_audience_pruduct_t10 t   ) ") //+ where >=StartDate <= EndDate
		return sql
	}
	str := ""
	for _, item := range i {
		str += fmt.Sprintf("'%s',", item)

	}
	str = str[0 : len(str)-1]
	str = fmt.Sprintf("( %s )", str)
	return str

}
