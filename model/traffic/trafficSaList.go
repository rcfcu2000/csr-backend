package traffic

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespTrafficListData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_List)
	if len(r.ProductId) != 0 {
		sqlp.SetSql(SQL_TRAFFIC_P_List)
	}
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_TRAFFIC_List sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_TRAFFIC_List sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespTrafficListData) Prodata() {

}

const SQL_TRAFFIC_List = `
with shop as 

(

select 

tertiary_source as source_type_3, -- 三级来源

sum(visitors_count) as shop_visitors_count, -- 访客数

sum(buyers_paid) as shop_paid_buyers, -- 支付买家数

sum(add_to_cart_count)/sum(visitors_count) as shop_add_car_rate,-- 加购率

sum(add_to_cart_pay_buyers_count)/sum(add_to_cart_count) as shop_add_cart_conversion_rate --  加购转化率

from biz_shop_traffic s

where tertiary_source <> '汇总'

and date >=:startdate  and date <=:enddate  

-- and date >=:startdate and date <= :enddate 

-- and src = '每一次访问来源'  流量归属原则参数

group by tertiary_source

),

shop_all as (

select 

sum(shop_visitors_count) as all_shop_visitors_count, -- 访客数

sum(shop_paid_buyers) as all_shop_paid_buyers -- 支付买家数

from shop 

),

pro as (

select 

tertiary_source as source_type_3, -- 三级来源

sum(visitors_count) as visitors_count, -- 访客数

-- 访客TGI visitor_tgi

-- 购买TGI buy_tgi

sum(paid_amount)/sum(buyers_paid) as customer_unit_price, -- 客单价

sum(buyers_paid) as paid_buyers, -- 支付买家数

sum(buyers_paid)/sum(visitors_count) as pay_conversion_rate, -- 支付转化率

sum(add_to_cart_count)/sum(visitors_count) as add_car_rate,-- 加购率

sum(add_to_cart_count) as add_to_cart_buyers, -- 加购人数

sum(add_to_cart_pay_buyers_count)/sum(add_to_cart_count) as add_cart_conversion_rate, --  加购转化率

sum(paid_amount)/sum(visitors_count) as uv

from biz_shop_traffic p

where tertiary_source <> '汇总'

and date >=:startdate  and date <=:enddate  

and '全店' = '全店'

-- and date >=:startdate and date <= :enddate 

-- and src = '每一次访问来源'  流量归属原则参数

-- and :productid     = '全店'

group by tertiary_source

union all 

select 

p.source_type_3 as source_type_3, -- 三级来源

sum(p.visitors_count) as visitors_count, -- 访客数

-- 访客TGI visitor_tgi

-- 购买TGI buy_tgi

sum(p.paid_amount)/sum(p.total_paid_buyers) as customer_unit_price, -- 客单价

sum(p.total_paid_buyers) as paid_buyers, -- 支付买家数

sum(p.total_paid_buyers)/sum(p.visitors_count) as pay_conversion_rate, -- 支付转化率

sum(p.add_to_carts)/sum(p.visitors_count) as add_car_rate,-- 加购率

sum(p.add_to_carts) as add_to_cart_buyers, -- 加购人数

sum(p.add_to_cart_and_paid_buyers)/sum(p.add_to_carts) as add_cart_conversion_rate, --  加购转化率

sum(p.paid_amount)/sum(p.visitors_count) as uv

from biz_product_traffic_stats p

left join biz_product bp  on p.product_id = bp.product_id

where  p.source_type_3  <> '汇总'

and p.statistic_date >=:startdate  and p.statistic_date <=:enddate  

-- and p.product_id = '全店'

-- and statistic_date >=:startdate and statistic_date <= :enddate 

-- and src = '每一次访问来源'  流量归属原则参数

-- and product_id  = :productid     

-- and bp.responsible IN :resperson  

group by source_type_3

),

pro_all as 

(

select 

sum(visitors_count) as all_visitors_count, -- 访客数

sum(paid_buyers) as all_paid_buyers -- 支付买家数

from pro

),

channel as (

select 

tertiary_source as source_type_3, -- 三级来源

sum(case when src ='第一次访问来源' then buyers_paid else 0 end)/sum(case when src ='每一次访问来源' then buyers_paid else 0 end) as zc,

sum(case when src ='最后一次访问来源' then buyers_paid else 0 end)/sum(case when src ='每一次访问来源' then buyers_paid else 0 end) as sg,

sum(buyers_paid) as czc

from biz_shop_traffic p

where tertiary_source <> '汇总'

and date >=:startdate  and date <=:enddate  

and '全店' = '全店'

-- and date >=:startdate and date <= :enddate 

-- and :productid     = '全店'

group by tertiary_source

union all 

select 

p.source_type_3 as source_type_3, -- 三级来源

sum(case when src ='第一次访问来源' then p.total_paid_buyers else 0 end)/sum(case when src ='每一次访问来源' then p.total_paid_buyers else 0 end) as zc,

sum(case when src ='最后一次访问来源' then p.total_paid_buyers else 0 end)/sum(case when src ='每一次访问来源' then p.total_paid_buyers else 0 end) as sg,

sum(p.total_paid_buyers) as czc

from biz_product_traffic_stats p

left join biz_product bp  on p.product_id = bp.product_id

where  p.source_type_3  <> '汇总'

and p.statistic_date >=:startdate  and p.statistic_date <=:enddate  

-- and p.product_id = '全店'

-- and statistic_date >=:startdate and statistic_date <= :enddate 

-- and product_id  = :productid     

-- and bp.responsible IN :resperson  

group by source_type_3

)



select 

pro.source_type_3, -- 三级来源

pro.visitors_count, -- 访客数

(pro.visitors_count/pro_all.all_visitors_count)/(shop.shop_visitors_count/shop_all.all_shop_visitors_count) as visitor_tgi,-- 访客TGI 

(pro.paid_buyers/pro_all.all_paid_buyers)/(shop.shop_paid_buyers/shop_all.all_shop_paid_buyers) as buy_tgi,-- 购买TGI 

case when channel.czc = 0 then '纯种草' when channel.zc<channel.sg then '偏收割' else '偏种草' end as channel_attribute,

channel.zc-channel.sg as channel_diff,

pro.customer_unit_price, -- 客单价

pro.paid_buyers, -- 支付买家数

pro.pay_conversion_rate, -- 支付转化率

pro.add_car_rate,-- 加购率

pro.add_to_cart_buyers, -- 加购人数

pro.add_cart_conversion_rate, --  加购转化率

shop.shop_add_car_rate,

shop.shop_add_cart_conversion_rate,

pro.uv,

pro.visitors_count/pro_all.all_visitors_count as product_visitor_percentage,

shop.shop_visitors_count/shop_all.all_shop_visitors_count as shop_visitor_percentage,

pro.paid_buyers/pro_all.all_paid_buyers as product_buyer_percentage,

shop.shop_paid_buyers/shop_all.all_shop_paid_buyers as shop_buyer_percentage

from pro

left join shop on pro.source_type_3 = shop.source_type_3

left join channel on pro.source_type_3 = channel.source_type_3

left join pro_all on 1=1

left join shop_all on 1=1

LIMIT :offset , :pageSize  

 
`

