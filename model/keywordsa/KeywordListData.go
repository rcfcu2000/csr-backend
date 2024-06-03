package keywordsa

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespKeywordListData) GetData(r ReqKeywordSaAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_KEYWORD_SA_KeywordList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println(" SQL_KEYWORD_SA_KeywordList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_KEYWORD_SA_KeywordList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespKeywordListData) Prodata() {

}

const SQL_KEYWORD_SA_KeywordList = `

select keyword_name as keyword, sum(clicktraffic) as visitors_count, IFNULL( sum(buyer_count)/sum(clicktraffic),0.0)  as payment_conversion_rate
from wanxiang_keywords t 
where t.datetimekey>=:startdate and t.datetimekey<=:enddate 
GROUP BY keyword_name

`
