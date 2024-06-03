package keywordsa

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespKeywordTrendListData) GetData(r ReqKeywordSaAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_KEYWORD_SA_KeywordtTrendList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println(" SQL_KEYWORD_SA_KeywordtTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_KEYWORD_SA_KeywordtTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespKeywordTrendListData) Prodata() {

}

const SQL_KEYWORD_SA_KeywordtTrendList = `

WITH t1 AS (
		select datetimekey as date,  sum(clicktraffic) as visitors_count, IFNULL( sum(buyer_count)/sum(clicktraffic),0.0)  as payment_conversion_rate
		from wanxiang_keywords t 
		where t.datetimekey>=:startdate   and t.datetimekey<=:enddate  
		and IF(':keyword_like' = '', TRUE, keyword_name LIKE '%:keyword_like%')
		GROUP BY datetimekey -- , keyword_name
	),
	t2 as(
		select statistic_date as date,   sum(visitor_count) as visitors_count, IFNULL( sum(direct_payment_buyer_count)/sum(visitor_count),0.0)  as payment_conversion_rate
		from biz_shop_keyword t 
		where t.statistic_date>=:startdate   and t.statistic_date<=:enddate  
		and src_type ="手淘搜索"
		and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
		GROUP BY statistic_date --  keyword_name
	),
	t3 as(
		select statistic_date as date,   sum(visitor_count) as visitors_count, IFNULL( sum(direct_payment_buyer_count)/sum(visitor_count),0.0)  as payment_conversion_rate
		from biz_shop_keyword t 
		where t.statistic_date>=:startdate   and t.statistic_date<=:enddate  
		and src_type ="直通车"
		and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
		GROUP BY statistic_date --  keyword_name
 
	)
	,
	t4 as(  -- 行业
		select 0 as visitors_count, 0 as cr
	)
	
	select  
	DATE_FORMAT(t1.date, '%m-%d')  as date,
	t1.visitors_count as clicks, t2.visitors_count as visitors_count_free, t3.visitors_count as visitors_count_notfree, 0 as industry_clicks,
	t1.payment_conversion_rate as cr, t2.payment_conversion_rate as cr_free, t3.payment_conversion_rate as cr_notfree, 0 as cr_industry
	from t1  
	left join t2 on t1.date = t2.date -- and t1.keyword_name = t2.keyword_name
	left join t3 on t1.date = t3.date -- and t1.keyword_name = t2.keyword_name
	order by t1.date 

`

// select keyword_name, sum(clicktraffic) as visitors_count, IFNULL( sum(buyer_count)/sum(clicktraffic),0.0)  as payment_conversion_rate
// from wanxiang_keywords t
// where t.datetimekey>="2024-01-01" and t.datetimekey<="2024-01-11"
// GROUP BY keyword_name
