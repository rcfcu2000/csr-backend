package crowdsa

import (
	"errors"
	"fmt"
	"xtt/global"
	base "xtt/model/base"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespCrowdGmv20List) GetData(r ReqCrowdSaAllSearch) error {

	db := base.GetDBByShopName(r.ShopName)
	if db == nil {
		global.GVA_LOG.Error("get db by shop error", zap.String("ShopName", r.ShopName))
		return errors.New("get db by shop error")
	}
	if len(r.StartDate) > len("2024-01") {
		r.StartDate = r.StartDate[0:7]
		r.EndDate = r.EndDate[0:7]
	}
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_CROWD_Gmv20)
	if r.SecondarySource == "" && r.TertiarySource == "" {
		sqlp.SetSql(SQL_CROWD_Gmv20_bank)
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
func (i *RespCrowdGmv20List) Prodata() {

}

const SQL_CROWD_Gmv20 = `
	select  crowd_type, sum(paid_amount) as gmv, sum(visitors) as visitors_count,
	0 as crowd_tgi, sum(paid_amount) /sum(paid_buyers)  as customer_unit_price,
	sum(paid_buyers)/sum(visitors) as payment_conversion_rate 
	-- sum(add_to_cart_count)/sum(visitors) as add_car_rate
	from biz_shop_audience_channel_t20 t 
	WHERE t.year_month >=:startdate  and t.year_month <= :enddate  
	and secondary_source=:secondary_source  and (tertiary_source=:tertiary_source or  :tertiary_source ="" )
	group by crowd_type 

`

const SQL_CROWD_Gmv20_bank = `
	select  crowd_type, sum(paid_amount) as gmv, sum(visitors) as visitors_count,
	0 as crowd_tgi, sum(paid_amount) /sum(paid_buyers)  as customer_unit_price,
	sum(paid_buyers)/sum(visitors) as payment_conversion_rate
	-- sum(add_to_cart_count)/sum(visitors) as add_car_rate
	from biz_shop_audience_channel_t20 t 
	WHERE t.year_month >=:startdate  and t.year_month <= :enddate  
	-- and secondary_source="" and tertiary_source=""
	group by crowd_type
	LIMIT :offset , :pageSize  

`
