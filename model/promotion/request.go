package promotion

import (
	"xtt/model/common"
)

type ReqPromotionAllSearch struct {
	// 店铺筛选字段，用于存储店铺筛选条件的详细信息
	ShopFilter string `json:"shop_filter"` // 对应于店铺筛选条件的JSON键值
	// 商品负责人字段，记录商品负责人信息
	ProductManager []string `json:"product_manager"` // 商品负责人姓名或ID，对应于JSON中的"商品负责人"
	// 当期货盘字段，包含当前待选货盘的信息
	CurrentInventory []string `json:"current_inventory"` // 描述或标识当期货盘的详细信息
	// 场景分类字段，指定某种特定场景的分类
	SceneCategory []string `json:"scene_category"` // 场景分类的具体名称或ID
	StartDate     string   `json:"start_date"`     // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate       string   `json:"end_date"`       // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	// 货盘字段，局部货盘
	Pallet []string `json:"pallet"`
	// 出价方式字段，表示商品或服务的出价模式, 局部变量
	BidType []string `json:"bid_type"`
	// 关键词筛选字段，用于记录关键词搜索或过滤条件 局部变量
	KeywordFilter []string `json:"keyword_filter"`
	// 人群筛选字段，用于存储对目标人群的筛选条件 局部变量
	AudienceFilter []string `json:"audience_filter"`
	//输入 product_id   列表
	Ids []string `json:"ids"`

	PageNum  int `json:"pageNum"`  //分页，页号
	PageSize int `json:"pageSize"` //分页，每页数目
}

type ReqPromotionThendSearch struct {
	// 店铺筛选字段，用于存储店铺筛选条件的详细信息
	ShopFilter string `json:"shop_filter"` // 对应于店铺筛选条件的JSON键值
	// 商品负责人字段，记录商品负责人信息
	ProductManager []string `json:"product_manager"` // 商品负责人姓名或ID，对应于JSON中的"商品负责人"
	// 当期货盘字段，包含当前待选货盘的信息
	CurrentInventory []string `json:"current_inventory"` // 描述或标识当期货盘的详细信息
	// 场景分类字段，指定某种特定场景的分类
	SceneCategory []string `json:"scene_category"` // 场景分类的具体名称或ID
	StartDate     string   `json:"start_date"`     // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	EndDate       string   `json:"end_date"`       // 格式如："2023-01-01" // 日期 - 数据统计的时间点
	// 货盘字段，局部货盘
	Pallet []string `json:"pallet"`
	// 出价方式字段，表示商品或服务的出价模式, 局部变量
	BidType []string `json:"bid_type"`
	// 关键词筛选字段，用于记录关键词搜索或过滤条件 局部变量
	KeywordFilter []string `json:"keyword_filter"`
	// 人群筛选字段，用于存储对目标人群的筛选条件 局部变量
	AudienceFilter []string `json:"audience_filter"`

	//输入 product_id / plan_id 列表
	Ids []string `json:"ids"`
}

