package keywordsa

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespScKeywordListData) GetData(r ReqKeywordSaAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_KEYWORD_SA_ScKeywordListFree)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println(" SQL_KEYWORD_SA_ScKeywordListFree sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.RecordsFree).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_KEYWORD_SA_ScKeywordListFree sql\n", mainSql)
		return e
	}

	sqlp.SetSql(SQL_KEYWORD_SA_ScKeywordListNotFree)
	(&r).SetToSQLProccesor(sqlp)
	mainSql = sqlp.GetResult()

	fmt.Println(" SQL_KEYWORD_SA_ScKeywordListNotFree sql\n", mainSql)
	e = global.GVA_DB.Raw(mainSql).Scan(&i.RecordsNotFree).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_KEYWORD_SA_ScKeywordListNotFree sql\n", mainSql)
		return e
	}

	i.Prodata()
	return nil

}
func (i *RespScKeywordListData) Prodata() {

}

const SQL_KEYWORD_SA_ScKeywordListFree = `

select keyword as keyword,   sum(visitor_count) as visitors_count, IFNULL( sum(direct_payment_buyer_count)/sum(visitor_count),0.0)  as payment_conversion_rate
from biz_shop_keyword t 
where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
and src_type ="手淘搜索"
GROUP BY keyword

`

const SQL_KEYWORD_SA_ScKeywordListNotFree = `

select keyword as keyword,   sum(visitor_count) as visitors_count, IFNULL( sum(direct_payment_buyer_count)/sum(visitor_count),0.0)  as payment_conversion_rate
from biz_shop_keyword t 
where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
and src_type ="直通车"
GROUP BY keyword

`
