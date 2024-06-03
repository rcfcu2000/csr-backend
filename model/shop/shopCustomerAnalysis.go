package shop

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *CustomerAnalysis) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_CustomerAnalysis)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("CustomerAnalysis sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("CustomerAnalysis sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *CustomerAnalysis) Prodata() {

}

const SQL_SHOP_CustomerAnalysis = `

select members_count as  "total_membership_count",
members_yesterday as  "members_recruited_yesterday",
membership_transaction_penetration as  "member_transaction_penetration",
industry_standard_penetration as  "industry_standard_penetration",
industry_premium_penetration as  "industry_premium_penetration",
average_member_order_value as  "member_average_order_value" 
 
from biz_shop_member  
where    statistic_date=:enddate ;

`

////////////////////////////////////////////////////////////////////////////
//

func (i *CustomerAnalysisTrendList) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_CustomerAnalysisTrend)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	i.Records = []CustomerAnalysisTrendNode{}

	fmt.Println("CustomerAnalysisTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("CustomerAnalysisTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *CustomerAnalysisTrendList) Prodata() {

}

const SQL_SHOP_CustomerAnalysisTrend = `

select 
DATE_FORMAT( statistic_date, '%Y-%m-%d')  as date, 
members_count as  "total_membership_count",
members_yesterday as  "members_recruited_yesterday",
membership_transaction_penetration as  "member_transaction_penetration",
industry_standard_penetration as  "industry_standard_penetration",
industry_premium_penetration as  "industry_premium_penetration",
average_member_order_value as  "member_average_order_value" 
 
from   biz_shop_member
where  statistic_date>=:startdate and statistic_date<=:enddate  ;
 
`
