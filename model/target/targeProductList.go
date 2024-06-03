package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductListData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_ProductList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespPalletTargetListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPalletTargetListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}

func (i *RespProductListData) Prodata() {

}

const SQL_TARGET_ProductList = `
SELECT
product_id as product_id,
product_name as product_name,
category_lv3 as category_lv3,
sum(paid_amount) as gmv,
sum(gmv_target) as target_gmv,
sum(gmv_target) as month_gmv,
sum(paid_amount)/sum(gmv_target) as target_gmv_rate,
( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  time_schedule,
 
sum(spend) AS spend,
sum(spend) /sum(paid_amount)   as "promotion_percentage",
sum(paid_amount) / sum(spend)  as "composite_roi",
sum(paid_amount) / sum(buyer_count)  as "customer_unit_price",
sum(buyer_count)  as "paid_buyers",
sum(jlr) as profit,
sum(profit_target) as  profit_target, -- 经营利润目标
sum(jlr)/sum(profit_target) as  profit_rate  -- 经营利润达成率

FROM
v_target_bpp2_wx_bpt tmain
where
tmain.statistic_date >= :startdate 
AND tmain.statistic_date <=  :enddate  
group by  	  product_id, product_name,category_lv3
LIMIT :offset , :pageSize  
`
