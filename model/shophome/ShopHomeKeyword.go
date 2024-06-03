package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomeKeywordData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_Keyword_ztc)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_Keyword sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.ZtcRecords).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_Keyword sql\n", mainSql)
		return e
	}
	i.Prodata(r)
	return nil

}
func (i *RespShopHomeKeywordData) Prodata(r ReqShopHomeAllSearch) {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_Keyword_search)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_Keyword sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.SearchRecords).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_Keyword sql\n", mainSql)
		return
	}
}

const SQL_SHOPHOME_Keyword = `
WITH t1 AS (
	SELECT
		src_type AS src,
		keyword AS keyword,
		sum(visitor_count) AS visitors
   from biz_product_keyword
	WHERE
		src_type = '直通车'
	AND statistic_date >= :startdate  
	AND statistic_date <= :enddate  
	GROUP BY
		src_type,
		keyword
order BY sum(visitor_count) desc LIMIT 0, 30
),
 t2 AS (
	SELECT
		src_type AS src,
		keyword AS keyword,
		sum(visitor_count) AS visitors from biz_product_keyword  
where src_type = '手淘搜索'and statistic_date >= :startdate   and  statistic_date <= :enddate  
group by src_type,keyword 
order by sum(visitor_count) desc 
LIMIT 0, 30
)

select src,keyword,visitors from t1 union all select src,keyword,visitors
from t2
`

const SQL_SHOPHOME_Keyword_ztc = `
WITH t1 AS (
	SELECT
		src_type AS src,
		keyword AS keyword,
		sum(visitor_count) AS visitors
   from biz_product_keyword
	WHERE
		src_type = '直通车'
	AND statistic_date >= :startdate  
	AND statistic_date <= :enddate  
	GROUP BY
		src_type,
		keyword
order BY sum(visitor_count) desc LIMIT 0, 30
),
 t2 AS (
	SELECT
		src_type AS src,
		keyword AS keyword,
		sum(visitor_count) AS visitors from biz_product_keyword  
where src_type = '手淘搜索' and statistic_date >= :startdate   and  statistic_date <= :enddate  
group by src_type,keyword 
order by sum(visitor_count) desc 
LIMIT 0, 30
)

select src,keyword,visitors from t1  
`

const SQL_SHOPHOME_Keyword_search = `
WITH t1 AS (
	SELECT
		src_type AS src,
		keyword AS keyword,
		sum(visitor_count) AS visitors
   from biz_product_keyword
	WHERE
		src_type = '直通车'
	AND statistic_date >= :startdate  
	AND statistic_date <= :enddate  
	GROUP BY
		src_type,
		keyword
order BY sum(visitor_count) desc LIMIT 0, 30
),
 t2 AS (
	SELECT
		src_type AS src,
		keyword AS keyword,
		sum(visitor_count) AS visitors from biz_product_keyword  
where src_type = '手淘搜索' and statistic_date >= :startdate   and  statistic_date <= :enddate  
group by src_type,keyword 
order by sum(visitor_count) desc 
LIMIT 0, 30
)

select src,keyword,visitors from t2  
`
