package shop

//返回店铺单页所有数据
type RespShopAllData struct {
	ShopIndex            ShopIndex            `json:"shop_index"`            //重要指标
	Content              Content              `json:"content"`               //内容
	ShopServiceAnalysis  ShopServiceAnalysis  `json:"shop_service_analysis"` //店铺服务分析
	CustomerService      CustomerService      `json:"customer_service"`      //客户服务
	CustomerAnalysis     CustomerAnalysis     `json:"customer_analysis"`     //客服分析
	CustomerLossAnalysis CustomerLossAnalysis `json:"customer_lossAnalysis"` //客户流失分析
}

//返回店铺单页所有趋势数据
type RespShopTrendData struct {
	ShopIndexTrend           ShopIndexTrendList            `json:"shop_index_trend"`            //重要指标
	ContentTrend             ContentTrendList              `json:"content_trend"`               //内容
	ShopServiceAnalysisTrend ShopServiceAnalysisTrendList  `json:"shop_service_analysis_trend"` //店铺服务分析
	CustomerServiceTrend     CustomerServiceTrendList      `json:"customer_service_trend"`      //客户服务分析
	CustomerAnalysisTrend    CustomerAnalysisTrendList     `json:"customer_analysis_trend"`     //客户分析
	CustomerLossTrend        CustomerLossAnalysisTrendList `json:"customer_loss_trend"`         //客户流失分析
}

//指标
type ShopIndex struct {
	// 动销率字段
	TurnoverRate float64 `json:"turnover_rate"`
	// 店铺层级字段
	Level int `json:"level"`
	// 店铺排行字段
	Ranking int `json:"ranking"`
}

//指标日趋势返回数据
type ShopIndexTrendList struct {
	Records []ShopIndexTrendNode `json:"records"`
}

//指标趋势
type ShopIndexTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 动销率字段
	TurnoverRate float64 `json:"turnover_rate"`
	// 店铺层级字段
	Level int `json:"level"`
	// 店铺排行字段
	Ranking int `json:"ranking"`
}

// 内容
type Content struct {
	Records []ContentTrendNode `json:"records"`
}

type ContentNode struct {
	// 内容类型字段，
	Type string `json:"type"` //内容类型
	// 种草成交金额字段，
	Amount float64 `json:"amount"` // 种草成交金额字段
}

//内容趋势返回数据
type ContentTrendList struct {
	Records []ContentTrendNode `json:"records"`
}

// 内容
type ContentTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 内容类型字段，
	Type string `json:"type"` //内容类型
	// 种草成交金额字段，
	Amount float64 `json:"amount"` // 种草成交金额字段
	// 占比字段，
	Proportion float64 `json:"proportion"`
}

// 店铺服务分析
type ShopServiceAnalysis struct {

	// 首次品退率，即第一次购买后发生的退货比率
	FirstProductReturnRate float64 `json:"first_product_return_rate"`

	// 首次品退率是否超过同行， 表示该产品首次退货率是否高于行业平均水平
	FirstProductReturnRateAbovePeers float64 `json:"first_product_return_rate_above_peers"`

	// 首次品退率与行业平均的对比数值
	FirstProductReturnRateVsIndustryAverage float64 `json:"first_product_return_rate_vs_industry_average"`

	// 首次品退率与行业优秀标准的对比数值
	FirstProductReturnRateVsIndustryExcellent float64 `json:"first_product_return_rate_vs_industry_excellent"`

	// 商品差评率，即收到的负向评价的比例
	ProductNegativeReviewRate float64 `json:"product_negative_review_rate"`

	// 商品差评率是否超过同行， 表示商品差评率是否高于行业平均水平
	ProductNegativeReviewRateAbovePeers float64 `json:"product_negative_review_rate_above_peers"`

	// 商品差评率与行业平均的对比数值
	ProductNegativeReviewRateVsIndustryAverage float64 `json:"product_negative_review_rate_vs_industry_average"`

	// 商品差评率与行业优秀标准的对比数值
	ProductNegativeReviewRateVsIndustryExcellent float64 `json:"product_negative_review_rate_vs_industry_excellent"`

	// 退款率，即退款请求成功的比率
	RefundRate float64 `json:"refund_rate"`

	// 退款成功金额，统计周期内成功退款的总额
	RefundSuccessfulAmount float64 `json:"refund_successful_amount"`

	// 违规数量，记录商家在统计周期内的违规操作次数
	ViolationCount int `json:"violation_count"`
}

//店铺服务分析趋势返回数据
type ShopServiceAnalysisTrendList struct {
	Records []ShopServiceAnalysisTrendNode `json:"records"`
}

// 店铺服务分析
type ShopServiceAnalysisTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 首次品退率，即第一次购买后发生的退货比率
	FirstProductReturnRate float64 `json:"first_product_return_rate"`

	// 首次品退率是否超过同行， 表示该产品首次退货率是否高于行业平均水平
	FirstProductReturnRateAbovePeers float64 `json:"first_product_return_rate_above_peers"`

	// 首次品退率与行业平均的对比数值
	FirstProductReturnRateVsIndustryAverage float64 `json:"first_product_return_rate_vs_industry_average"`

	// 首次品退率与行业优秀标准的对比数值
	FirstProductReturnRateVsIndustryExcellent float64 `json:"first_product_return_rate_vs_industry_excellent"`

	// 商品差评率，即收到的负向评价的比例
	ProductNegativeReviewRate float64 `json:"product_negative_review_rate"`

	// 商品差评率是否超过同行， 表示商品差评率是否高于行业平均水平
	ProductNegativeReviewRateAbovePeers float64 `json:"product_negative_review_rate_above_peers"`

	// 商品差评率与行业平均的对比数值
	ProductNegativeReviewRateVsIndustryAverage float64 `json:"product_negative_review_rate_vs_industry_average"`

	// 商品差评率与行业优秀标准的对比数值
	ProductNegativeReviewRateVsIndustryExcellent float64 `json:"product_negative_review_rate_vs_industry_excellent"`

	// 退款率，即退款请求成功的比率
	RefundRate float64 `json:"refund_rate"`

	// 退款成功金额，统计周期内成功退款的总额
	RefundSuccessfulAmount float64 `json:"refund_successful_amount"`

	// 违规数量，记录商家在统计周期内的违规操作次数
	ViolationCount int `json:"violation_count"`
}

