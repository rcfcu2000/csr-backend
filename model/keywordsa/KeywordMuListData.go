package keywordsa

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

// 明细表格
func (i *RespKeywordMuListData) GetData(r ReqKeywordSaAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_KEYWORDSA_MuList)
	if r.Keyword != "" {
		sqlp.SetSql(SQL_KEYWORD_Detail)
	} else {
		i.ProdataSum(r)
	}
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println(" SQL_KEYWORD_SA_KeywordMuList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_KEYWORD_SA_KeywordMuList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespKeywordMuListData) Prodata() {

}
func (i *RespKeywordMuListData) ProdataSum(r ReqKeywordSaAllSearch) {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_KEYWORDSA_MuList_Sum)

	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	fmt.Println(" SQL_KEYWORDSA_MuList_Sum sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Sum).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println(" SQL_KEYWORDSA_MuList_Sum sql\n", mainSql)
		return
	}

}

const SQL_KEYWORDSA_MuList = `

WITH t1 AS (
	select keyword_name as keyword, 
	round(sum(clicktraffic),0) as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	sum(buyer_count) as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from wanxiang_keywords t 
	-- where t.datetimekey>="2024-01-01" and t.datetimekey<="2024-01-11" 
	where t.datetimekey>=:startdate  and t.datetimekey<=:enddate  
	-- and IF(':keyword_like' = '', TRUE, keyword_name LIKE '%:keyword_like%')
	GROUP BY keyword_name 
	having sum(clicktraffic)+sum(buyer_count)>0
	union all 
	select 
	keyword as keyword, 
	0 as clicks,  -- 无界词-点击量
	sum(visitor_count) as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	sum(direct_payment_buyer_count) as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from biz_shop_keyword t 
	--  where t.statistic_date>="2024-01-01" and t.statistic_date<="2024-01-11" 
	where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
	and src_type ="手淘搜索"
	-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
	GROUP BY keyword 
		having sum(visitor_count)+sum(direct_payment_buyer_count)>0
	union all 
		select 
		keyword as keyword,
		0 as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	sum(visitor_count) as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	sum(direct_payment_buyer_count) as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from biz_shop_keyword t 
	-- where t.statistic_date>="2024-01-01" and t.statistic_date<="2024-01-11" 
	where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
	and src_type ="直通车"
	-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
	GROUP BY keyword
	having sum(visitor_count)+sum(direct_payment_buyer_count)>0
	union all 
	select keyword, 
	0 as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	sum(click_count) as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	sum( click_count*conversion_rate ) as buyer_industry -- 行业-买家数
from  biz_industry_keyword 
	--  where statistic_date>="2024-03-01" and statistic_date<="2024-03-11" 
	where statistic_date>=:startdate  and statistic_date<=:enddate  
-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
group by keyword
	having sum(click_count)+sum(click_count*conversion_rate)>0
),
t2 as (
select  
keyword as keyword,
1 as count, 
	sum(clicks) as clicks,  -- 无界词-点击量
	sum(visitors_count_free) as visitors_count_free, -- 生参免费词访客
	sum(visitors_count_notfree) as visitors_count_notfree, -- 生参付费词访客
	sum(industry_clicks) as industry_clicks, -- 行业点击量
	sum(buyer) as buyer, 
	sum(buyer_free) as buyer_free,
	sum(buyer_notfree) as buyer_notfree, 
	sum(buyer_industry) as buyer_industry, 
	sum(buyer)/sum(clicks) as cr, -- 无界词转化率 cr =  conversion rate
	sum(buyer_free)/sum(visitors_count_free) as cr_free, -- 生参免费词转化率
	sum(buyer_notfree)/sum(visitors_count_notfree) as cr_notfree, -- 生参付费词转化率
	sum(buyer_industry)/sum(industry_clicks) as cr_industry -- 行业-转化率
from t1  
group by keyword
)
select 
case 
when clicks+visitors_count_free+visitors_count_notfree = 0 and industry_clicks > 1000 then '机会词'
when visitors_count_free+visitors_count_notfree+industry_clicks = 0  then '流量词'
when visitors_count_notfree > 0 and clicks = 0  then '智能词'
when visitors_count_free > 0 and industry_clicks > 0 and clicks = 0 and visitors_count_notfree = 0  then '潜力词'
else '其他' end as   keyword,  -- key_type,
sum(count) as count, 
	sum(clicks) as clicks,  -- 无界词-点击量
	sum(visitors_count_free) as visitors_count_free, -- 生参免费词访客
	sum(visitors_count_notfree) as visitors_count_notfree, -- 生参付费词访客
	sum(industry_clicks) as industry_clicks, -- 行业点击量
	sum(buyer)/sum(clicks) as cr, -- 无界词转化率 cr =  conversion rate
	sum(buyer_free)/sum(visitors_count_free) as cr_free, -- 生参免费词转化率
	sum(buyer_notfree)/sum(visitors_count_notfree) as cr_notfree, -- 生参付费词转化率
	sum(buyer_industry)/sum(industry_clicks) as cr_industry -- 行业-转化率
from t2
group by 
case 
when clicks+visitors_count_free+visitors_count_notfree = 0 and industry_clicks > 1000 then '机会词'
when visitors_count_free+visitors_count_notfree+industry_clicks = 0  then '流量词'
when visitors_count_notfree > 0 and clicks = 0  then '智能词'
when visitors_count_free > 0 and industry_clicks > 0 and clicks = 0 and visitors_count_notfree = 0  then '潜力词'
else '其他' end
`

