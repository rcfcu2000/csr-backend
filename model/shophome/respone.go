package shophome

// 返回重要指标数据
type RespShopHomeIndexData struct {
	// 访客数
	Visitors float64 `json:"visitors"`
	// GMV
	Gmv float64 `json:"gmv"`
	// 客单价
	CustomerUnitPrice float64 `json:"customer_unit_price"` // // 客单价
	// 支付转化率
	ConversionRatePayment float64 `json:"conversion_rate_payment"` // 支付转化率，即实际支付买家数与访客数的比例。
	// 推广花费
	Spend float64 `json:"spend"`
	// 推广占比
	PromotionPercentage float64 `json:"promotion_percentage"`
	// 净利润
	Profit float64 `json:"profit"`
	// 净利润率
	ProfitRate float64 `json:"profit_rate"`
}

// 返回流量渠道数据
type RespShopHomeTrafficData struct {
	Records []TrafficL3Node `json:"records"`
}

type TrafficL3Node struct {
	// 三级来源
	L3 string `json:"l3"`
	// 访客数
	Visitors float64 `json:"visitors"`
	// GMV
	Gmv float64 `json:"gmv"`
}

// 返回Gmv访客数据
type RespShopHomeGmvVisitorsData struct {
	Records []GmvTargetNode `json:"records"`
	Sum     GmvTargetNode   `json:"sum"`
}

type GmvTargetNode struct {
	// 货盘
	Pallet string `json:"pallet"`
	// GMV - 累计GMV
	Gmv float64 `json:"gmv"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
	// 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`
}

// 返回汇总趋势数据
type RespShopHomeSumTrendData struct {
	Records []SumTrendNode `json:"records"`
}

type SumTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 访客数
	Visitors float64 `json:"visitors"`
	// GMV
	Gmv float64 `json:"gmv"`
}

// 返回关键词对比数据
type RespShopHomeKeywordData struct {
	ZtcRecords    []KeywordNode `json:"ztc_records"`
	SearchRecords []KeywordNode `json:"search_records"`
}

type KeywordNode struct {
	// 搜索词来源  直通车 或 搜索
	Src string `json:"src"`
	// 关键词
	Keyword string `json:"keyword"`
	// 访客数
	Visitors float64 `json:"visitors"`
}

// 返回推广分析数据
type RespShopHomePromotionData struct {
	Records []PromotionNode `json:"records"`
}

type PromotionNode struct {

	// 场景分类
	Scene string `json:"scene"`
	// 推广花费
	spend string `json:"spend"`
	// 成交成本
	TransactionCost float64 `json:"transaction_cost"`
	// GMV
	Gmv string `json:"gmv"`
	// GMV占比
	GmvPercentage string `json:"scene_percentage"`
	// 推广ROI
	Roi string `json:"roi"`
	// 点击量
	Clicks string `json:"clicks"`
	// 点击率
	ClicksRate string `json:"clicks_rate"`
	// CPC
	CPC string `json:"cpc"`
	// 加购率
	AddToCartRate float64 `json:"add_to_cart_rate"`
	// 加购成本
	AddToCartCost float64 `json:"add_to_cart_cost"`
	// 旺旺咨询量
	AliWangWangInquiries int64 `json:"ali_wang_wang_inquiries"`
}

// 返回店铺数据
type RespShopHomeExperienceScoreData struct {
	Records []ExperienceScoreNode `json:"records"`
}

type ExperienceScoreNode struct {

	// 日期
	Date string `json:"date"`
	// 综合体验分
	OverallExperienceScore string `json:"overall_experience_score"`
	// 商品体验
	ProductExperienceScore string `json:"product_experience_score"`
	// 物流体验
	LogisticsExperienceScore string `json:"logistics_experience_score"`
	// 服务体验
	Service_ExperienceScore string `json:"service_experience_score"`
}
