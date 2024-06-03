package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespCategoryTargetListData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	if r.Lv3 == "" {
		sqlp.SetSql(SQL_TARGET_CategoryTargetList)
	} else {
		sqlp.SetSql(SQL_TARGET_CategoryTargetList_Detail)
	}
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespCategoryTargetListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespCategoryTargetListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespCategoryTargetListData) Prodata() {

}

const SQL_TARGET_CategoryTargetList = `

SELECT
 
tmain.category_lv3 as category_lv3,
sum(paid_amount) as gmv,	
sum(gmv_target) as target_gmv,
sum(paid_amount)/sum(gmv_target) as target_gmv_rate, 
( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  target_day_rate,
sum(jlr) as profit, 
sum(profit_target) as  profit_rate, -- 经营利润目标
sum(jlr)/sum(profit_target) as  profit_rate  -- 经营利润达成率

FROM
v_target_bpp2_wx_bpt tmain
where
tmain.statistic_date >= :startdate 
AND tmain.statistic_date <=  :enddate  
group by  	  tmain.category_lv3
`

const SQL_TARGET_CategoryTargetList_Detail = `
SELECT
product_id as product_id,
product_name as product_name,
sum(paid_amount) as gmv,	
sum(gmv_target) as target_gmv,
sum(paid_amount)/sum(gmv_target) as target_gmv_rate, 
( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  target_day_rate,
sum(jlr) as profit, 
sum(profit_target) as  profit_rate, -- 经营利润目标
sum(jlr)/sum(profit_target) as  profit_rate  -- 经营利润达成率

FROM
v_target_bpp2_wx_bpt tmain
where
tmain.statistic_date >= :startdate 
AND tmain.statistic_date <=  :enddate  
AND  category_lv3 = :lv3 
group by  	  product_id, product_name
LIMIT :offset , :pageSize  
`
