package promotion

func (i *RespPromotionAllData) Random() {
	(&i.PromotionIndex1).Random()
	(&i.PromotionIndex2).Random()
	(&i.BidTypeAnalysis).Random()
	(&i.PalletCost).Random()
	(&i.KeywordCost).Random()
	(&i.CrowdSpend).Random()
	// (&i.ProductAnalysis).Random()
	// (&i.PlanAnalysis).Random()
}
func (i *RespPromotionSearchData) Random() {
	(&i.BidTypeAnalysis).Random()
	(&i.PalletCost).Random()
	(&i.KeywordCost).Random()
	(&i.CrowdSpend).Random()
}

func (i *RespPromotionIndex1) Random() {

}
func (i *RespPromotionIndex2) Random() {
	bidTypeList := []string{"场景", "关键词", "人群"}
	for _, item := range bidTypeList {
		i.Records = append(i.Records, PromotionIndex2Node{SceneCategory: item})
	}
}
func (i *RespBidTypeAnalysis) Random() {
	bidTypeList := []string{"出价方式1", "出价方式2", "出价方式3"}
	for _, item := range bidTypeList {
		i.Records = append(i.Records, BidTypeAnalysisNode{BidType: item})
	}

}
func (i *RespPalletCost) Random() {
	palletList := []string{"S", "A", "B", "C", "D", "-"}
	for _, item := range palletList {
		i.Records = append(i.Records, PalletCostNode{Pallet: item})
	}
}
func (i *RespKeywordCost) Random() {
	l := []string{"kw1", "kw2", "kw3", "kw4", "kw5"}
	for _, item := range l {
		i.Records = append(i.Records, KeywordCostNode{Keyword: item})
	}
}
func (i *RespCrowdSpend) Random() {
	l := []string{"智能定向", "推荐", "店铺", "互动", "深度"}
	for _, item := range l {
		i.Records = append(i.Records, CrowdSpendNode{Crowd: item})
	}
}
func (i *RespProductAnalysis) Random() {
	l := []string{"S", "A", "B", "C", "D", "-"}
	for _, item := range l {
		i.Records = append(i.Records, ProductAnalysisNode{Pallet: item})
	}
}
func (i *RespPlanAnalysis) Random() {
	// l := []string{"plan_A", "plan_B", "plan_C", "plan_D"}
	// for _, item := range l {
	// 	i.Records = append(i.Records, PlanAnalysisNode{CampaignName: item})
	// }
}
