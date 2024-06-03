package keywordsa

//返回关键词数据列表
type RespKeywordListData struct {
	Records []KeywordNode `json:"records"`
}

// 关键词数据
type KeywordNode struct {
	// 日期
	Date string `json:"date"`
	// 词
	Keyword string `json:"keyword"`
	// 访客数
	VisitorsCount float64 `json:"visitors_count"` //点击量
	// Gmv
	// PaidAmount float64 `json:"gmv"`
	// 支付转化率 - 成功支付订单数占下单数的比例
	PaymentConversionRate float64 `json:"payment_conversion_rate"` // 订单支付转化率 支付转化率 = 总成交笔数 / 点击量
}

//返回生参付费词&免费词数据列表
type RespScKeywordListData struct {
	RecordsFree    []KeywordNode `json:"records_free"`
	RecordsNotFree []KeywordNode `json:"records_notfree"`
}

//返回生参行业词数据列表
type RespIndustryKeywordListData struct {
	Records []KeywordNode `json:"records"`
}

//返回词趋势数据列表
type RespKeywordTrendListData struct {
	Records []KeywordMuNode `json:"records"`
}

// 关键词数据
type KeywordMuNode struct {
	// 日期
	Date string `json:"date"`
	// 词
	Keyword string `json:"keyword"`
	// 计数
	Count int64 `json:"count"`
	//   无界词-点击量
	Clicks float64 `json:"clicks"` //点击量
	// 生参免费词访客
	VisitorsCountFree float64 `json:"visitors_count_free"` //生参免费词访客
	//生参付费词访客
	VisitorsCountNotfree float64 `json:"visitors_count_notfree"` //生参付费词访客
	//   行业点击量
	IndustryClicks float64 `json:"industry_clicks"` //行业点击量

	//   无界词转化率
	Cr float64 `json:"cr"` //无界词转化率 cr =  conversion rate
	// 生参免费词转化率
	CrFree float64 `json:"cr_free"` //生参免费词转化率
	//生参付费词转化率
	CrNotfree float64 `json:"cr_notfree"` //生参付费词转化率
	//   行业-转化率
	CrIndustry float64 `json:"cr_industry"` //行业-转化率
}

//返回关键词明细数据列表
type RespKeywordMuListData struct {
	Records []KeywordMuNode `json:"records"`
	Sum     KeywordMuNode   `json:"sum"`
}
