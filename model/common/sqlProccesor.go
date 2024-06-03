package common

import (
	"fmt"
	"strings"
	// biz "xtt/model/biz"
	// "xtt/model/promotion"
)

type KeyVal struct {
	Key       string
	ValInt    int64
	ValString string
	ValFloat  float64
	ValType   string // int float string
}

type SQLProccesor struct {
	RawSql    string
	KeyValMap map[string]KeyVal
}

func (i *SQLProccesor) SetSql(sql string) {
	i.RawSql = sql
}

func (i *SQLProccesor) SetKeyVal(kv KeyVal) {
	if i.KeyValMap == nil {
		i.KeyValMap = map[string]KeyVal{}
	}
	i.KeyValMap[kv.Key] = kv
}

func (i *SQLProccesor) GetResult() string {
	if i.KeyValMap == nil {
		i.KeyValMap = map[string]KeyVal{}
	}
	result := i.RawSql
	for _, v := range i.KeyValMap {
		if v.ValType == "string" {
			result = strings.ReplaceAll(result, v.Key+" ", fmt.Sprintf("'%s'"+" ", v.ValString))
		}
		if v.ValType == "stringlike" {
			result = strings.ReplaceAll(result, v.Key, v.ValString)
		}
		if v.ValType == "stringsrc" {
			result = strings.ReplaceAll(result, v.Key+" ", fmt.Sprintf("%s"+" ", v.ValString))
		}
		if v.ValType == "int" {
			result = strings.ReplaceAll(result, v.Key+" ", fmt.Sprintf("%d"+" ", v.ValInt))
		}
		if v.ValType == "float" {
			result = strings.ReplaceAll(result, v.Key+" ", fmt.Sprintf("%.4f"+" ", v.ValFloat))
		}
		if v.ValType == "list" {
			result = strings.ReplaceAll(result, v.Key+" ", fmt.Sprintf("%s"+" ", v.ValString))
		}
	}
	return result
}

// func (i *SQLProccesor) SetParam_ReqAll(r biz.ReqHuopanAllSearch) {
// 	i.SetKeyVal(KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
// 	i.SetKeyVal(KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
// 	i.SetKeyVal(KeyVal{Key: ":pallet", ValType: "list", ValString: ToInStringStr(r.CurrentInventory)}) // "('A','B')"
// 	i.SetKeyVal(KeyVal{Key: ":pallet_change", ValType: "list", ValString: ToInStringInt(r.InventoryChange)})
// 	i.SetKeyVal(KeyVal{Key: ":resperson", ValType: "string", ValString: r.ProductManager})
// 	i.SetKeyVal(KeyVal{Key: ":category_lv1", ValType: "string", ValString: r.PrimaryCategory})
// 	i.SetKeyVal(KeyVal{Key: ":category_lv2", ValType: "string", ValString: r.SecondaryCategory})
// 	i.SetKeyVal(KeyVal{Key: ":category_lv3", ValType: "string", ValString: r.TertiaryCategory})
// }

// func (i *SQLProccesor) SetParam_ReqRange(r biz.ReqPrductPriceRangeSearch) {
// 	i.SetKeyVal(KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
// 	i.SetKeyVal(KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
// 	i.SetKeyVal(KeyVal{Key: ":pallet", ValType: "list", ValString: ToInStringStr(r.CurrentInventory)}) // "('A','B')"
// 	i.SetKeyVal(KeyVal{Key: ":pallet_change", ValType: "list", ValString: ToInStringInt(r.InventoryChange)})
// 	i.SetKeyVal(KeyVal{Key: ":resperson", ValType: "string", ValString: r.ProductManager})
// 	i.SetKeyVal(KeyVal{Key: ":category_lv1", ValType: "string", ValString: r.PrimaryCategory})
// 	i.SetKeyVal(KeyVal{Key: ":category_lv2", ValType: "string", ValString: r.SecondaryCategory})
// 	i.SetKeyVal(KeyVal{Key: ":category_lv3", ValType: "string", ValString: r.TertiaryCategory})
// }

// func (i *SQLProccesor) SetParam_ReqAllPromotion(r promotion.ReqPromotionAllSearch) {
// 	i.SetKeyVal(KeyVal{Key: ":startdate", ValType: "string", ValString: r.StartDate})
// 	i.SetKeyVal(KeyVal{Key: ":enddate", ValType: "string", ValString: r.EndDate})
// 	i.SetKeyVal(KeyVal{Key: ":shop", ValType: "string", ValString: r.ShopFilter})
// 	i.SetKeyVal(KeyVal{Key: ":pallet", ValType: "list", ValString: ToInStringStr(r.CurrentInventory)}) // "('A','B')"
// 	i.SetKeyVal(KeyVal{Key: ":scene", ValType: "list", ValString: ToInStringStr(r.SceneCategory)})
// 	i.SetKeyVal(KeyVal{Key: ":resperson", ValType: "list", ValString: ToInStringStr(r.ProductManager)})
// 	i.SetKeyVal(KeyVal{Key: ":subpallet", ValType: "list", ValString: ToInStringStr(r.Pallet)})
// 	i.SetKeyVal(KeyVal{Key: ":bidtype", ValType: "list", ValString: ToInStringStr(r.BidType)})
// 	i.SetKeyVal(KeyVal{Key: ":keyword", ValType: "list", ValString: ToInStringStr(r.KeywordFilter)})
// 	i.SetKeyVal(KeyVal{Key: ":crowd", ValType: "list", ValString: ToInStringStr(r.AudienceFilter)})

// }