const SQL_TRAFFIC_P_List = `
with shop as 

(

select 

tertiary_source as source_type_3, -- 三级来源

sum(visitors_count) as shop_visitors_count, -- 访客数

sum(buyers_paid) as shop_paid_buyers, -- 支付买家数

sum(add_to_cart_count)/sum(visitors_count) as shop_add_car_rate,-- 加购率

sum(add_to_cart_pay_buyers_count)/sum(add_to_cart_count) as shop_add_cart_conversion_rate --  加购转化率

from biz_shop_traffic s

where tertiary_source <> '汇总'

and date >=:startdate  and date <=:enddate  

-- and date >=:startdate and date <= :enddate 

-- and src = '每一次访问来源'  流量归属原则参数

group by tertiary_source

),

shop_all as (

select 

sum(shop_visitors_count) as all_shop_visitors_count, -- 访客数

sum(shop_paid_buyers) as all_shop_paid_buyers -- 支付买家数

from shop 

),

pro as (

select 

tertiary_source as source_type_3, -- 三级来源

sum(visitors_count) as visitors_count, -- 访客数

-- 访客TGI visitor_tgi

-- 购买TGI buy_tgi

sum(paid_amount)/sum(buyers_paid) as customer_unit_price, -- 客单价

sum(buyers_paid) as paid_buyers, -- 支付买家数

sum(buyers_paid)/sum(visitors_count) as pay_conversion_rate, -- 支付转化率

sum(add_to_cart_count)/sum(visitors_count) as add_car_rate,-- 加购率

sum(add_to_cart_count) as add_to_cart_buyers, -- 加购人数

sum(add_to_cart_pay_buyers_count)/sum(add_to_cart_count) as add_cart_conversion_rate, --  加购转化率

sum(paid_amount)/sum(visitors_count) as uv

from biz_shop_traffic p

where tertiary_source <> '汇总'

and date >=:startdate  and date <=:enddate  

and '全店' = '全店'


-- and date >=:startdate and date <= :enddate 

-- and src = '每一次访问来源'  流量归属原则参数

-- and :productid     = '全店'

group by tertiary_source

union all 

select 

p.source_type_3 as source_type_3, -- 三级来源

sum(p.visitors_count) as visitors_count, -- 访客数

-- 访客TGI visitor_tgi

-- 购买TGI buy_tgi

sum(p.paid_amount)/sum(p.total_paid_buyers) as customer_unit_price, -- 客单价

sum(p.total_paid_buyers) as paid_buyers, -- 支付买家数

sum(p.total_paid_buyers)/sum(p.visitors_count) as pay_conversion_rate, -- 支付转化率

sum(p.add_to_carts)/sum(p.visitors_count) as add_car_rate,-- 加购率

sum(p.add_to_carts) as add_to_cart_buyers, -- 加购人数

sum(p.add_to_cart_and_paid_buyers)/sum(p.add_to_carts) as add_cart_conversion_rate, --  加购转化率

sum(p.paid_amount)/sum(p.visitors_count) as uv

from biz_product_traffic_stats p

left join biz_product bp  on p.product_id = bp.product_id

where  p.source_type_3  <> '汇总'

and p.statistic_date >=:startdate  and p.statistic_date <=:enddate  

and p.product_id = :productid 

-- and statistic_date >=:startdate and statistic_date <= :enddate 

-- and src = '每一次访问来源'  流量归属原则参数

-- and product_id  = :productid     

-- and bp.responsible IN :resperson  

group by source_type_3

),

pro_all as 

(

select 

sum(visitors_count) as all_visitors_count, -- 访客数

sum(paid_buyers) as all_paid_buyers -- 支付买家数

from pro

),

channel as (

select 

tertiary_source as source_type_3, -- 三级来源

sum(case when src ='第一次访问来源' then buyers_paid else 0 end)/sum(case when src ='每一次访问来源' then buyers_paid else 0 end) as zc,

sum(case when src ='最后一次访问来源' then buyers_paid else 0 end)/sum(case when src ='每一次访问来源' then buyers_paid else 0 end) as sg,

sum(buyers_paid) as czc

from biz_shop_traffic p

where tertiary_source <> '汇总'

and date >=:startdate  and date <=:enddate  

and '全店' = '全店'



-- and date >=:startdate and date <= :enddate 

-- and :productid     = '全店'

group by tertiary_source

union all 

select 

p.source_type_3 as source_type_3, -- 三级来源

sum(case when src ='第一次访问来源' then p.total_paid_buyers else 0 end)/sum(case when src ='每一次访问来源' then p.total_paid_buyers else 0 end) as zc,

sum(case when src ='最后一次访问来源' then p.total_paid_buyers else 0 end)/sum(case when src ='每一次访问来源' then p.total_paid_buyers else 0 end) as sg,

sum(p.total_paid_buyers) as czc

from biz_product_traffic_stats p

left join biz_product bp  on p.product_id = bp.product_id

where  p.source_type_3  <> '汇总'

and p.statistic_date >=:startdate  and p.statistic_date <=:enddate  

-- and p.product_id = '全店'

and p.product_id = :productid 

-- and statistic_date >=:startdate and statistic_date <= :enddate 

-- and product_id  = :productid     

-- and bp.responsible IN :resperson  

group by source_type_3

)



select 

pro.source_type_3, -- 三级来源

pro.visitors_count, -- 访客数

(pro.visitors_count/pro_all.all_visitors_count)/(shop.shop_visitors_count/shop_all.all_shop_visitors_count) as visitor_tgi,-- 访客TGI 

(pro.paid_buyers/pro_all.all_paid_buyers)/(shop.shop_paid_buyers/shop_all.all_shop_paid_buyers) as buy_tgi,-- 购买TGI 

case when channel.czc = 0 then '纯种草' when channel.zc<channel.sg then '偏收割' else '偏种草' end as channel_attribute,

channel.zc-channel.sg as channel_diff,

pro.customer_unit_price, -- 客单价

pro.paid_buyers, -- 支付买家数

pro.pay_conversion_rate, -- 支付转化率

pro.add_car_rate,-- 加购率

pro.add_to_cart_buyers, -- 加购人数

pro.add_cart_conversion_rate, --  加购转化率

shop.shop_add_car_rate,

shop.shop_add_cart_conversion_rate,

pro.uv,

pro.visitors_count/pro_all.all_visitors_count as product_visitor_percentage,

shop.shop_visitors_count/shop_all.all_shop_visitors_count as shop_visitor_percentage,

pro.paid_buyers/pro_all.all_paid_buyers as product_buyer_percentage,

shop.shop_paid_buyers/shop_all.all_shop_paid_buyers as shop_buyer_percentage

from pro

left join shop on pro.source_type_3 = shop.source_type_3

left join channel on pro.source_type_3 = channel.source_type_3

left join pro_all on 1=1

left join shop_all on 1=1

LIMIT :offset , :pageSize  

 
`
