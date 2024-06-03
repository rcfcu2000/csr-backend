package traffic

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespNewOldCustomerListData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_NewOld)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespNewOldCustomerListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespNewOldCustomerListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespNewOldCustomerListData) Prodata() {

}

const SQL_TRAFFIC_NewOld = `

select bsc.statistic_date as "date" , total_customers, new_visits, 
new_visits/total_customers as "new_visits_percentage",  
non_purchase_return_visits/total_customers as "non_purchase_return_visits_percentage", 
purchased_customer_return_visits/total_customers as "purchased_customer_return_visits_percentage", 
(new_visit_conversions+ return_visit_conversions_non_purchasers + repeat_purchases )/total_customers as "conversion_rate", 
return_visit_payment_conversion_rate_non_purchasers,
new_visit_payment_conversion_rate, 
return_visit_payment_conversion_rate_purchasers


from biz_shop_customers bsc
where statistic_date >= "2024-01-01"  and statistic_date<="2024-01-11" 
order by statistic_date
 

`