// 客服
type CustomerService struct {

	// 客服销售额，记录客服团队在一定时期内的销售总额
	CustomerServiceSales float64 `json:"customer_service_sales"`

	// 客服销售额占比，即客服销售额占总销售额的比例
	CustomerServiceSalesRatio float64 `json:"customer_service_sales_ratio"`

	// 平均响应时长（秒），客服人员对客户咨询的平均响应时间
	AverageResponseTimeInSeconds float64 `json:"average_response_time_in_seconds"`

	// 客户满意率，衡量客户服务满意度的整体百分比
	CustomerSatisfactionRate float64 `json:"customer_satisfaction_rate"`

	// 询单转化率，即从顾客发起咨询到最终下单购买的转化比例
	InquiryConversionRate float64 `json:"inquiry_conversion_rate"`
}

//客服趋势返回数据
type CustomerServiceTrendList struct {
	Records []CustomerServiceTrendNode `json:"records"`
}

// 客服趋势
type CustomerServiceTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 客服销售额，记录客服团队在一定时期内的销售总额
	CustomerServiceSales float64 `json:"customer_service_sales"`

	// 客服销售额占比，即客服销售额占总销售额的比例
	CustomerServiceSalesRatio float64 `json:"customer_service_sales_ratio"`

	// 平均响应时长（秒），客服人员对客户咨询的平均响应时间
	AverageResponseTimeInSeconds float64 `json:"average_response_time_in_seconds"`

	// 客户满意率，衡量客户服务满意度的整体百分比
	CustomerSatisfactionRate float64 `json:"customer_satisfaction_rate"`

	// 询单转化率，即从顾客发起咨询到最终下单购买的转化比例
	InquiryConversionRate float64 `json:"inquiry_conversion_rate"`
}

// 客户分析
type CustomerAnalysis struct {
	// 会员人数，当前平台或业务的总会员数量
	TotalMembershipCount int64 `json:"total_membership_count"`

	// 昨日招募会员数，前一天新加入的会员数量
	MembersRecruitedYesterday int64 `json:"members_recruited_yesterday"`

	// 会员成交渗透率，会员群体中产生交易行为的比例
	MemberTransactionPenetration float64 `json:"member_transaction_penetration"`

	// 行业达标渗透率，与行业内其他企业相比，达到标准水平的会员渗透率
	IndustryStandardPenetration float64 `json:"industry_standard_penetration"`

	// 行业优质渗透率，在行业内表现出优秀品质特征的会员渗透率
	IndustryPremiumPenetration float64 `json:"industry_premium_penetration"`

	// 会员客单价，会员用户平均每单消费金额
	MemberAverageOrderValue float64 `json:"member_average_order_value"`
}

//客户分析趋势返回数据
type CustomerAnalysisTrendList struct {
	Records []CustomerAnalysisTrendNode `json:"records"`
}

// 客户分析趋势
type CustomerAnalysisTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 会员人数，当前平台或业务的总会员数量
	TotalMembershipCount int64 `json:"total_membership_count"`

	// 昨日招募会员数，前一天新加入的会员数量
	MembersRecruitedYesterday int64 `json:"members_recruited_yesterday"`

	// 会员成交渗透率，会员群体中产生交易行为的比例
	MemberTransactionPenetration float64 `json:"member_transaction_penetration"`

	// 行业达标渗透率，与行业内其他企业相比，达到标准水平的会员渗透率
	IndustryStandardPenetration float64 `json:"industry_standard_penetration"`

	// 行业优质渗透率，在行业内表现出优秀品质特征的会员渗透率
	IndustryPremiumPenetration float64 `json:"industry_premium_penetration"`

	// 会员客单价，会员用户平均每单消费金额
	MemberAverageOrderValue float64 `json:"member_average_order_value"`
}

// 客户分析
type CustomerLossAnalysis struct {
	// 流失金额，一定时期内因客户流失而损失的总金额
	AmountOfLoss float64 `json:"amount_of_loss"`

	// 流失人数，一定时期内流失的客户总人数
	LostMembers int64 `json:"lost_members"`

	// 引起流失的店铺数，在统计期内导致客户流失的店铺个数
	StoresCausingLoss int64 `json:"stores_causing_loss"`
}

//客户损失分析趋势返回数据
type CustomerLossAnalysisTrendList struct {
	Records []CustomerLossAnalysisTrendNode `json:"records"`
}

// 客户分析趋势
type CustomerLossAnalysisTrendNode struct {
	// 日期
	Date string `json:"date"`
	// 流失金额，一定时期内因客户流失而损失的总金额
	AmountOfLoss float64 `json:"amount_of_loss"`

	// 流失人数，一定时期内流失的客户总人数
	LostMembers int64 `json:"lost_members"`

	// 引起流失的店铺数，在统计期内导致客户流失的店铺个数
	StoresCausingLoss int64 `json:"stores_causing_loss"`
}
