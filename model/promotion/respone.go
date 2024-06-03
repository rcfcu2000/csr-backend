package promotion

//返回推广分析页面所有数据
type RespPromotionAllData struct {
	PromotionIndex1 RespPromotionIndex1 `json:"promotionIndex1"` //重要指标1
	PromotionIndex2 RespPromotionIndex2 `json:"promotionIndex2"` //重要指标1
	BidTypeAnalysis RespBidTypeAnalysis `json:"bidTypeAnalysis"` //货盘花费
	PalletCost      RespPalletCost      `json:"palletCost"`      //货盘花费
	KeywordCost     RespKeywordCost     `json:"keywordCost"`     //关键词花费
	CrowdSpend      RespCrowdSpend      `json:"crowdSpend"`      //人群花费
	// ProductAnalysis RespProductAnalysis `json:"productAnalysis"` //商品分析
	// PlanAnalysis    RespPlanAnalysis    `json:"planAnalysis"`    //计划分析
}

//4个图表信息
type RespPromotionSearchData struct {
	BidTypeAnalysis RespBidTypeAnalysis `json:"bidTypeAnalysis"` //货盘花费
	PalletCost      RespPalletCost      `json:"palletCost"`      //货盘花费
	KeywordCost     RespKeywordCost     `json:"keywordCost"`     //关键词花费
	CrowdSpend      RespCrowdSpend      `json:"crowdSpend"`      //人群花费
}

//推广效果指标 重要指标1
type RespPromotionIndex1 struct {
	// 推广花费
	PromotionCost float64 `json:"promotion_cost"`

	// 推广产生的GMV（商品交易总额）
	PromotionGMV float64 `json:"promotion_gmv"`

	// 全店GMV
	OverallGMV float64 `json:"overall_gmv"`

	// 花费占比（推广花费占全店花费的比例）
	CostPercentage float64 `json:"cost_percentage"`

	// 推广GMV占比（推广GMV占全店GMV的比例）
	PromotionGMVPercentage float64 `json:"promotion_gmv_percentage"`

	// 推广流量占比（推广带来的流量占全店流量的比例）
	PromotionTrafficPercentage float64 `json:"promotion_traffic_percentage"` // 推广流量占比

	// 推广加购率（推广带来的加购次数占推广流量的比例）
	PromotionAddToCartRate float64 `json:"promotion_add_to_cart_rate"`

	// 全店加购率（全店加购次数占全店流量的比例）
	OverallAddToCartRate float64 `json:"overall_add_to_cart_rate"`

	// 推广转化率（推广带来订单的转化率）
	PromotionConversionRate float64 `json:"promotion_conversion_rate"`

	// 全店转化率（全店订单的总体转化率）
	OverallConversionRate float64 `json:"overall_conversion_rate"`

	// 推广投资回报率（ROI）
	PromotionROI float64 `json:"promotion_roi"`

	// 全店投资回报率（ROI）
	OverallROI float64 `json:"overall_roi"`
}

//推广效果指标  重要指标2
type PromotionIndex2Node struct {
	// 场景分类
	SceneCategory string `json:"scene_category"`

	// 花费
	Spend float64 `json:"spend"`

	// 成交成本（平均每次成交所需花费）
	TransactionCost float64 `json:"transaction_cost"`

	// GMV（商品交易总额）
	GMV float64 `json:"gmv"`

	// 渠道占比（此推广渠道在总渠道中的占比）
	ChannelPercentage float64 `json:"channel_percentage"`

	// 推广投资回报率（ROI）
	PromotionROI float64 `json:"promotion_roi"`

	// 点击量
	Clicks int64 `json:"clicks"`

	// 点击率（点击量占展现量的比例）
	ClickThroughRate float64 `json:"click_through_rate"`

	// 平均每次点击成本（CPC）
	CPC float64 `json:"cpc"`

	// 加购率（用户添加购物车次数占点击量的比例）
	AddToCartRate float64 `json:"add_to_cart_rate"`

	// 加购成本（平均每产生一次加购行为所需的花费）
	AddToCartCost float64 `json:"add_to_cart_cost"`

	// 旺旺咨询量
	AliWangWangInquiries int64 `json:"ali_wang_wang_inquiries"`

	//////////////////////////////////
	// 不显示，用到

	GuidedVisits   int64 `json:"guided_visits"`   // 访问数
	GuidedVisitors int64 `json:"guided_visitors"` // 访问客数
	ShopcartCount  int64 `json:"shopcart_count"`  //总购物车数
	BuyerCount     int64 `json:"buyer_count"`     //买家数
	GmvCount       int64 `json:"gmv_count"`       //成交笔数

}