/*
with bpp as (
    select
    bpp.product_id,
    bpc.pallet, bpc.pallet_change, bpc.pre_pallet,
    SUM(returning_buyers_paid) as lmgzfrs,
    SUM(returning_buyers_paid) / SUM(order_placed_buyers) as lkzb,
    SUM(order_amount) as gmv,
    SUM(successful_refund_amount) as refund,
    SUM(bpp.order_placed_buyers) / SUM(bpp.visitors_count) as 支付转化率,
    SUM(bpp.order_quantity) / SUM(bpp.visitors_count) as 收藏率,
    SUM(bpp.add_to_cart_buyers) / SUM(bpp.visitors_count) as 加购率,
    SUM(bpp.order_quantity) / SUM(bpp.order_placed_buyers) as 连带率
    FROM
        biz_product_performance bpp
    JOIN
        biz_product bp ON bp.product_id = bpp.product_id
    JOIN
        biz_product_classes bpc ON bpc.product_id = bpp.product_id and bpc.statistic_date = bpp.statistic_date
WHERE
    (bpc.statistic_date >= :startdate OR :startdate IS NULL OR :startdate = '')
    AND (bpc.statistic_date <= :enddate OR :enddate IS NULL OR :enddate = '')
    AND (bpc.pallet = :pallet OR :pallet IS NULL OR :pallet = '')
    AND (bpc.pallet_change = :pallet_change OR :pallet_change IS NULL)
    AND (bp.responsible = :resperson OR :resperson IS NULL OR :resperson = '')
    AND (bp.category_lv1 = :category_lv1 OR :category_lv1 IS NULL OR :category_lv1 = '')
    AND (bp.category_lv2 = :category_lv2 OR :category_lv2 IS NULL OR :category_lv2 = '')
    AND (bp.category_lv3 = :category_lv3 OR :category_lv3 IS NULL OR :category_lv3 = '')
    group by bpp.product_id,bpc.pallet, bpc.pallet_change, bpc.pre_pallet
),
bpt1 as (
    SELECT
        bpt.product_id,
        SUM(bpt.paid_amount) AS gmv,
        SUM(bpt.visitors_count) AS visitors_count
    FROM
        biz_product_traffic_stats bpt
    JOIN
        biz_product bp ON bp.product_id = bpt.product_id
    JOIN
        biz_product_classes bpc ON bpc.product_id = bpt.product_id AND bpc.statistic_date = bpt.statistic_date
    WHERE
        bpc.statistic_date >= :startdate AND bpc.statistic_date <= :enddate AND bpc.pallet = :pallet AND bp.responsible = :resperson
        AND bpt.source_type_1 = '平台流量' AND bpt.source_type_2 = '汇总' AND bpt.source_type_3 = '汇总'
    GROUP BY
        bpt.product_id
),
bpt2 as (
    SELECT
        bpt.product_id,
        SUM(bpt.paid_amount) AS search_gmv,
        SUM(bpt.visitors_count) AS search_count
    FROM
        biz_product_traffic_stats bpt
    JOIReplace()
        biz_product bp ON bp.product_id = bpt.product_id
    JOIN
        biz_product_classes bpc ON bpc.product_id = bpt.product_id AND bpc.statistic_date = bpt.statistic_date
    WHERE
        bpc.statistic_date >= :startdate AND bpc.statistic_date <= :enddate AND bpc.pallet = :pallet AND bp.responsible = :resperson
        AND bpt.source_type_1 = '平台流量' AND bpt.source_type_2 = '手淘搜索' AND bpt.source_type_3 = '汇总'
    GROUP BY
        bpt.product_id
)
select
    bpp.product_id,
    CASE
        WHEN (bpp.gmv - bpp.refund) > 0 THEN
            ((bpp.gmv - bpp.refund) * bp.estimated_gross_profit_margin
            - bpp.gmv * bp.delivery_cost
            - bp.labor_cost * DATEDIFF(:enddate,:startdate))
            / (bpp.gmv - bpp.refund)
        ELSE 0
    END AS 净利润率,
    CASE
        WHEN bpp.lkzb >= 0.1 AND bpp.lmgzfrs >= 50 AND bpp.gmv >= 5000 THEN '潜力爆品'
        WHEN bpp.lkzb >= 0.1 AND bpp.lmgzfrs >= 50 THEN '回流品'
        WHEN bpp.lkzb >= 0.1 THEN '潜在回流品'
        WHEN bpt2.search_gmv / bpt1.gmv >= 0.3 THEN '刚需品/季节品'
        WHEN bpp.gmv < 5000 THEN '待定品'
        ELSE '拉新品'
    END AS 产品分类,
   CASE
        WHEN bpt1.visitors_count > 0 THEN bpt2.search_count / bpt1.visitors_count
        ELSE 0
    END AS 搜索访客占比,
    CASE
        WHEN bpt1.gmv > 0 THEN bpt2.search_gmv / bpt1.gmv
        ELSE 0
    END AS 搜索GMV占比,
    bpp.lkzb AS 老客占比,
    bpp.pallet AS 本期货盘,
    bpp.pre_pallet AS 上期货盘,
    bpp.pallet_change AS 货盘变化,
    bp.estimated_gross_profit_margin AS 预估毛利率,
    bp.category_lv3 AS 三级类目,
    bp.product_name AS 商品名称
FROM bpp
JOIN
    biz_product bp ON bp.product_id = bpp.product_id
JOIN
    bpt1 ON bpt1.product_id = bpp.product_id
JOIN
    bpt2 ON bpt2.product_id = bpp.product_id
*/
