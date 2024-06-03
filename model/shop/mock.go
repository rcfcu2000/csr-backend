package shop

import "math/rand"

func (i *ShopIndexTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, ShopIndexTrendNode{Date: date, Ranking: (rand.Intn(20))})
	}
}
func (i *ContentTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, ContentTrendNode{Date: date, Type: "短视频", Amount: float64(rand.Intn(10000))})
		i.Records = append(i.Records, ContentTrendNode{Date: date, Type: "直播", Amount: float64(rand.Intn(10000))})
		i.Records = append(i.Records, ContentTrendNode{Date: date, Type: "图文", Amount: float64(rand.Intn(10000))})
	}
}
func (i *ShopServiceAnalysisTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, ShopServiceAnalysisTrendNode{Date: date, RefundSuccessfulAmount: float64(rand.Intn(2000))})
	}
}
func (i *CustomerServiceTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, CustomerServiceTrendNode{Date: date, CustomerServiceSales: float64(rand.Intn(2000))})
	}
}
func (i *CustomerAnalysisTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, CustomerAnalysisTrendNode{Date: date, TotalMembershipCount: int64(rand.Intn(200))})
	}
}
func (i *CustomerLossAnalysisTrendList) Random(datelist []string) {
	for _, date := range datelist {
		i.Records = append(i.Records, CustomerLossAnalysisTrendNode{Date: date, AmountOfLoss: float64(rand.Intn(2000))})
	}
}