const SQL_KEYWORD_Detail = `
WITH t1 AS (
	select keyword_name as keyword, 
	round(sum(clicktraffic),0) as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	sum(buyer_count) as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from wanxiang_keywords t 
	-- where t.datetimekey>="2024-01-01" and t.datetimekey<="2024-01-11" 
	where t.datetimekey>=:startdate  and t.datetimekey<=:enddate  
	-- and IF(':keyword_like' = '', TRUE, keyword_name LIKE '%:keyword_like%')
	GROUP BY keyword_name 
	having sum(clicktraffic)+sum(buyer_count)>0
	union all 
	select 
	keyword as keyword, 
	0 as clicks,  -- 无界词-点击量
	sum(visitor_count) as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	sum(direct_payment_buyer_count) as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from biz_shop_keyword t 
	-- where t.statistic_date>="2024-01-01" and t.statistic_date<="2024-01-11" 
	where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
	and src_type ="手淘搜索"
	-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
	GROUP BY keyword 
		having sum(visitor_count)+sum(direct_payment_buyer_count)>0
	union all 
		select 
		keyword as keyword,
		0 as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	sum(visitor_count) as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	sum(direct_payment_buyer_count) as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from biz_shop_keyword t 
	-- where t.statistic_date>="2024-01-01" and t.statistic_date<="2024-01-11" 
	where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
	and src_type ="直通车"
	-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
	GROUP BY keyword
	having sum(visitor_count)+sum(direct_payment_buyer_count)>0
	union all 
	select keyword, 
	0 as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	sum(click_count) as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	sum( click_count*conversion_rate ) as buyer_industry -- 行业-买家数
from  biz_industry_keyword 
	-- where statistic_date>="2024-03-01" and statistic_date<="2024-03-11" 
 	where statistic_date>=:startdate  and statistic_date<=:enddate  
-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
group by keyword
	having sum(click_count)+sum(click_count*conversion_rate)>0
),
t2 as (
select  
keyword as keyword,
1 as count, 
	sum(clicks) as clicks,  -- 无界词-点击量
	sum(visitors_count_free) as visitors_count_free, -- 生参免费词访客
	sum(visitors_count_notfree) as visitors_count_notfree, -- 生参付费词访客
	sum(industry_clicks) as industry_clicks, -- 行业点击量
	sum(buyer) as buyer, 
	sum(buyer_free) as buyer_free,
	sum(buyer_notfree) as buyer_notfree, 
	sum(buyer_industry) as buyer_industry, 
	sum(buyer)/sum(clicks) as cr, -- 无界词转化率 cr =  conversion rate
	sum(buyer_free)/sum(visitors_count_free) as cr_free, -- 生参免费词转化率
	sum(buyer_notfree)/sum(visitors_count_notfree) as cr_notfree, -- 生参付费词转化率
	sum(buyer_industry)/sum(industry_clicks) as cr_industry -- 行业-转化率
from t1  
group by keyword
)
 
select keyword,
count, 
clicks,  -- 无界词-点击量
visitors_count_free, -- 生参免费词访客
 visitors_count_notfree, -- 生参付费词访客
industry_clicks, 
cr, -- 无界词转化率 cr =  conversion 
cr_free, -- 生参免费词
cr_notfree, 
cr_industry    -- 行业-转化率
from t2

where 
case 
when clicks+visitors_count_free+visitors_count_notfree = 0 and industry_clicks > 1000 then '机会词'
when visitors_count_free+visitors_count_notfree+industry_clicks = 0  then '流量词'
when visitors_count_notfree > 0 and clicks = 0  then '智能词'
when visitors_count_free > 0 and industry_clicks > 0 and clicks = 0 and visitors_count_notfree = 0  then '潜力词'
else '其他' end   =:keyword   -- '智能词'
LIMIT :offset , :pageSize  
`

