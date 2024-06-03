package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomePromotionData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_Promotion)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_SQL_SHOPHOME_PromotionIndex sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_Promotion sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespShopHomePromotionData) Prodata() {

}

const SQL_SHOPHOME_Promotion = `
WITH t1 AS (
	select sum(gmv) AS all_gmv 
  from wanxiang_product w
  WHERE w.datetimekey >= :startdate  
	AND w.datetimekey <= :enddate  
) 
 select w.promotion_type AS "scene",
 sum(spend) AS "spend",
 sum(spend) / sum(gmv_count) AS "transaction_cost",
 sum(gmv) AS "gmv",
 sum(gmv) / t1.all_gmv AS "scene_percentage",
 sum(gmv) / sum(spend) AS "roi",
 sum(w.clicktraffic) AS "clicks",
 sum(w.clicktraffic) / sum(w.impressions) AS "click_through_rate",
 sum(spend) / sum(w.clicktraffic) AS "cpc",
 sum(shopcart_count) / sum(w.clicktraffic) AS "add_to_cart_rate",
 sum(spend) / sum(shopcart_count) AS "add_to_cart_cost",
 sum(wangwang_count) AS "ali_wang_wang_inquiries"
FROM
	wanxiang_product w
left JOIN t1 ON 1 = 1
WHERE w.datetimekey >= :startdate  
AND w.datetimekey <= :enddate  
GROUP BY promotion_type , t1.all_gmv 
`