func (r *ReqPromotionAllSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
	if len(r.SceneCategory) == 0 {
		r.SceneCategory = []string{"场景推广", "关键词推广", "精准人群推广"}
	}
	if len(r.CurrentInventory) == 0 {
		r.CurrentInventory = []string{"S+", "S", "A", "B", "C", "D", "-"}
	}

	if len(r.BidType) == 0 {
		r.BidType = []string{"最大化拿成交",
			"控ROI",
			"控成本点击",
			"最大化拿点击",
			"最大化拿量",
			"控成本投放",
			"控投产比投放",
			"套餐包", "手动", "其他"}
	}

	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":shop", ValType: "string", ValString: r.ShopFilter})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pallet", ValType: "list", ValString: common.ListToInStringStr(r.CurrentInventory, "pallet")}) // "('A','B')"
	sqlp.SetKeyVal(common.KeyVal{Key: ":scene", ValType: "list", ValString: common.ListToInStringStr(r.SceneCategory, "scene")})
	if len(r.ProductManager) == 0 {
		sqlp.SetKeyVal(common.KeyVal{Key: ":resperson", ValType: "stringsrc", ValString: GetAllProductManager()})
	} else {
		sqlp.SetKeyVal(common.KeyVal{Key: ":resperson", ValType: "list", ValString: common.ListToInStringStr(r.ProductManager, "resperson")})
	}
	if len(r.Pallet) == 0 {
		r.Pallet = []string{"S+", "S", "A", "B", "C", "D", "-"}
	}
	sqlp.SetKeyVal(common.KeyVal{Key: ":subpallet", ValType: "list", ValString: common.ListToInStringStr(r.Pallet, "")})
	sqlp.SetKeyVal(common.KeyVal{Key: ":bidtype", ValType: "list", ValString: common.ListToInStringStr(r.BidType, "bidtype")})
	sqlp.SetKeyVal(common.KeyVal{Key: ":keyword", ValType: "list", ValString: common.ListToInStringStr(r.KeywordFilter, "keyword")})
	sqlp.SetKeyVal(common.KeyVal{Key: ":crowd", ValType: "list", ValString: common.ListToInStringStr(r.AudienceFilter, "crowd")})
	sqlp.SetKeyVal(common.KeyVal{Key: ":ids", ValType: "list", ValString: common.PidToInStringStr(r.Ids)})

	if r.PageNum <= 0 || r.PageSize <= 0 {
		r.PageNum = 1
		r.PageSize = 9999
	}
	ofset := (r.PageNum - 1) * r.PageSize
	sqlp.SetKeyVal(common.KeyVal{Key: ":offset", ValType: "int", ValInt: int64(ofset)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pageSize", ValType: "int", ValInt: int64(r.PageSize)})
}

func GetAllProductManager() string {
	return "(select DISTINCT responsible  from biz_product  ) "
}

func (r *ReqPromotionThendSearch) SetToSQLProccesor(sqlp *common.SQLProccesor) {
	if len(r.SceneCategory) == 0 {
		r.SceneCategory = []string{"场景推广", "关键词推广", "精准人群推广"}
	}
	if len(r.CurrentInventory) == 0 {
		r.CurrentInventory = []string{"S+", "S", "A", "B", "C", "D", "-"}
	}
	if len(r.BidType) == 0 {
		r.BidType = []string{"最大化拿成交",
			"控ROI",
			"控成本点击",
			"最大化拿点击",
			"最大化拿量",
			"控成本投放",
			"控投产比投放",
			"套餐包", "手动", "其他"}
	}

	sqlp.SetKeyVal(common.KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
	sqlp.SetKeyVal(common.KeyVal{Key: ":shop", ValType: "string", ValString: r.ShopFilter})
	sqlp.SetKeyVal(common.KeyVal{Key: ":pallet", ValType: "list", ValString: common.PalletToInStringStr(r.CurrentInventory)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":scene", ValType: "list", ValString: common.ListToInStringStr(r.SceneCategory, "scene")})
	if len(r.ProductManager) == 0 {
		sqlp.SetKeyVal(common.KeyVal{Key: ":resperson", ValType: "stringsrc", ValString: GetAllProductManager()})
	} else {
		sqlp.SetKeyVal(common.KeyVal{Key: ":resperson", ValType: "list", ValString: common.ListToInStringStr(r.ProductManager, "")})
	}
	if len(r.Pallet) == 0 {
		r.Pallet = []string{"S+", "S", "A", "B", "C", "D", "-"}
	}

	sqlp.SetKeyVal(common.KeyVal{Key: ":subpallet", ValType: "list", ValString: common.PalletToInStringStr(r.Pallet)})
	sqlp.SetKeyVal(common.KeyVal{Key: ":bidtype", ValType: "list", ValString: common.ListToInStringStr(r.BidType, "bidtype")})
	sqlp.SetKeyVal(common.KeyVal{Key: ":keyword", ValType: "list", ValString: common.ListToInStringStr(r.KeywordFilter, "keyword")})
	sqlp.SetKeyVal(common.KeyVal{Key: ":crowd", ValType: "list", ValString: common.ListToInStringStr(r.AudienceFilter, "crowd")})

	sqlp.SetKeyVal(common.KeyVal{Key: ":ids", ValType: "list", ValString: common.PidToInStringStr(r.Ids)})

	// if r.PageNum <= 0 || r.PageSize <= 0 {
	// 	r.PageNum = 1
	// 	r.PageSize = 9999
	// }
	// ofset := (r.PageNum - 1) * r.PageSize
	// sqlp.SetKeyVal(common.KeyVal{Key: ":offset", ValType: "int", ValInt: int64(ofset)})
	// sqlp.SetKeyVal(common.KeyVal{Key: ":pageSize", ValType: "int", ValInt: int64(r.PageSize)})

}
