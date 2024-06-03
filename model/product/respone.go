package product

//返回商品单页所有数据
type RespProductAllData struct {
	ProductIndex RespProductIndex `json:"productIndex"` //重要指标
}

//商品单页指标
type RespProductIndex struct {
	// 访客数
	VisitorsCount uint `gorm:"column:visitors_count;comment:访客数。" json:"visitors_count"` //  访客数
	// Gmv
	Gmv float64 `json:"gmv"`
	// 支付转化率 - 成功支付订单数占下单数的比例
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率
	// 搜索访客占比 - 来自搜索渠道访问商品详情页的访客占总访客数的比例
	SearchVisitorRatio float64 `json:"search_visitor_ratio"` // 搜索渠道访问商品的访客比例
	// 搜索GMV占比 - 来自搜索渠道产生的GMV占总GMV的比例
	SearchGMVRatio float64 `gorm:"column:search_gmv_ratio;comment:搜索GMV占比" json:"search_gmv_ratio"` // 搜索渠道贡献的GMV占比
	// 老客占比 - 老客户购买商品的数量或金额占总数量或金额的比例
	ReturningCustomerRatio float64 `gorm:"column:returning_customer_ratio;comment:老客户占比" json:"returning_customer_ratio"` // 老顾客购买占比
	// 退款率 - 发生退款的订单占总订单数的比例
	RefundRate float64 `json:"refund_rate"` // 退款率
	// 免费搜索点击率
	FreeSearchClickRate float64 `json:"free_search_click_rate"` // 免费搜索点击率
	// 连带购买也子类目宽度
	BundlePurchase float64 `json:"bundle_purchase"` //连带购买
	//复购率
	RepeatPurchaseRate float64 `json:"repeat_purchase_rate"` // 复购率
}

//指标日趋势返回数据
type RespIndexTrendList struct {
	Records []IndexTrendNode `json:"records"`
}

//商品单页指标日数据
type IndexTrendNode struct {
	// date
	Date string `json:"date"`
	// 访客数
	VisitorsCount uint `gorm:"column:visitors_count;comment:访客数。" json:"visitors_count"` //  访客数
	// Gmv
	Gmv float64 `json:"gmv"`
	// 支付转化率 - 成功支付订单数占下单数的比例
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率
	// 搜索访客占比 - 来自搜索渠道访问商品详情页的访客占总访客数的比例
	SearchVisitorRatio float64 `json:"search_visitor_ratio"` // 搜索渠道访问商品的访客比例
	// 搜索GMV占比 - 来自搜索渠道产生的GMV占总GMV的比例
	SearchGMVRatio float64 `gorm:"column:search_gmv_ratio;comment:搜索GMV占比" json:"search_gmv_ratio"` // 搜索渠道贡献的GMV占比
	// 老客占比 - 老客户购买商品的数量或金额占总数量或金额的比例
	ReturningCustomerRatio float64 `json:"returning_customer_ratio"` // 老顾客购买占比
	// 退款率 - 发生退款的订单占总订单数的比例
	RefundRate float64 `json:"refund_rate"` // 百分比形式
	// 免费搜索点击率
	FreeSearchClickRate float64 `json:"free_search_click_rate"` // 免费搜索点击率
	// 连带购买也子类目宽度
	BundlePurchase float64 `json:"bundle_purchase"` //连带购买
	//复购率
	RepeatPurchaseRate float64 `json:"repeat_purchase_rate"` // 复购率
}

//3 张图
type RespChart3 struct {
	PricePower RespPricePower `json:"price_power"` //价格力
	Sku        RespSku        `json:"sku"`         //sku
	Review     RespReview     `json:"review"`      //评价
}

// SKU 占比分析
type RespPricePower struct {
	Records []PricePowerNode `json:"records"`
}

//价格力
type PricePowerNode struct {
	// 日期
	Date string `json:"date"`
	// 价格力星级
	PPLevel float64 `json:"pp_level"` // 价格力星级  int64
	// 单价
	UnitPrice float64 `json:"unit_price"` // 单价
}

// SKU 占比分析
type RespSku struct {
	Records []SKUNode `json:"records"`
}

// SKU 占比分析
type SKUNode struct {
	// // 日期，格式按照实际需求定义，例如："2022-01-01" 或者 int64 时间戳等
	// Date string `json:"date"`
	// SKU名称，商品的唯一标识符
	SKUName string `json:"sku_name"`
	// 支付金额，单位通常为元或者其他货币单位，精度可能需要小数点后两位
	PayAmount float64 `json:"pay_amount"` // 支付金额
	// 支付买家数，完成支付的用户数量
	PayBuyers int `json:"pay_buyers"` // 支付买家数
	// 支付件数，该时间段内成功支付的商品件数
	PayQuantity int `json:"pay_quantity"` // 支付件数
	// 加购件数，该时间段内被加入购物车的商品件数
	AddToCartCount int `json:"add_to_cart_count"` // 加购件数
}

