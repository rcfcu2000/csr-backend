package traffic

//渠道分布 数据
type RespTrafficChannelData struct {
	Records []TrafficChannelNode `json:"records"`
}

// 渠道数据
type TrafficChannelNode struct {
	// 日期
	Date string `json:"date"`
	// 三级来源信息
	TertiarySource string `json:"tertiary_source"`
	// 访客数
	VisitorsCount int64 `json:"visitors_count"`
	// Gmv
	PaidAmount float64 `json:"gmv"`
}

//渠道gmv visitor 趋势数据
type RespTrafficTrendData struct {
	Records []TrafficChannelNode `json:"records"`
}

//货盘数据数据
type RespPalletData struct {
	Records []TrafficPalletNode `json:"records"`
}

// 货盘数据
type TrafficPalletNode struct {
	// pallet 从上到下按照S,A,B,C,D,- 排序  ORDER BY FIELD(status, 'active', 'pending', 'inactive');
	Pallet string `json:"pallet"`
	// count
	Count string `json:"count"`
}

//商品信息数据返回
type RespProductInfoListData struct {
	Records []ProductInfoNode `json:"records"`
}

// 商品信息数据 需要补全
type ProductInfoNode struct {
	// 产品ID
	ProductId string `json:"product_id"`
	// 商品简称
	ProductAlias string `json:"product_alias"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 	责任人
	Responsible string `json:"responsible"`
	// 二级类目名称
	Category_lv2 string `json:"category_lv2"`
	// 流量归属原则
	Belong string `json:"belong"`
	// 货盘
	Pallet string `json:"pallet"`
	// GMV(减退款)
	Gmv float64 `json:"gmv"`
	// 净利润
	Profit int64 `json:"profit"`
	// 本期货盘
	CurPallet string `json:"cur_pallet"`
	// 上期货盘
	PrePalleta string `json:"pre_pallet"`
	// 货盘变化
	PalletChange int64 `json:"pallet_change"`
	// 老客占比
	OldPercentage float64 `json:"old_percentage"`
	// 付费GMV占比
	PayGmvPercentage float64 `json:"pay_gmv_percentage"`
	// 推广花费占比
	SpendPercentage float64 `json:"spend_percentage"`
	// 综合得分
	OverallScore float64 `json:"overall_score"`
	// 加购效率
	AddCarEfficiency float64 `json:"add_car_efficiency"`
	// 退款效率
	RefundEfficiency float64 `json:"refund_efficiency""`
	// 复购效率
	Repurchase_efficiency float64 `json:"repurchase_efficiency"`
	// 流量效率
	Loss_efficiency float64 `json:"loss_efficiency"`
	// 转化效率
	ConversionEfficiency float64 `json:"conversion_efficiency"`
	// 支付转化率 支付转化率=支付买家数/访客数
	PayConversionRate float64 `json:"pay_conversion_rate"`
	// UV价值  UV价值=支付金额/商品访客数
	UV float64 `json:"uv"`
	// 跳失率 跳失率 =商品详情页跳出人数/访客数
	BounceRate float64 `json:"bounce_rate"`
	// 平均停留时长 平均停留时长=访客停留总时长/访客数
	AvgStayDuration float64 `json:"avg_stay_duration"`
	// 访问深度  访问深度=商品浏览量/访客数
	DepthVisit float64 `json:"depth_visit"`
}

//流量信息数据返回
type RespTrafficListData struct {
	Records []TrafficNode `json:"records"`
}

