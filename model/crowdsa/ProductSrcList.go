package crowdsa

import (
	"errors"
	"fmt"
	"xtt/global"
	base "xtt/model/base"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductSrcList) GetData(r ReqCrowdSaAllSearch) error {
	db := base.GetDBByShopName(r.ShopName)
	if db == nil {
		global.GVA_LOG.Error("get db by shop error", zap.String("ShopName", r.ShopName))
		return errors.New("get db by shop error")
	}
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_CROWD_ProductSrcList)
	if r.SecondarySource != "" {
		sqlp.SetSql(SQL_CROWD_ProductSrcList_3)
	}
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("  sql\n", mainSql)
	e := db.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("  sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespProductSrcList) Prodata() {

}

const SQL_CROWD_ProductSrcList = `
	select  source_type_2 as secondary_source, sum(paid_amount) as gmv, sum(visitors_count) as visitors_count,
	0 as crowd_tgi, sum(paid_amount) /sum(direct_paid_buyers)  as customer_unit_price,
	sum(direct_paid_buyers)/sum(visitors_count) as payment_conversion_rate  
	-- sum(add_to_carts)/sum(visitors_count) as add_car_rate
	from biz_product_traffic_stats t 
	WHERE t.statistic_date >=:startdate  and t.statistic_date <= LAST_DAY(:enddate  ) 
	 
	and product_id in :ids 
	and src="每一次访问来源"
	-- and source_type_2=:secondary_source  and source_type_3=:tertiary_source 
	group by source_type_2
	LIMIT :offset , :pageSize  
	
`

const SQL_CROWD_ProductSrcList_3 = `
	select  source_type_2 as secondary_source, source_type_3 as tertiary_source, sum(paid_amount) as gmv, 
	sum(visitors_count) as visitors_count,
	0 as crowd_tgi, sum(paid_amount) /sum(direct_paid_buyers)  as customer_unit_price,
	sum(direct_paid_buyers)/sum(visitors_count) as payment_conversion_rate 
	-- sum(add_to_carts)/sum(visitors_count) as add_car_rate
	from biz_product_traffic_stats t 
	WHERE t.statistic_date >=:startdate  and t.statistic_date <= LAST_DAY(:enddate  ) 
	 
	and product_id in :ids 
	and src="每一次访问来源"
	and source_type_2=:secondary_source  
	group by source_type_2, source_type_3
	LIMIT :offset , :pageSize  
`
