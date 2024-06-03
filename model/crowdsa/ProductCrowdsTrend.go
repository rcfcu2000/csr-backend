package crowdsa

import (
	"errors"
	"fmt"
	"xtt/global"
	base "xtt/model/base"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductCrowdsTrendList) GetData(r ReqCrowdSaAllSearch) error {
	db := base.GetDBByShopName(r.ShopName)
	if db == nil {
		global.GVA_LOG.Error("get db by shop error", zap.String("ShopName", r.ShopName))
		return errors.New("get db by shop error")
	}
	sqlp := &common.SQLProccesor{}
	if len(r.StartDate) > len("2024-01") {
		r.StartDate = r.StartDate[0:7]
		r.EndDate = r.EndDate[0:7]
	}
	sqlp.SetSql(SQL_CROWD_ProductCrowdsTrendList)
	if len(r.Ids) == 0 {
		sqlp.SetSql(SQL_CROWD_ProductCrowdsTrendList_noid)
	}

	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("  SQL_CROWD_ProductCrowdsTrendList sql\n", mainSql)
	e := db.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_CROWD_ProductCrowdsTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespProductCrowdsTrendList) Prodata() {

}

const SQL_CROWD_ProductCrowdsTrendList = `
 
select   t.year_month as date ,  product_id, product_name,  sum(paid_amount) as gmv, 
sum(visitors) as visitors_count,
0 as crowd_tgi, sum(paid_amount) /sum(paid_buyers)  as customer_unit_price,
sum(paid_buyers)/sum(visitors) as payment_conversion_rate 
-- sum(add_to_cart_count)/sum(visitors) as add_car_rate
from biz_shop_audience_pruduct_t10 t 
WHERE t.year_month >=:startdate  and t.year_month <= :enddate   
and product_id in :ids 
and crowd_type =:crowd_type 
group by t.year_month, product_id, product_name   -- crowd_type, 
LIMIT :offset , :pageSize  

`

const SQL_CROWD_ProductCrowdsTrendList_noid = `

select   t.year_month as date ,   sum(paid_amount) as gmv, 
sum(visitors) as visitors_count,
0 as crowd_tgi, sum(paid_amount) /sum(paid_buyers)  as customer_unit_price,
sum(paid_buyers)/sum(visitors) as payment_conversion_rate 
-- sum(add_to_cart_count)/sum(visitors) as add_car_rate
from biz_shop_audience_pruduct_t10 t 
WHERE t.year_month >=:startdate  and t.year_month <= :enddate   
and crowd_type =:crowd_type 
group by t.year_month    -- crowd_type, 
LIMIT :offset , :pageSize  

`