type RespPromotionIndex2 struct {
	Records []PromotionIndex2Node `json:"records"`
}

// 出价类型分析
type BidTypeAnalysisNode struct {
	// 出价类型
	BidType string `json:"bid_type"`

	// 花费
	Spend float64 `json:"spend"`

	// GMV（商品交易总额）
	GMV float64 `json:"gmv"`

	// GMV占比（此次广告活动产生的GMV占总体GMV的比例）
	GMVPercentage float64 `json:"gmv_percentage"`

	// 点击量
	Clicks int64 `json:"clicks"`

	// 点击率（点击量占展现量的比例）
	ClickThroughRate float64 `json:"click_through_rate"`

	// 加购成本（每产生一次加购行为的平均成本）
	AddToCartCost float64 `json:"add_to_cart_cost"`

	// 推广投资回报率（ROI）
	PromotionROI float64 `json:"promotion_roi"`

	// 成交成本（平均每次成交所需花费）
	TransactionCost float64 `json:"transaction_cost"`

	// 推广转化率（从点击到最终成交的转化比例）
	ConversionRate float64 `json:"conversion_rate"`
}

type RespBidTypeAnalysis struct {
	Records []BidTypeAnalysisNode `json:"records"`
}

//  货盘花费
type PalletCostNode struct {
	Pallet string `json:"pallet"` // 当期货盘

	// 花费字段，记录广告推广过程中消耗的总费用
	Cost     float64 `json:"cost"`      // 广告花费总额，单位通常为元或其他货币单位
	CostRate float64 `json:"cost_rate"` // 广告花费总额占比

	// GMV字段，记录广告推广带来的总交易额
	GMV     float64 `json:"gmv"`      // 广告推广所促成的总成交金额
	GMVRate float64 `json:"gmv_rate"` // gmv占比

	// ROI字段，即投资回报率，衡量广告投入产出效益
	ROI float64 `json:"roi"` // ROI = （GMV - 花费） / 花费
	//产品个数
	ProductNum float64 `json:"product_num"`
}

type RespPalletCost struct {
	Records []PalletCostNode `json:"records"`
}

// 5.	关键词花费（TOP20）
type KeywordCostNode struct {
	Keyword string `json:"keyword"` //关键词
	// 花费字段，记录单个广告或一组广告的具体消费金额
	Cost float64 `json:"cost"` // 广告花费总额，单位通常为元或其他货币单位
	// 花费占比字段，表示该广告花费占整体预算的比例
	CostPercentage float64 `json:"cost_percentage"` // 花费占比，计算公式为：当前广告花费 / 总预算 * 100%
	// ROI字段，表示广告投资回报率，衡量广告效益
	ROI float64 `json:"roi"` // ROI = （总收入或总价值 - 花费） / 花费

}

type RespKeywordCost struct {
	Records []KeywordCostNode `json:"records"`
}