// 评价分析数据
type RespReview struct {
	Records []ReviewNode `json:"records"`
}

// 评价关键词数据
type ReviewNode struct {
	Keyword string `json:"keyword"` // 评价关键词
	Count   int    `json:"count"`   // 出现次数，即关键词在评价中出现的累计频次
}

//关系词分析数据返回数据
type RespKeywordList struct {
	Count   int64         `json:"count"`
	Records []KeywordNode `json:"records"`
	Sum     []KeywordNode `json:"sum"`
}

// 关系词分析数据
type KeywordNodeDel struct {
	// 类型，代表不同类型的营销活动或者统计数据类型
	Type string `json:"type" comment:"类型，标识不同的营销活动或统计维度"`

	// 关键词，可能与广告投放或商品搜索相关的核心词汇
	Keyword string `json:"keyword" comment:"关键词，如广告投放关键词或搜索关键词"`

	// 访客数，统计周期内的页面访问者数量
	VisitorCount int `json:"visitor_count" comment:"访客数，表示期间内访问该页面或参与活动的独立访客数"`

	// 加购率，计算公式为加入购物车的商品数量除以访客数
	AddToCartRate float64 `json:"add_to_cart_rate" comment:"加购率，即加入购物车的用户占总访客的比例"`

	// 转化率，计算公式为完成购买的订单数除以访客数
	ConversionRate float64 `json:"conversion_rate" comment:"转化率，即最终转化为购买行为的访客占比"`

	// 粉丝支付买家数，统计期内通过关注店铺并产生购买行为的买家数量
	FansPaidBuyersCount int `json:"fans_paid_buyers_count" comment:"粉丝支付买家数，指关注店铺后下单并完成支付的买家数"`

	// 直接支付买家数，统计期内未经过其他环节直接下单并完成支付的买家数量
	DirectPaidBuyersCount int `json:"direct_paid_buyers_count" comment:"直接支付买家数，表示没有经过收藏或加购等中间环节直接下单付款的买家数"`
}

// 关系词分析数据
type KeywordNode struct {

	// 关键词，可能与广告投放或商品搜索相关的核心词汇
	Keyword string `json:"keyword" comment:"关键词，如广告投放关键词或搜索关键词"`

	// Search访客数，统计周期内的页面访问者数量
	SearchVisitorCount int `json:"search_visitor_count"`

	// 加购率，计算公式为加入购物车的商品数量除以访客数
	SearchAddToCartRate float64 `gorm:"column:search_cart_addition_rate" json:"search_add_to_cart_rate"`

	// 转化率，计算公式为完成购买的订单数除以访客数
	SearchConversionRate float64 `json:"search_conversion_rate" comment:"转化率，即最终转化为购买行为的访客占比"`

	// 粉丝支付买家数，统计期内通过关注店铺并产生购买行为的买家数量
	SearchFansPaidBuyersCount int `json:"search_fans_paid_buyers_count" comment:"粉丝支付买家数，指关注店铺后下单并完成支付的买家数"`

	// 直接支付买家数，统计期内未经过其他环节直接下单并完成支付的买家数量
	SearchDirectPaidBuyersCount int `json:"search_direct_paid_buyers_count"`

	// Ztc 访客数，统计周期内的页面访问者数量
	ZtcVisitorCount int `json:"ztc_visitor_count"`

	// Ztc 加购率，计算公式为加入购物车的商品数量除以访客数
	ZtcAddToCartRate float64 `json:"ztc_add_to_cart_rate"`

	// Ztc 转化率，计算公式为完成购买的订单数除以访客数
	ZtcConversionRate float64 `json:"ztc_conversion_rate" comment:"转化率，即最终转化为购买行为的访客占比"`

	// Ztc 粉丝支付买家数，统计期内通过关注店铺并产生购买行为的买家数量
	ZtcFansPaidBuyersCount int `json:"ztc_fans_paid_buyers_count" comment:"粉丝支付买家数，指关注店铺后下单并完成支付的买家数"`

	// Ztc 直接支付买家数，统计期内未经过其他环节直接下单并完成支付的买家数量
	ZtcDirectPaidBuyersCount int `json:"ztc_direct_paid_buyers_count"`

	// 合计 访客数，统计周期内的页面访问者数量
	ZjVisitorCount int `json:"zj_visitor_count"`

	// 合计 加购率，计算公式为加入购物车的商品数量除以访客数
	ZjAddToCartRate float64 `json:"zj_add_to_cart_rate"`

	// 合计 转化率，计算公式为完成购买的订单数除以访客数
	ZjConversionRate float64 `json:"zj_conversion_rate" comment:"转化率，即最终转化为购买行为的访客占比"`

	// 合计 粉丝支付买家数，统计期内通过关注店铺并产生购买行为的买家数量
	ZjFansPaidBuyersCount int `json:"zj_fans_paid_buyers_count" comment:"粉丝支付买家数，指关注店铺后下单并完成支付的买家数"`

	// 合计 直接支付买家数，统计期内未经过其他环节直接下单并完成支付的买家数量
	ZjDirectPaidBuyersCount int `json:"zj_direct_paid_buyers_count"`
}

