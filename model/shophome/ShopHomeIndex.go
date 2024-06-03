package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomeIndexData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_Index)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_Index sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_Index sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespShopHomeIndexData) Prodata() {

}

const SQL_SHOPHOME_Index = `

with t1 as (
	select 
	IFNULL(sum(visitors_count),0) as visitors,
	IFNULL(sum(paid_amount),0) as gmv,
	IFNULL( sum(paid_amount)/sum(paid_buyers),0) as customer_unit_price,
	IFNULL( sum(paid_buyers)/sum(visitors_count),0) as conversion_rate_payment,
	sum(bpp.paid_amount- bpp.successful_refund_amount)	AS gmv_refund,
	sum((bpp.paid_amount- bpp.successful_refund_amount) * bp.estimated_gross_profit_margin
	- bpp.paid_quantity* bp.delivery_cost) AS jlr1
	from biz_product_performance bpp
	LEFT JOIN biz_product bp ON bpp.product_id = bp.product_id 
	where bpp.statistic_date >= :startdate 
	and  bpp.statistic_date <= :enddate 
	),t2 as 
	(
	SELECT
	sum( wxp.spend ) AS spend,
	sum( wxp.gmv ) AS wxp_gmv
	FROM
	wanxiang_product wxp
	where wxp.datetimekey >= :startdate 
	and  wxp.datetimekey <= :enddate 
	)
	select t1.visitors,
	t1.gmv,
	t1.customer_unit_price,
	t1.conversion_rate_payment,
	t2.spend,
	t2.wxp_gmv/t1.gmv as promotion_percentage,
	t1.jlr1-t2.spend as profit,
	(t1.jlr1-t2.spend)/t1.gmv_refund as profit_rate
	from t1 left join t2 on 1=1

`
