package traffic

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductInfoListData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_ProductList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_TRAFFIC_ProductList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_TRAFFIC_ProductList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespProductInfoListData) Prodata() {

}

const SQL_TRAFFIC_ProductList = `

select  
	product_id, product_name as product_alias, product_name, responsible, category_lv2, pallet
-- 	gmv,profit,cur_pallet, pre_pallet,pallet_change, 
--   old_percentage
-- 	pay_gmv_percentage, spend_percentage, overall_score, add_car_efficiency, 
-- 	repurchase_efficiency, loss_efficiency,conversion_efficiency, pay_conversion_rate,
-- 	uv, bounce_rate, avg_stay_duration, belong, depth_visit
 
from biz_pallet_product  bpp
where statistic_date >= "2024-01-01"  and statistic_date<="2024-01-11" 
order by statistic_date
LIMIT :offset , :pageSize  

`
