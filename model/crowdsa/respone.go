package crowdsa

//返回人群Gmv数据列表
type RespCrowdGmvListData struct {
	Records []CrowdGmvNode `json:"records"`
}

// 人群Gmv数据
type CrowdGmvNode struct {
	// 日期
	Date string `json:"date"`
	// 人群类型
	CrowdType string `json:"crowd_type"`
	// gmv
	Gmv float64 `json:"gmv"`
	// 访客数
	VisitorsCount int64 `json:"visitors_count"`
	//人群TGI
	CrowdTGI float64 `json:"crowd_tgi"` // 人群TGI
	//客单价 支付金额/支付买家数
	CustomerUnitPrice float64 `json:"customer_unit_price"` // // 客单价
	// 支付转化率 - =支付买家数/访客数
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率 支付转化率 = 总成交笔数 / 点击量
	//人群加购率 =加购人数/访客数
	AddCarRate float64 `json:"add_car_rate"`
}

//返回人群Gmv趋势数据列表
type RespCrowdGmvTrendData struct {
	Records []CrowdGmvNode `json:"records"`
}

//商品Crowd分类
type RespProductCrowd10List struct {
	Records []ProductCrowd10Node `json:"records"`
}

// 产品你人群10数据
type ProductCrowd10Node struct {
	// 日期
	Date string `json:"date"`
	// 人群类型
	CrowdType string `json:"crowd_type"`
	// ProductId
	ProductId string `json:"product_id"`
	// ProductName
	ProductName string `json:"product_name"`
	// gmv
	Gmv float64 `json:"gmv"`
	// 支付买家数
	PaidBuyers int64 `json:"paid_buyers"`
	// 访客数
	VisitorsCount int64 `json:"visitors_count"`
	//人群TGI
	CrowdTGI float64 `json:"crowd_tgi"` // 人群TGI
	//客单价 支付金额/支付买家数
	CustomerUnitPrice float64 `json:"customer_unit_price"` // // 客单价
	// 支付转化率 - =支付买家数/访客数
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率 支付转化率 = 总成交笔数 / 点击量
}

//返回某一商品按人群分类
type RespProductCrowdsList struct {
	Records []ProductCrowd10Node `json:"records"`
}

const sql04 = `
select    crowd_type, , product_id, product_name,  sum(paid_amount) as gmv, sum(visitors) as visitors_count,
0 as crowd_tgi, sum(paid_amount) /sum(paid_buyers)  as customer_unit_price,
sum(paid_buyers)/sum(visitors) as payment_conversion_rate, 
sum(add_to_cart_count)/sum(visitors) as add_car_rate
from biz_shop_audience_product_t10 t 
WHERE t.year_month >="2024-01-01" and t.year_month <= LAST_DAY("2024-01-31" ) 
and product_id =:productid 
group by crowd_type, product_id, product_name


`

//返回某一商品按人群分类趋势数据
type RespProductCrowdsTrendList struct {
	Records []ProductCrowd10Node `json:"records"`
}

const sql05 = `
select   year_month as date ,  product_id, product_name,  sum(paid_amount) as gmv, sum(visitors) as visitors_count,
0 as crowd_tgi, sum(paid_amount) /sum(paid_buyers)  as customer_unit_price,
sum(paid_buyers)/sum(visitors) as payment_conversion_rate, 
sum(add_to_cart_count)/sum(visitors) as add_car_rate
from biz_shop_audience_product_t10 t 
WHERE t.year_month >="2024-01-01" and t.year_month <= LAST_DAY("2024-01-31" ) 
and product_id =:productid 
group by year_month, product_id, product_name   -- crowd_type, 
`

//返回商品流量来源
type RespProductSrcList struct {
	Records []ProductSrcNode `json:"records"`
}

// 商品流量来源数据
type ProductSrcNode struct {
	// 日期
	Date string `json:"date"`
	// 人群类型
	CrowdType string `json:"crowd_type"`
	// 二级来源
	SecondarySource string `json:"secondary_source"`
	// 三级来源
	TertiarySource string `json:"tertiary_source"`
	// ProductId
	ProductId string `json:"product_id"`
	// ProductName
	ProductName string `json:"product_name"`
	// gmv
	Gmv float64 `json:"gmv"`
	// 支付买家数
	PaidBuyers int64 `json:"paid_buyers"`
	// 访客数
	VisitorsCount int64 `json:"visitors_count"`
	//人群TGI
	CrowdTGI float64 `json:"crowd_tgi"` // 人群TGI
	//客单价 支付金额/支付买家数
	CustomerUnitPrice float64 `json:"customer_unit_price"` // // 客单价
	// 支付转化率 - =支付买家数/访客数
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率 支付转化率 = 总成交笔数 / 点击量
}

//返回人群流量来源
type RespCrowdSrcList struct {
	Records []ProductSrcNode `json:"records"`
}

// 商品流量来源数据
type CrowdSrcNode struct {
	// 日期
	Date string `json:"date"`
	// 人群类型
	CrowdType string `json:"crowd_type"`
	// 二级来源
	SecondarySource string `json:"secondary_source"`
	// 三级来源
	TertiarySource string `json:"tertiary_source"`
	// gmv
	Gmv float64 `json:"gmv"`
	// 支付买家数
	PaidBuyers int64 `json:"paid_buyers"`
	// 访客数
	VisitorsCount int64 `json:"visitors_count"`
	//人群TGI
	CrowdTGI float64 `json:"crowd_tgi"` // 人群TGI
	//客单价 支付金额/支付买家数
	CustomerUnitPrice float64 `json:"customer_unit_price"` // // 客单价
	// 支付转化率 - =支付买家数/访客数
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率 支付转化率 = 总成交笔数 / 点击量
}

//返回人群流量来源
type RespCrowdGmv20List struct {
	Records []CrowdGmv20Node `json:"records"`
}

// 商品流量来源数据
type CrowdGmv20Node struct {
	// 日期
	Date string `json:"date"`
	// 人群类型
	CrowdType string `json:"crowd_type"`
	// 二级来源
	SecondarySource string `json:"secondary_source"`
	// 三级来源
	TertiarySource string `json:"tertiary_source"`
	// gmv
	Gmv float64 `json:"gmv"`
}