// 6.	人群花费（TOP20）
type CrowdSpendNode struct {
	Crowd string `json:"crowd"` // 人群

	// 花费字段，记录针对特定人群的广告具体消费金额
	Spend float64 `json:"spend"` // 针对目标人群的广告花费总额，单位通常为元或其他货币单位

	// 花费占比字段，表示该人群对应的广告花费占总体花费的比例
	SpendPercentage float64 `json:"spend_percentage"` // 目标人群花费占比，计算公式为：该人群花费 / 总广告花费 * 100%

	// ROI字段，表示针对该人群的广告投资回报率
	ROI float64 `json:"roi"` // ROI = （该人群带来的总收入或总价值 - 对该人群的花费） / 对该人群的花费

}

type RespCrowdSpend struct {
	Records []CrowdSpendNode `json:"records"`
}

// 7.	商品分析
type ProductAnalysisNode struct {
	// 商品编号
	ProductId string `json:"product_id"`

	// 当期货盘情况
	Pallet string `json:"pallet"`

	// 商品名称
	ProductName string `json:"product_name"`

	ProductAlias string `json:"product_alias"` //商品简称

	// 花费
	Cost float64 `json:"cost"`

	// 花费占比（该商品的花费占总花费的比例）
	CostPercentage float64 `json:"cost_percentage"`

	// 花费趋势（可能是时间段内的花费变化趋势数据）
	// CostTrend interface{} `json:"cost_trend"` // 这里使用interface{}，实际类型取决于具体的数据结构（如数组或对象）

	// GMV（商品交易总额）
	GMV float64 `json:"gmv"`

	// GMV占比（该商品的GMV占总GMV的比例）
	GMVPercentage float64 `json:"gmv_percentage"`

	// GMV趋势（同上，可能为GMV的变化趋势数据）
	// GMVTrend interface{} `json:"gmv_trend"`

	// ROI（投资回报率）
	ROI float64 `json:"roi"`

	// ROI趋势（ROI随时间的变化趋势）
	// ROITrend interface{} `json:"roi_trend"`

	// 直接ROI
	DirectROI float64 `json:"direct_roi"`

	// 间接ROI
	IndirectROI float64 `json:"indirect_roi"`

	// 新客占比（新客户带来的交易占比）
	NewCustomerPercentage float64 `json:"new_customer_percentage"`

	// 老客占比（老客户带来的交易占比）
	ExistingCustomerPercentage float64 `json:"existing_customer_percentage"`

	// 点击量
	Clicks int64 `json:"clicks"`

	// 点击率
	ClickThroughRate float64 `json:"click_through_rate"`

	// 转化率（从点击到购买的转化比例）
	ConversionRate float64 `json:"conversion_rate"`

	// 客单价（每位顾客平均购买金额）
	AverageOrderValue float64 `json:"average_order_value"`

	// 直接成交金额
	DirectTransactionAmount float64 `json:"direct_transaction_amount"`

	// 直接成交笔数
	DirectTransactionCount int64 `json:"direct_transaction_count"`

	// 间接成交金额
	IndirectTransactionAmount float64 `json:"indirect_transaction_amount"`

	// 间接成交笔数
	IndirectTransactionCount int64 `json:"indirect_transaction_count"`

	// 收藏加购率
	FavoriteAddToCartRate float64 `json:"favorite_add_to_cart_rate"`

	// 三级类目
	LevelThreeCategory string `json:"level_three_category"`
}

// 商品分析返回
type RespProductAnalysis struct {
	Count   int64                 `json:"count"`
	Records []ProductAnalysisNode `json:"records"`
	Sum     []ProductAnalysisNode `json:"sum"`
}

///////////////////////////////////////////////
//商品趋势数据
type ProductTrendNode struct {
	// 商品编号
	ProductId string `json:"product_id"`
	// 日期
	Date string `json:"date"`
	// 花费趋势（可能是时间段内的花费变化趋势数据）
	SpendTrend float64 `json:"spend_trend"` // 这里使用interface{}，实际类型取决于具体的数据结构（如数组或对象）
	// GMV趋势（GMV随时间的变化趋势）
	GMVTrend float64 `json:"gmv_trend"`
	// ROI趋势（ROI随时间的变化趋势）
	ROITrend float64 `json:"roi_trend"`
}