const SQL_KEYWORDSA_MuList_Sum = `

WITH t1 AS (
	select keyword_name as keyword, 
	round(sum(clicktraffic),0) as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	sum(buyer_count) as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from wanxiang_keywords t 
	-- where t.datetimekey>="2024-01-01" and t.datetimekey<="2024-01-11" 
	where t.datetimekey>=:startdate  and t.datetimekey<=:enddate  
	-- and IF(':keyword_like' = '', TRUE, keyword_name LIKE '%:keyword_like%')
	GROUP BY keyword_name 
	having sum(clicktraffic)+sum(buyer_count)>0
	union all 
	select 
	keyword as keyword, 
	0 as clicks,  -- 无界词-点击量
	sum(visitor_count) as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	sum(direct_payment_buyer_count) as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from biz_shop_keyword t 
	--  where t.statistic_date>="2024-01-01" and t.statistic_date<="2024-01-11" 
	where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
	and src_type ="手淘搜索"
	-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
	GROUP BY keyword 
		having sum(visitor_count)+sum(direct_payment_buyer_count)>0
	union all 
		select 
		keyword as keyword,
		0 as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	sum(visitor_count) as visitors_count_notfree, -- 生参付费词访客
	0 as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	sum(direct_payment_buyer_count) as buyer_notfree, -- 生参付费词买家数
	0 as buyer_industry -- 行业-买家数
	from biz_shop_keyword t 
	-- where t.statistic_date>="2024-01-01" and t.statistic_date<="2024-01-11" 
	where t.statistic_date>=:startdate  and t.statistic_date<=:enddate  
	and src_type ="直通车"
	-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
	GROUP BY keyword
	having sum(visitor_count)+sum(direct_payment_buyer_count)>0
	union all 
	select keyword, 
	0 as clicks,  -- 无界词-点击量
	0 as visitors_count_free, -- 生参免费词访客
	0 as visitors_count_notfree, -- 生参付费词访客
	sum(click_count) as industry_clicks, -- 行业点击量
	0 as buyer, -- 无界词买家数
	0 as buyer_free, -- 生参免费词买家数
	0 as buyer_notfree, -- 生参付费词买家数
	sum( click_count*conversion_rate ) as buyer_industry -- 行业-买家数
from  biz_industry_keyword 
	--  where statistic_date>="2024-03-01" and statistic_date<="2024-03-11" 
	where statistic_date>=:startdate  and statistic_date<=:enddate  
-- and IF(':keyword_like' = '', TRUE, keyword LIKE '%:keyword_like%')
group by keyword
	having sum(click_count)+sum(click_count*conversion_rate)>0
),
t2 as (
select  
keyword as keyword,
1 as count, 
	sum(clicks) as clicks,  -- 无界词-点击量
	sum(visitors_count_free) as visitors_count_free, -- 生参免费词访客
	sum(visitors_count_notfree) as visitors_count_notfree, -- 生参付费词访客
	sum(industry_clicks) as industry_clicks, -- 行业点击量
	sum(buyer) as buyer, 
	sum(buyer_free) as buyer_free,
	sum(buyer_notfree) as buyer_notfree, 
	sum(buyer_industry) as buyer_industry, 
	sum(buyer)/sum(clicks) as cr, -- 无界词转化率 cr =  conversion rate
	sum(buyer_free)/sum(visitors_count_free) as cr_free, -- 生参免费词转化率
	sum(buyer_notfree)/sum(visitors_count_notfree) as cr_notfree, -- 生参付费词转化率
	sum(buyer_industry)/sum(industry_clicks) as cr_industry -- 行业-转化率
from t1  
group by keyword
)
select 
 
sum(count) as count, 
	sum(clicks) as clicks,  -- 无界词-点击量
	sum(visitors_count_free) as visitors_count_free, -- 生参免费词访客
	sum(visitors_count_notfree) as visitors_count_notfree, -- 生参付费词访客
	sum(industry_clicks) as industry_clicks, -- 行业点击量
	sum(buyer)/sum(clicks) as cr, -- 无界词转化率 cr =  conversion rate
	sum(buyer_free)/sum(visitors_count_free) as cr_free, -- 生参免费词转化率
	sum(buyer_notfree)/sum(visitors_count_notfree) as cr_notfree, -- 生参付费词转化率
	sum(buyer_industry)/sum(industry_clicks) as cr_industry -- 行业-转化率
from t2
 
`
