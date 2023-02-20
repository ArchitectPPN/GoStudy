package BusinessOrderGoProductionScan

import (
	"database/sql"
	"fmt"
)

func GetCaseInfoByCaseId(dbObj *sql.DB, caseId string) CaseInfo {
	caseIdSql := "SELECT " +
		"ecc.id_c,ec.name,ecc.is_teenager_c,ecc.state_c " +
		"FROM ea_case ec " +
		"INNER JOIN ea_case_cstm ecc ON ecc.id_c = ec.id " +
		"WHERE ecc.id_c = ? "

	var caseInfo CaseInfo
	err := dbObj.QueryRow(caseIdSql, caseId).Scan(&caseInfo.CaseId, &caseInfo.CaseCode, &caseInfo.TeenagerInfo, &caseInfo.CaseState)
	if err != nil {
		fmt.Println("查询出错：", err)
		return caseInfo
	}

	return caseInfo
}

func GetBusinessOrderList(dbObj *sql.DB, limit int) map[int]string {
	getBusinessOrderSql := "SELECT id FROM ea_businessorder WHERE deleted = 0 LIMIT ?, 100"
	rows, err := dbObj.Query(getBusinessOrderSql, limit)

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return make(map[int]string, 20)
	}

	defer rows.Close()

	businessOrderIdList := make(map[int]string, 20)
	index := 0
	var businessOrderId string
	for rows.Next() {
		err := rows.Scan(&businessOrderId)
		if err != nil {
			fmt.Printf("scan faled, err:%v \n", err)
		}

		businessOrderIdList[index] = businessOrderId
		index++
	}

	return businessOrderIdList
}

func GetBusinessOrderGetBusinessOrderInfo(dbObj *sql.DB, businessOrderId string) (businessOrderInfo BusinessOrderInfo) {
	getBusinessOrderInfoSql := "SELECT code_c as BusinessOrderCode, status_c as businessOrderStatus " +
		"FROM ea_businessorder eb " +
		"JOIN ea_businessorder_cstm ebc ON eb.id = ebc.id_c " +
		"WHERE eb.deleted = 0 " +
		"AND ebc.id_c = ?"

	_ = dbObj.QueryRow(getBusinessOrderInfoSql, businessOrderId).Scan(&businessOrderInfo.BusinessOrderCode, &businessOrderInfo.BusinessOrderStatus)

	return
}

func GetBusinessOrderExtInfo(dbObj *sql.DB, businessOrderId string) (businessOrderExtInfo BusinessOrderExtInfo) {
	getBusinessOrderExtInfoSql := "SELECT value FROM ea_business_info_ext WHERE parent_id = ? AND field = 'requireJaw'"

	_ = dbObj.QueryRow(getBusinessOrderExtInfoSql, businessOrderId).Scan(&businessOrderExtInfo.requiredJaw)

	return
}

func GetBusinessRelationReceipt(dbObj *sql.DB, businessOrderId string) map[int]BusinessOrderReceipt {
	getBusinessOrderReceiptSql := "SELECT related_id, jaw, receipt_type, qualification, in_use, go_next FROM ea_businessorder_ea_receipt_relation WHERE outward_order_id = ?"
	rows, err := dbObj.Query(getBusinessOrderReceiptSql, businessOrderId)

	receiptList := make(map[int]BusinessOrderReceipt)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	var businessOrderReceipt BusinessOrderReceipt
	index := 0
	for rows.Next() {
		_ = rows.Scan(&businessOrderReceipt.receiptId, &businessOrderReceipt.receiptJaw, &businessOrderReceipt.receiptType, &businessOrderReceipt.qualification, &businessOrderReceipt.inUse, &businessOrderReceipt.goNext)
		receiptList[index] = businessOrderReceipt
		index++
	}

	return receiptList
}

// 复购的订单和被复购的订单, 患者不是同一个, 根据 ea_patients_ea_businessorder_1_c

func GetBusinessOrderProduction(dbObj *sql.DB, businessOrderId string) string {
	getBusinessOrderProductionSql := " SELECT ep.id FROM ea_businessorder_ea_production_1_c ebep1c " +
		" INNER JOIN ea_production ep ON ep.id = ebep1c.ea_businessorder_ea_production_1ea_production_idb AND ep.deleted = 0 " +
		" INNER JOIN ea_production_cstm epm ON epm.id_c = ep.id " +
		" WHERE epm.state_c != '5' AND ebep1c.ea_businessorder_ea_production_1ea_businessorder_ida = ? "

	productionId := ""
	_ = dbObj.QueryRow(getBusinessOrderProductionSql, businessOrderId).Scan(&productionId)

	return productionId
}