// 流量信息数据
type TrafficNode struct {
	// 产品ID
	ProductId string `json:"product_id"`
	// 	责任人
	Responsible string `json:"responsible"`
	// 流量归属原则
	Belong string `json:"belong"`
	// 货盘
	Pallet string `json:"pallet"`
	// 三级来源
	Source_type_3 string `json:"source_type_3"`
	// 访客数
	VisitorsCount int64 `json:"visitors_count"`
	// 访客TGI
	VisitorTGI float64 `json:"visitor_tgi"`
	// 购买TGI
	BuyTGI float64 `json:"buy_tgi"`
	// 渠道属性
	ChannelAttribute string `json:"channel_attribute "`
	// 渠道属性差异
	ChannelDiff float64 `json:"channel_diff"`
	// 客单价
	CustomerUnitPrice float64 `json:"customer_unit_price"` // // 客单价
	// 支付买家数
	PaidBuyers int64 `json:"paid_buyers"` // // 支付买家数
	// 支付转化率
	PayConversionRate float64 `json:"pay_conversion_rate"`
	// 加购率
	AddCarRate int64 `json:"add_car_rate"`
	// 加购人数
	AddToCartBuyers float64 `json:"add_to_cart_buyers"`
	// 加购转化率
	AddCartConversionRate float64 `json:"add_cart_conversion_rate"`
	// 全店加购率
	ShopAddCartRate float64 `json:"shop_add_car_rate"`
	// 全店加购转化率
	ShopAddCartConversionRate float64 `json:"shop_add_cart_conversion_rate"`
	// UV价值
	UV float64 `json:"uv"`
	// 本品访客占比  -本品渠道访客占比=商品流量来源数据中单个渠道访客/全部渠道访客
	ProductVisitorPercentage float64 `json:"product_visitor_percentage"`
	// 全店访客占比 - 店铺本渠道访客占比 =店铺流量来源数据的单个渠道访客/全部渠道访客
	ShopVisitorPercentage float64 `json:"shop_visitor_percentage"`
	// 本品买家占比 -本品渠道支付买家占比=商品流量来源数据中单个渠道支付买家/全部渠道支付买家
	ProductBuyerPercentage float64 `json:"product_buyer_percentage"`
	// 全店买家占比 -店铺本渠道支付买家占比 =店铺流量来源数据的单个渠道支付买家/全部渠道支付买家
	ShopBuyerPercentage float64 `json:"shop_buyer_percentage"`
}

//新老客户数据
type RespNewOldCustomerListData struct {
	Records []NewOldCustomerNode `json:"records"`
}

// 货盘数据
type NewOldCustomerNode struct {
	// 日期
	Date string `json:"date"`
	// 店铺客户数
	TotalCustomers float64 `json:"total_customers"`
	// 客户新访
	NewVisits float64 `json:"new_visits"`
	// 客户新访占比
	NewVisitsPercentage float64 `json:"new_visits_percentage"`
	// 未购客户回访占比
	NonPurchaseReturnVisitsPercentage float64 `json:"non_purchase_return_visits_percentage"`
	// 已购客户回访占比
	PurchasedCustomerReturnVisitsPercentage float64 `json:"purchased_customer_return_visits_percentage"`
	// 店铺转化率
	ConversionRate float64 `json:"conversion_rate"`
	// 客户新访支付转化率
	NewVisitPaymentConversionRate float64 `json:"new_visit_payment_conversion_rate"`
	// 未购客户回访支付转化率
	ReturnVisitPaymentConversionRateNonPurchasers float64 `json:"return_visit_payment_conversion_rate_non_purchasers"`
	// 已购客户回访支付转化率
	ReturnVisitPaymentConversionRatePurchasers float64 `json:"return_visit_payment_conversion_rate_purchasers"`
	// 店铺UV价值
	ShopUv float64 `json:"shop_uv"`
	// 客户新访UV价值
	NewVisitUv float64 `json:"new_visit_uv"`
	// 未购客户回访UV价值
	NonPurchaserUv float64 `json:"non_purchaser_uv"`
	// 已购客户回访UV价值
	PurchaserUv float64 `json:"purchaser_uv"`
}

// 店铺客户数
// 客户新访
// 客户新访占比
// 未购客户回访占比
// 已购客户回访占比
// 店铺转化率
// 客户新访支付转化率
// 未购客户回访支付转化率
// 已购客户回访支付转化率
// 店铺UV价值
// 客户新访UV价值
// 未购客户回访UV价值
// 已购客户回访UV价值

//流量通道列表
type RespChannelsData struct {
	Records []ChannelNode `json:"records"`
}

type ChannelNode struct {
	Channel string `json:"channel"`
}
