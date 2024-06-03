package shop

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *CustomerLossAnalysis) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_CustomerLossAnalysis)
	(&r).SetToSQLProccesor(sqlp)

	mainSql := sqlp.GetResult()

	fmt.Println("CustomerLossAnalysis sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("CustomerLossAnalysis sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *CustomerLossAnalysis) Prodata() {

}

const SQL_SHOP_CustomerLossAnalysis = `

select 
churn_amount as  "amount_of_loss",
churn_count as  "lost_members",
shop_count_involved as  "stores_causing_loss" 
  
from biz_shop_competition t
where (t.shop_name=:shop_name ) and t.statistic_date=:enddate ;


`

///////////////////////////////////////////////////////////////////////////
//

func (i *CustomerLossAnalysisTrendList) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_CustomerLossAnalysisTrendList)
	(&r).SetToSQLProccesor(sqlp)

	mainSql := sqlp.GetResult()
	i.Records = []CustomerLossAnalysisTrendNode{}

	fmt.Println("CustomerLossAnalysisTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("CustomerLossAnalysisTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *CustomerLossAnalysisTrendList) Prodata() {

}

const SQL_SHOP_CustomerLossAnalysisTrendList = `

select  DATE_FORMAT(t.statistic_date, '%Y-%m-%d')  as date,  
churn_amount as  "amount_of_loss",
churn_count as  "lost_members",
shop_count_involved as  "stores_causing_loss" 
  
from biz_shop_competition t
where t.statistic_date>=:startdate and t.statistic_date<=:enddate and t.shop_name=:shop_name ;

`
