package target

type RespTargetIndexData struct {
	// GMV
	Gmv float64 `json:"gmv"`
	// 净利润
	Profit float64 `json:"profit"`
	// 净利润率
	ProfitRate float64 `json:"profit_rate"`
	// 推广花费
	Spend float64 `json:"spend"`
	// 推广占比
	PromotionPercentage float64 `json:"promotion_percentage"`
	// 累计GMV
	MonthGmv float64 `json:"month_gmv"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
	// 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`
}

// 分货盘目标达成数据返回
type RespPalletTargetListData struct {
	Records []PalletTargetNode `json:"records"`
}

// 分货盘目标达成
type PalletTargetNode struct {
	// 货盘
	Pallet string `json:"pallet"`
	// 累计GMV
	MonthGmv float64 `json:"month_gmv"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
	// 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`
}

// Gmv 趋势列表返回
type RespGmvTrendListData struct {
	Records []GmvTrendNode `json:"records"`
}

// Gmv趋势数据
type GmvTrendNode struct {
	// 日期
	Date string `json:"date"`
	// GMV
	Gmv float64 `json:"gmv"`
	// 推广GMV
	PromotionGmv float64 `json:"promotion_gmv"`
	// 推广花费
	Spend float64 `json:"spend"`
}

// 责任人目标数据返回
type RespProductManagerTargetListData struct {
	Records []ProductManagerTargetNode `json:"records"`
}

// 责任人目标数据
type ProductManagerTargetNode struct {
	// 责任人
	Manager string `json:"manager"`
	// 三级类目
	CategoryLv3 string `json:"category_lv3"`
	// GMV
	Gmv float64 `json:"gmv"`
	// 时间进度, 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
	// 净利润
	Profit float64 `json:"profit"`
	// 经营利润目标
	TargetProfit float64 `json:"profit_rate"`

	// 经营利润达成率
	ProfitTargetRate float64 `json:"profit_target_rate"`
}

// 三级类目目标数据返回
type RespCategoryTargetListData struct {
	Records []CategoryTargetNode `json:"records"`
}

// 三级类目目标数据
type CategoryTargetNode struct {
	// 	日期
	// 	三级类目
	// 	商品简称
	// 	GMV
	// 	时间进度
	// 	GMV达成率
	// 	GMV目标
	// 责任人
	Manager string `json:"manager"`
	// 三级类目
	CategoryLv3 string `json:"category_lv3"`
	// 商品简称
	ProductName string `json:"product_name"`
	// GMV
	Gmv float64 `json:"gmv"`
	// 时间进度, 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
}

// 货盘目标详细数据返回
type RespPalletTargetInfoListData struct {
	Records []PalletTargetInfoNode `json:"records"`
}

// 货盘目标详细数据
// 需要补充
type PalletTargetInfoNode struct {
	// 货盘
	Pallet string `json:"pallet"`
	// 累计GMV
	MonthGmv float64 `json:"month_gmv"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
	// 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`

	// 推广花费
	Spend float64 `json:"spend"`
	// 推广预算
	MonthlyBudget float64 `json:"monthly_budget"`
	// 推广占比            -- 推广占比 =推广花费/GMV
	PromotionPercentage float64 `json:"promotion_percentage"`
	// 推广目标占比 -- 推广目标占比 =推广预算/GMV目标
	PromotionTargetPercentage float64 `json:"promotion_target_percentage"`
	// 推广差异  -- 推广差异 = 推广花费-推广花费预算
	PromotionDiff float64 `json:"promotion_diff"`
	// 时间进度 -- 时间进度=结束日期的天数/结束日期所在月的总天数
	TimeSchedule float64 `json:"time_schedule"`
	// 综合ROI  -- 商品每日数据
	CompositeRoi float64 `json:"composite_roi"`
	// 客单价 -- 客单价 = 支付金额/支付买家数
	CustomerUnitPrice float64 `json:"customer_unit_price"`
	// 净利润 -- 万相台-宝贝主体报表 商品每日数据	净利润=（b.支付金额-b.成功退款金额）*c.预估毛利率-b.支付件数*c.发货费用-c.人工费用*天数-a.花费
	Profit float64 `json:"profit"`
	// 支付人数 -- 商品每日数据	支付买家数
	PaidBuyers float64 `json:"paid_buyers"`
}

// 产品目标详细数据返回
type RespProductListData struct {
	Records []ProductTargetNode `json:"records"`
}

// 产品目标详细数据
// 需要补充
type ProductTargetNode struct {

	//商品ID
	ProductId string `json:"product_id"`
	//商品简称
	ProductName string `json:"product_name"`
	// 累计GMV
	MonthGmv float64 `json:"month_gmv"`
	// GMV目标
	TargetGmv float64 `json:"target_gmv"`
	// GMV达成率
	TargetGmvRate float64 `json:"target_gmv_rate"`
	// 达成率目标
	TargetDayRate float64 `json:"target_day_rate"`

	// 推广花费
	Spend float64 `json:"spend"`
	// 推广预算
	MonthlyBudget float64 `json:"monthly_budget"`
	// 推广占比            -- 推广占比 =推广花费/GMV
	PromotionPercentage float64 `json:"promotion_percentage"`
	// 推广目标占比 -- 推广目标占比 =推广预算/GMV目标
	PromotionTargetPercentage float64 `json:"promotion_target_percentage"`
	// 推广差异  -- 推广差异 = 推广花费-推广花费预算
	PromotionDiff float64 `json:"promotion_diff"`
	// 时间进度 -- 时间进度=结束日期的天数/结束日期所在月的总天数
	TimeSchedule float64 `json:"time_schedule"`
	// 综合ROI  -- 商品每日数据
	CompositeRoi float64 `json:"composite_roi"`
	// 客单价 -- 客单价 = 支付金额/支付买家数
	CustomerUnitPrice float64 `json:"customer_unit_price"`
	// 净利润 -- 万相台-宝贝主体报表 商品每日数据	净利润=（b.支付金额-b.成功退款金额）*c.预估毛利率-b.支付件数*c.发货费用-c.人工费用*天数-a.花费
	Profit float64 `json:"profit"`
	// 支付人数 -- 商品每日数据	支付买家数
	PaidBuyers float64 `json:"paid_buyers"`
	// 类目名称  -- 商品主数据	三级类目
	CategoryLv3 string `json:"category_lv3"`
}