//产品趋势返回
type ProductOneTrend struct {
	// 商品编号
	ProductId string             `json:"product_id"`
	Records   []ProductTrendNode `json:"records"`
}

//产品趋势返回
type RespProductTrend struct {
	Records []ProductOneTrend `json:"records"`
	Sum     ProductOneTrend   `json:"sum"`
}

///////////////////////////////////////////////
// 8.计划数据
type PlanAnalysisNode struct {
	// 计划编号
	PlanId string `json:"plan_id"`

	// 出价类型
	BidType string `json:"bid_type"`

	// 计划名称
	CampaignName string `json:"campaign_name"`

	// 花费
	Spend float64 `json:"spend"`

	// 花费趋势（可能是时间段内的花费变化趋势数据）
	// SpendTrend interface{} `json:"spend_trend"` // 这里使用interface{}，实际类型取决于具体的数据结构（如数组或对象）

	// GMV（商品交易总额）
	GMV float64 `json:"gmv"`

	// GMV趋势（GMV随时间的变化趋势）
	// GMVTrend interface{} `json:"gmv_trend"`

	// ROI（投资回报率）
	ROI float64 `json:"roi"`

	// ROI趋势（ROI随时间的变化趋势）
	// ROITrend interface{} `json:"roi_trend"`

	// 直接ROI
	DirectROI float64 `json:"direct_roi"`

	// 间接ROI
	IndirectROI float64 `json:"indirect_roi"`

	// 点击量
	Clicks int64 `json:"clicks"`

	// 点击率
	ClickThroughRate float64 `json:"click_through_rate"`

	// 每次点击成本（Cost Per Click）
	CPC float64 `json:"cpc"`

	// 转化率（从点击到购买的转化比例）
	ConversionRate float64 `json:"conversion_rate"`

	// 加购成本（每增加一个购物车的成本）
	AddToCartCost float64 `json:"add_to_cart_cost,omitempty"` // （可选字段，如果数据不适用则不用）

	// 成交成本（每次成交的平均成本）
	TransactionCost float64 `json:"transaction_cost,omitempty"` // （可选字段，如果数据不适用则不用）

	// 直接成交金额
	DirectTransactionAmount float64 `json:"direct_transaction_amount"`

	// 直接成交笔数
	DirectTransactionCount int64 `json:"direct_transaction_count"`

	// 间接成交金额
	IndirectTransactionAmount float64 `json:"indirect_transaction_amount"`

	// 间接成交笔数
	IndirectTransactionCount int64 `json:"indirect_transaction_count"`

	// 计划场景
	PromotionType string `json:"promotion_type"`
}

// 计划分析返回
type RespPlanAnalysis struct {
	Count   int64              `json:"count"`
	Records []PlanAnalysisNode `json:"records"`
	Sum     []PlanAnalysisNode `json:"sum"`
}

///////////////////////////////////////////////

//计划趋势数据
type PlanTrendNode struct {
	// 计划编号
	PlanId string `json:"plan_id"`
	// 日期
	Date string `json:"date"`
	// 花费趋势（可能是时间段内的花费变化趋势数据）
	SpendTrend float64 `json:"spend_trend"`
	// GMV趋势（GMV随时间的变化趋势）
	GMVTrend float64 `json:"gmv_trend"`
	// ROI趋势（ROI随时间的变化趋势）
	ROITrend float64 `json:"roi_trend"`
}

//计划趋势数据
type PlanOneTrend struct {
	PlanId  string          `json:"plan_id"`
	Records []PlanTrendNode `json:"records"`
}

//计划趋势返回
type RespPlanTrend struct {
	Records []PlanOneTrend `json:"records"`
	Sum     PlanOneTrend   `json:"sum"`
}