//每日明细
type RespProductDay struct {
	Count   int64                    `json:"count"`
	Records []ProductPerformanceNode `json:"records"`
	Sum     ProductPerformanceNode   `json:"sum"`
}

// ProductPerformance 结构体用于表示商品在指定日期的各项关键运营指标。
type ProductPerformanceNode struct {
	// 日期，记录各项数据对应的日期时间
	Date string `json:"date" comment:"日期，记录各项指标对应的时间点"`

	// 商品访客数，该商品在当天被访问的独立访客数
	ProductVisitorCount int `json:"product_visitor_count" comment:"商品访客数，表示该商品在当日被访问的独立访客数"`

	// GMV，Gross Merchandise Volume，商品成交总额
	GMV float64 `json:"gmv" comment:"GMV，商品成交总额，即商品在当天的总销售额"`

	// 支付转化率，计算公式为实际支付订单数除以商品访客数
	PaymentConversionRate float64 `json:"payment_conversion_rate" comment:"支付转化率，表示实际完成支付的订单数占商品访客数的比例"`

	// 搜索访客占比，搜索渠道带来的访客在整个商品访客中的比例
	SearchVisitorRatio float64 `json:"search_visitor_ratio" comment:"搜索访客占比，表示通过搜索进入商品详情页的访客占总商品访客的比例"`

	// 搜索GMV占比，搜索渠道产生的GMV占总GMV的比例
	SearchGMVRatio float64 `json:"search_gmv_ratio" comment:"搜索GMV占比，表示通过搜索渠道带来的销售额占总GMV的比例"`

	// 老客占比，历史购买过的顾客再次购买的占比
	ReturningCustomerRatio float64 `json:"returning_customer_ratio" comment:"老客占比，表示历史有购买记录的客户再次购买的次数占总购买次数的比例"`

	// 退款率，计算公式为退款订单金额除以GMV
	RefundRate float64 `json:"refund_rate" comment:"退款率，表示退款订单金额占总GMV的比例"`

	// 价格力星级，反映商品价格竞争力的评级
	PricePowerStars int `json:"price_power_stars" comment:"价格力星级，评估商品价格优势和市场竞争力的等级"`

	// 价格力额外曝光，由于价格竞争力获得的额外流量曝光量
	PricePowerExtraExposure float64 `gorm:"column:price_power_extra_exposure;comment:价格力额外曝光" json:"price_power_extra_exposure" comment:"价格力额外曝光，表示因价格竞争力而获得的额外流量曝光次数"`

	// 免费搜索点击率，搜索结果中商品被点击的概率
	FreeSearchClickThroughRate float64 `gorm:"column:free_search_click_through_rate;comment:免费搜索点击率"  json:"free_search_click_through_rate" comment:"免费搜索点击率，表示商品在自然搜索结果中被点击的次数占展示次数的比例"`

	// 连带购买也子类目宽度，衡量商品关联销售能力的一个指标
	AssociatedPurchaseSubcategoryWidth float64 `json:"associated_purchase_subcategory_width" comment:"连带购买子类目宽度，表示商品在跨类目关联销售时所涉及的子类目宽度，非标准指标，可根据实际情况使用"`

	// 复购率，已购买过该商品的客户再次购买的比率
	RepeatPurchaseRate float64 `json:"repeat_purchase_rate" comment:"复购率，表示曾经购买过商品的客户再次购买该商品的频率"`

	// 推广花费，商品在当天投入的推广费用总额
	PromotionCost float64 `json:"promotion_cost" comment:"推广花费，记录该商品当日用于各类推广活动所消耗的费用总额"`

	// 推广ROI，Return On Investment，推广投资回报率
	PromotionROI float64 `json:"promotion_roi" comment:"推广ROI，计算公式为（GMV - 推广花费）/ 推广花费，反映推广投入产出效果的比例"`
}

//产品名称列表
type RespProductList struct {
	Records []ProductInfo `json:"records"`
}

//产品信息
type ProductInfo struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}
