package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPalletTargetInfoListData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_PalletTargetInfoList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespPalletTargetListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPalletTargetListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}

func (i *RespPalletTargetInfoListData) Prodata() {

}

const SQL_TARGET_PalletTargetInfoList = `
WITH t_id_pallet_gmv AS (
	SELECT
	sum(bpp.paid_amount) AS gmv,
	sum(bpp.spend) AS spend,
	sum(bpp.spend) /sum(bpp.paid_amount)   as "promotion_percentage",
	sum(bpp.paid_amount) / sum(bpp.spend)  as "composite_roi", 
	sum(bpp.paid_amount) / sum(bpp.buyer_count)  as "customer_unit_price", 
	sum(bpp.jlr)  as "profit", 
	sum(bpp.buyer_count)  as "paid_buyers", 

		max(
			CASE
			WHEN bpp.statistic_date = :enddate  THEN
				bpp.pallet
			ELSE
				'-'
			END
		) AS pallet
	FROM
		biz_pallet_product bpp
	WHERE
		bpp.statistic_date >= :startdate 
	AND bpp.statistic_date <= :enddate 
	GROUP BY
		pallet
),
 t_pallet_target_gmv AS (
	SELECT
		pallet,
		sum(gmv_target) AS gmv_target,
		sum(monthly_budget) AS monthly_budget -- 花费月预算
	FROM
		biz_shop_pallet_targets bspt
	WHERE
		bspt.statistic_date >= :startdate 
	AND bspt.statistic_date <= LAST_DAY(:enddate )
	GROUP BY
		pallet
) 
 

SELECT
	tmain.pallet,
	gmv ,
	spend  AS spend,
	gmv_target as month_gmv,
	gmv_target as target_gmv ,
	(gmv / IFNULL(gmv_target, 0)) AS target_gmv_rate,
			promotion_percentage   as "promotion_percentage",
			composite_roi  as "composite_roi", 
			customer_unit_price  as "customer_unit_price", 
			profit  as "profit", 
			paid_buyers  as "paid_buyers", 
			monthly_budget as "monthly_budget", 
			monthly_budget / gmv_target as "promotion_target_percentage",
			 spend - monthly_budget  as "promotion_diff",			
  ( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  time_schedule
 
FROM
	t_id_pallet_gmv tmain
LEFT JOIN t_pallet_target_gmv bspt ON bspt.pallet = tmain.pallet 
`
