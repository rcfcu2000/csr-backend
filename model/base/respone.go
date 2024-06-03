package base

//店铺列表列表
type RespShopListData struct {
	Records []ShopNode `json:"records"`
}

// 关键词数据
type ShopNode struct {
	// 店铺id
	ShopId string `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	//
	LinkName float64 `json:"linkname"`
}
