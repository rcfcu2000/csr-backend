package product

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespSku) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_Sku)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespSku sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespSku sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespSku) Prodata() {

}

const SQL_PRODUCT_Sku = `

select v.sku_id,       -- v.statistic_date as "date",
v.sku_name as "sku_name", --  sku_name
sum(v.payment_amount) as "pay_amount",  -- 支付金额
sum(v.buyer_count) as "pay_buyers",  -- 支付买家数
sum(v.item_sold_count) as "pay_quantity",  -- 支付件数
sum(v.add_to_cart_item_count) as "add_to_cart_count"   -- 加购件数

from v_bps_bp v
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
GROUP BY v.sku_id, sku_name
order by  v.sku_id 

`
