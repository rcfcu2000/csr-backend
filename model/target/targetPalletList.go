package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPalletTargetListData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_PalletList)
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
func (i *RespPalletTargetListData) Prodata() {

}

const SQL_TARGET_PalletList = `
WITH t_id_pallet_gmv AS (
	SELECT
		sum(bpp.paid_amount) AS gmv,
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
		sum(gmv_target) AS gmv_target
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
	gmv,
	gmv_target as month_gmv,
	gmv_target as target_gmv,
  (gmv / IFNULL(gmv_target, 0)) AS target_gmv_rate,
  ( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  target_day_rate
 
FROM
	t_id_pallet_gmv tmain
LEFT JOIN t_pallet_target_gmv bspt ON bspt.pallet = tmain.pallet 


	
`
