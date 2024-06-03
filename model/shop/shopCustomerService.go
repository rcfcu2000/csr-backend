package shop

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *CustomerService) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_CustomerService)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("CustomerService sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("CustomerService sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *CustomerService) Prodata() {

}

const SQL_SHOP_CustomerService = `

 
with  t1 as (

	select 
	sum(sales_revenue) as  "customer_service_sales",
	sum(sales_revenue)  as sales_revenue,
	avg(avg_response_time) as  "average_response_time_in_seconds",
	avg(customer_satisfaction_rate) as  "customer_satisfaction_rate", 
	avg(inquiry_conversion_rate) as  "inquiry_conversion_rate" 
	from biz_shop_customer_service t
	where (t.shop_name=:shop_name ) and t.statistic_date >=:startdate   and t.statistic_date <= :enddate 
	
	),t2 as (
	select sum(paid_amount) as paid_amount
	from biz_product_performance t
	where (t.shop_name=:shop_name ) and t.statistic_date >=:startdate   and t.statistic_date <= :enddate     
	)
	
	select 
	t1.customer_service_sales,
	t1.sales_revenue / t2.paid_amount as "customer_service_sales_ratio",
	t1.average_response_time_in_seconds,
	t1.customer_satisfaction_rate,
	t1.inquiry_conversion_rate
	
	from t1  left join t2 on 1=1

`

// select
// sales_revenue as  "customer_service_sales",
// -- sales_revenue /(select sum(paid_amount)
// -- from biz_product_performance
// -- where (t.shop_name=:shop_name ) and t.statistic_date=:enddate ) as "customer_service_sales_ratio",

// sales_revenue_ratio  as "customer_service_sales_ratio",
// avg_response_time as  "average_response_time_in_seconds",
// customer_satisfaction_rate as  "customer_satisfaction_rate",
// inquiry_conversion_rate as  "inquiry_conversion_rate"

// from biz_shop_customer_service t
// where (t.shop_name=:shop_name ) and t.statistic_date=:enddate ;

/////////////////////////////////////////////////////////////////
//

func (i *CustomerServiceTrendList) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_CustomerServiceTrendList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []CustomerServiceTrendNode{}

	fmt.Println("CustomerServiceTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("CustomerServiceTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *CustomerServiceTrendList) Prodata() {

}

const SQL_SHOP_CustomerServiceTrendList = `

select  DATE_FORMAT(t.statistic_date, '%Y-%m-%d')  as date, 
sales_revenue as  "customer_service_sales",
-- 0 as "customer_service_sales_ratio",
-- sales_revenue/ paid_amount as "customer_service_sales_ratio",
sales_revenue_ratio  as "customer_service_sales_ratio",
 
avg_response_time as  "average_response_time_in_seconds",
customer_satisfaction_rate as  "customer_satisfaction_rate", 
inquiry_conversion_rate as  "inquiry_conversion_rate" 
  
from biz_shop_customer_service t
left join(
	 select statistic_date, sum(paid_amount) as paid_amount 
	 from biz_product_performance   bpp
	 where   bpp.statistic_date>=:startdate  and bpp.statistic_date<=:enddate 
   GROUP BY statistic_date
) bpp on (t.statistic_date =  bpp.statistic_date)

where t.statistic_date>=:startdate and t.statistic_date<=:enddate and t.shop_name=:shop_name ;

`
