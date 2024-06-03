package keywordsa

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespIndustryKeywordListData) GetData(r ReqKeywordSaAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_IndustryKeywordList)

	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("  sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("  sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespIndustryKeywordListData) Prodata() {

}

const SQL_IndustryKeywordList = `
select keyword  as keyword, sum(visitor_count) as visitors_count, IFNULL( sum(visitor_count*conversion_rate)/sum(visitor_count),0.0)  as payment_conversion_rate
-- IFNULL( sum(buyer_count)/sum(clicktraffic),0.0)  as payment_conversion_rate
from biz_industry_keyword t 
where t.statistic_date>=:startdate and t.statistic_date<=:enddate  
GROUP BY keyword
`
