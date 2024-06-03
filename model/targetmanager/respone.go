package targetmanager

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

type RespPalletTargetListData struct {
	Records []PalletTarget `json:"records"`
}

func (i *RespPalletTargetListData) GetData(r ReqPalletTargetSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_PalletList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_TARGET_PalletList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_TARGET_PalletList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespPalletTargetListData) Prodata() {

}

const SQL_TARGET_PalletList = `
 
SELECT 
*  
FROM
	biz_shop_pallet_targets
where deleted_at =0   and pallet in :pallet 
and statistic_date 	>=:startdate  and  statistic_date <=:enddate  
LIMIT :offset , :pageSize  
	
`

///////////////////////////////////////////////////////////////////////////////
//产品目标

type RespProductTargetListData struct {
	Records []ProductTarget `json:"records"`
}

func (i *RespProductTargetListData) GetData(r ReqProductTargetSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_ProductList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespProductTargetListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductTargetListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespProductTargetListData) Prodata() {

}

const SQL_TARGET_ProductList = `
 
SELECT 
*  
FROM
	biz_shop_pallet_targets
where deleted_at =0   -- and pallet in :pallet 
and statistic_date 	>=:startdate  and  statistic_date <=:enddate  
LIMIT :offset , :pageSize  
	
`
