package shop

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

/* 缺
   "refund_rate": 0,
      "refund_successful_amount": 0,
      "violation_count": 0
*/

func (i *ShopServiceAnalysis) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_ShopService)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("ShopServiceAnalysis sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("ShopServiceAnalysis sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *ShopServiceAnalysis) Prodata() {

}

const SQL_SHOP_ShopService = `
select first_refund_rate as  "first_product_return_rate",
above_peer_first_refund_rate as  "first_product_return_rate_above_peers",
industry_avg_first_refund_rate as  "first_product_return_rate_vs_industry_average",
industry_excellent_first_refund_rate as  "first_product_return_rate_vs_industry_excellent",
product_negative_review_rate as  "product_negative_review_rate",
above_peer_product_negative_review_rate as  "product_negative_review_rate_above_peers",
industry_avg_product_negative_review_rate as  "product_negative_review_rate_vs_industry_average",
industry_excellent_product_negative_review_rate as  "product_negative_review_rate_vs_industry_excellent",
refund_rate/100.00 as "refund_rate",
refund_amount as  "refund_successful_amount",
violations_count as "violation_count"
 
from biz_shop_service t
left join biz_shop_dayinfo bsd on  t.statistic_date = bsd.statistic_date
where (t.shop_name=:shop_name ) and t.statistic_date=:enddate ;
 
`

// 退款率	指标	退款洞察	取结束日期的值
// 退款成功金额	指标	退款洞察	取结束日期的值
// 违规数量	指标	店铺违规	取结束日期的值

/*

FirstProductReturnRate float64 `json:"first_product_return_rate"`

// 首次品退率是否超过同行， 表示该产品首次退货率是否高于行业平均水平
FirstProductReturnRateAbovePeers float64 `json:"first_product_return_rate_above_peers"`

// 首次品退率与行业平均的对比数值
FirstProductReturnRateVsIndustryAverage float64 `json:"first_product_return_rate_vs_industry_average"`

// 首次品退率与行业优秀标准的对比数值
FirstProductReturnRateVsIndustryExcellent float64 `json:"first_product_return_rate_vs_industry_excellent"`

// 商品差评率，即收到的负向评价的比例
ProductNegativeReviewRate float64 `json:"product_negative_review_rate"`

// 商品差评率是否超过同行， 表示商品差评率是否高于行业平均水平
ProductNegativeReviewRateAbovePeers float64 `json:"product_negative_review_rate_above_peers"`

// 商品差评率与行业平均的对比数值
ProductNegativeReviewRateVsIndustryAverage float64 `json:"product_negative_review_rate_vs_industry_average"`

// 商品差评率与行业优秀标准的对比数值
ProductNegativeReviewRateVsIndustryExcellent float64 `json:"product_negative_review_rate_vs_industry_excellent"`

*/

////////////////////////////////////////////////////////////////////
//

func (i *ShopServiceAnalysisTrendList) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_ShopServiceTrendList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []ShopServiceAnalysisTrendNode{}

	fmt.Println("ShopServiceAnalysisTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("ShopServiceAnalysisTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *ShopServiceAnalysisTrendList) Prodata() {

}

const SQL_SHOP_ShopServiceTrendList = `
select  DATE_FORMAT(t.statistic_date, '%Y-%m-%d')  as date,  
first_refund_rate as  "first_product_return_rate",
above_peer_first_refund_rate as  "first_product_return_rate_above_peers",
industry_avg_first_refund_rate as  "first_product_return_rate_vs_industry_average",
industry_excellent_first_refund_rate as  "first_product_return_rate_vs_industry_excellent",
product_negative_review_rate as  "product_negative_review_rate",
above_peer_product_negative_review_rate as  "product_negative_review_rate_above_peers",
industry_avg_product_negative_review_rate as  "product_negative_review_rate_vs_industry_average",
industry_excellent_product_negative_review_rate as  "product_negative_review_rate_vs_industry_excellent",
refund_rate as "refund_rate",
refund_amount as  "refund_successful_amount",
violations_count as "violation_count"

from biz_shop_service t
left join biz_shop_dayinfo bsd on  t.statistic_date = bsd.statistic_date
where t.statistic_date>=:startdate and t.statistic_date<=:enddate and t.shop_name=:shop_name ;

`

// // 退款率，即退款请求成功的比率
// RefundRate float64 `json:"refund_rate"`

// // 退款成功金额，统计周期内成功退款的总额
// RefundSuccessfulAmount float64 `json:"refund_successful_amount"`

// // 违规数量，记录商家在统计周期内的违规操作次数
// ViolationCount int `json:"violation_count"`
