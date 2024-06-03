package product

import (
	"fmt"
	"strings"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductIndex) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_Index)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespProductIndex sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductIndex sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespProductIndex) Prodata() {

}

// 商品访客数	指标	商品每日数据	商品访客数
// GMV	指标	商品每日数据	支付金额
// 支付转化率	指标	商品每日数据	支付转化率=支付买家数/访客数
// 搜索访客占比	指标	商品流量数据	取来源类型=手淘搜索的访客数
// 搜索GMV占比	指标	商品每日数据
// 商品流量数据	分子取商品流量数据，来源效果=每一次，二级来源=手淘搜索，三级来源=汇总的支付金额；分母取商品每日数据的支付金额
// 老客占比	指标	商品每日数据	老买家人数占比=支付老买家数/支付买家数
// 退款率	指标	商品每日数据	退款率=成功退款金额/支付金额

const SQL_PRODUCT_Index = `

WITH t1 AS (
	select v.product_id as "product_id",
	sum(v.visitors_count) as "visitors_count", -- 商品访客数
	sum(v.paid_amount) as "gmv",  -- gmv
	sum(v.paid_buyers) / sum(v.visitors_count) as "payment_conversion_rate",  -- 订单支付转化率
	sum(v.returning_buyers_paid) / sum(v.paid_buyers) as "returning_customer_ratio",  -- 老客户占比
 
	CASE
	WHEN  sum(bpt_gmv) > 0 THEN sum(bpt_search_gmv) / sum(v.paid_amount)
	ELSE 0
	END AS search_gmv_ratio, -- 搜索渠道贡献的GMV占比
	
	CASE
	WHEN sum(bpt_visitors_count) > 0 THEN sum(bpt_search_count) / sum(v.visitors_count)
	ELSE 0
	END AS search_visitor_ratio,  -- 搜索渠道访问商品的访客比例	

	sum(v.successful_refund_amount) / sum(v.paid_amount) as "refund_rate"  -- 退款率
 
	from v_product_day_all  v -- v_bpp_bp =>v_product_day_all
	where (v.product_id =:productid )
	and ( v.statistic_date >=:startdate  and v.statistic_date <=:enddate   )
	 
)
	
	select  
	bpdm.product_id as "product_id",
	associated_leaf_category AS "bundle_purchase", -- 连带率
	repurchase_rate AS "repeat_purchase_rate", -- 复购率
	free_click_rate AS "free_search_click_rate",   -- 免费搜索点击率
	t1.*
	from biz_product_damo bpdm 
	left join t1 on bpdm.product_id = t1.product_id
	where bpdm.product_id =:productid   and statistic_date = :enddate  

`

// select v.product_id as "product_id",
// sum(v.visitors_count) as "visitors_count", -- 商品访客数
// sum(v.paid_amount) as "gmv",  -- gmv
// sum(v.paid_buyers) / sum(v.visitors_count) as "payment_conversion_rate",  -- 订单支付转化率
// sum(v.returning_buyers_paid) / sum(v.paid_buyers) as "returning_customer_ratio",  -- 老客户占比

// CASE
// WHEN  sum(bpt_gmv) > 0 THEN sum(bpt_search_gmv) / sum(bpt_gmv)
// ELSE 0
// END AS search_gmv_ratio, -- 搜索渠道贡献的GMV占比

// CASE
// WHEN sum(bpt_visitors_count) > 0 THEN sum(bpt_search_count) / sum(bpt_visitors_count)
// ELSE 0
// END AS search_visitor_ratio,  -- 搜索渠道访问商品的访客比例

// sum(v.successful_refund_amount) / sum(v.paid_amount) as "refund_rate",  -- 退款率

// sum(free_click_rate) as "free_search_click_rate",  -- 免费搜索点击率
// avg(associated_leaf_category) as "bundle_purchase",  -- 连带率
// avg(repurchase_rate) as "repeat_purchase_rate"    -- 复购率

// from v_product_day_all  v -- v_bpp_bp =>v_product_day_all
// where (v.product_id =:productid )
// and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
// order by  statistic_date

func (i *RespProductList) GetData(r ReqProductListSearch) error {

	// mainSql := fmt.Sprintf(SQL_PRODUCT_LIST, r.Key)
	mainSql := strings.ReplaceAll(SQL_PRODUCT_LIST, "%s", r.Key)
	fmt.Println("RespProductList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductList sql\n", mainSql)
		return e
	}
	// i.Prodata()
	return nil

}

const SQL_PRODUCT_LIST = `
select DISTINCT  product_id, product_name  from biz_product
where product_id LIKE "%%s%" or product_name LIKE "%%s%"

`
