package product

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

type DCount struct {
	Count int64 `json:"count"`
}

func (i *RespKeywordList) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_Keywordlist)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []KeywordNode{}

	fmt.Println("RespKeywordList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespKeywordList sql\n", mainSql)
		return e
	}
	i.Prodata()
	i.GetDataSum(r)
	//i.GetDataCount(r)
	return nil

}
func (i *RespKeywordList) Prodata() {

}

func (i *RespKeywordList) GetDataSum(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_Keywordlist_Sum)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Sum = []KeywordNode{}

	fmt.Println("RespKeywordList sum sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Sum).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespKeywordList sum sql\n", mainSql)
		return e
	}
	return nil
}

func (i *RespKeywordList) GetDataCount(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_Keywordlist_count)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	var count DCount

	fmt.Println("RespKeywordList count sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&count).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespKeywordList count sql\n", mainSql)
		return e
	}
	i.Count = int64(count.Count)
	return nil
}

const SQL_PRODUCT_Keywordlist_count = `

select  
COUNT(DISTINCT v.src_type,v.keyword) as count
from v_bpk_bp v
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )

`
const SQL_PRODUCT_Keywordlist_Sum = `
SELECT
 
SUM(case when v.src_type = '手淘搜索' then v.visitor_count else 0 end) AS search_visitor_count,
SUM(case when v.src_type = '直通车' then v.visitor_count else 0 end) AS ztc_visitor_count,
SUM(v.visitor_count) AS zj_visitor_count,
SUM(case when v.src_type = '手淘搜索' then v.cart_addition_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '手淘搜索' then v.visitor_count else 0 end) AS search_cart_addition_rate,
SUM(case when v.src_type = '直通车' then v.cart_addition_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '直通车' then v.visitor_count else 0 end) AS ztc_cart_addition_rate,
SUM(v.cart_addition_rate*v.visitor_count)/SUM(v.visitor_count) AS zj_cart_addition_rate,
SUM(case when v.src_type = '手淘搜索' then v.conversion_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '手淘搜索' then v.visitor_count else 0 end) AS search_conversion_rate,
SUM(case when v.src_type = '直通车' then v.conversion_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '直通车' then v.visitor_count else 0 end) AS ztc_conversion_rate,   
SUM(v.conversion_rate*v.visitor_count)/SUM(v.visitor_count) AS zj_conversion_rate, 
SUM(case when v.src_type = '手淘搜索' then v.fan_payment_buyer_count else 0 end) AS search_fans_paid_buyers_count,
SUM(case when v.src_type = '直通车' then v.fan_payment_buyer_count else 0 end) AS ztc_fans_paid_buyers_count,
SUM(v.fan_payment_buyer_count) AS zj_fans_paid_buyers_count,
SUM(case when v.src_type = '手淘搜索' then v.direct_payment_buyer_count else 0 end) AS search_direct_paid_buyers_count,
SUM(case when v.src_type = '直通车' then v.direct_payment_buyer_count else 0 end) AS ztc_direct_paid_buyers_count,
SUM(v.direct_payment_buyer_count) AS zj_direct_paid_buyers_count
FROM
biz_product_keyword v
WHERE 
v.product_id = :productid and v.statistic_date >=:startdate and v.statistic_date <= :enddate  
 
`

const SQL_PRODUCT_Keywordlist = `
SELECT
v.keyword as "keyword",
SUM(case when v.src_type = '手淘搜索' then v.visitor_count else 0 end) AS search_visitor_count,
SUM(case when v.src_type = '直通车' then v.visitor_count else 0 end) AS ztc_visitor_count,
SUM(v.visitor_count) AS zj_visitor_count,
SUM(case when v.src_type = '手淘搜索' then v.cart_addition_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '手淘搜索' then v.visitor_count else 0 end) AS search_cart_addition_rate,
SUM(case when v.src_type = '直通车' then v.cart_addition_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '直通车' then v.visitor_count else 0 end) AS ztc_cart_addition_rate,
SUM(v.cart_addition_rate*v.visitor_count)/SUM(v.visitor_count) AS zj_cart_addition_rate,
SUM(case when v.src_type = '手淘搜索' then v.conversion_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '手淘搜索' then v.visitor_count else 0 end) AS search_conversion_rate,
SUM(case when v.src_type = '直通车' then v.conversion_rate*v.visitor_count else 0 end)/SUM(case when v.src_type = '直通车' then v.visitor_count else 0 end) AS ztc_conversion_rate,   
SUM(v.conversion_rate*v.visitor_count)/SUM(v.visitor_count) AS zj_conversion_rate, 
SUM(case when v.src_type = '手淘搜索' then v.fan_payment_buyer_count else 0 end) AS search_fans_paid_buyers_count,
SUM(case when v.src_type = '直通车' then v.fan_payment_buyer_count else 0 end) AS ztc_fans_paid_buyers_count,
SUM(v.fan_payment_buyer_count) AS zj_fans_paid_buyers_count,
SUM(case when v.src_type = '手淘搜索' then v.direct_payment_buyer_count else 0 end) AS search_direct_paid_buyers_count,
SUM(case when v.src_type = '直通车' then v.direct_payment_buyer_count else 0 end) AS ztc_direct_paid_buyers_count,
SUM(v.direct_payment_buyer_count) AS zj_direct_paid_buyers_count
FROM
biz_product_keyword v
WHERE 
v.product_id = :productid and v.statistic_date >=:startdate and v.statistic_date <= :enddate  
group by v.keyword

-- LIMIT :offset , :pageSize  

`
