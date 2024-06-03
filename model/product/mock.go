package product

import "math/rand"

func (i *RespProductIndex) Random() {

}

func (i *RespPricePower) Random(datelist []string) {
	i.Records = []PricePowerNode{}
	// for _, date := range datelist {
	// 	i.Records = append(i.Records, PricePowerNode{Date: date})
	// }
}

func (i *RespSku) Random() {
	i.Records = []SKUNode{}
	// i.Records = append(i.Records, SKUNode{SKUName: "尺寸30"})
	// i.Records = append(i.Records, SKUNode{SKUName: "尺寸40"})

}

func (i *RespReview) Random() {
	i.Records = []ReviewNode{}
	//i.Records = append(i.Records, ReviewNode{Keyword: "不满意", Count: 10})
	//i.Records = append(i.Records, ReviewNode{Keyword: "不想要", Count: 5})

}

func (i *RespProductDay) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, ProductPerformanceNode{Date: date})
	}
}

func (i *RespKeywordList) Random() {
	i.Records = []KeywordNode{}
	// i.Records = append(i.Records, KeywordNode{Keyword: "席梦思", Type: "手淘搜索"})
	// i.Records = append(i.Records, KeywordNode{Keyword: "席梦思", Type: "直通车"})
	// i.Records = append(i.Records, KeywordNode{Keyword: "床垫", Type: "手淘搜索"})
	// i.Records = append(i.Records, KeywordNode{Keyword: "床垫", Type: "直通车"})
}

func (i *RespIndexTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, IndexTrendNode{Date: date, Gmv: float64(rand.Intn(100000))})
	}
}
