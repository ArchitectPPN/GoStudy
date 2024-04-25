package ScanReceiptNoStl

import (
	"database/sql"
	"fmt"
)

func getStageList(db *sql.DB, startIndex int) [10000]string {
	receiptListSql := `
	SELECT id FROM ea_stage WHERE deleted = 0 ORDER BY id LIMIT ?, 10000
`
	rows, err := db.Query(receiptListSql, startIndex)

	defer rows.Close()

	var stageList [10000]string
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return stageList
	}

	index := 0
	for rows.Next() {
		err := rows.Scan(&stageList[index])
		if err != nil {
			fmt.Printf("scan faled, err:%v \n", err)
		}
		index++
	}

	return stageList
}

func getStageInfo(db *sql.DB, stageId string) stageInfo {
	stageInfoSql := `
	SELECT id_c, is_produce_c FROM ea_stage_cstm WHERE id_c = ?
`
	var stageDetail stageInfo
	err := db.QueryRow(stageInfoSql, stageId).Scan(&stageDetail.id, &stageDetail.produce)
	if err != nil {
		return stageDetail
	}

	return stageDetail
}

// 获取阶段下的收货记录列表
func getStageReceiptList(db *sql.DB, stageId string) map[int]stageReceipt {
	stageReceiptListSql := `
	SELECT 
	    erc.id_c, 
	    erc.material_type_c, 
	    erc.is_qualified_c, 
	    erc.barcode_u_status_c, 
	    erc.barcode_l_status_c,
	    erc.barcode_u_c,
	    erc.barcode_l_c
	FROM ea_stage_ea_receipt_1_c eser1c 
	    INNER JOIN ea_receipt er ON er.id = eser1c.ea_stage_ea_receipt_1ea_receipt_idb AND eser1c.deleted = 0
	    INNER JOIN ea_receipt_cstm erc ON er.id = erc.id_c AND er.deleted = 0
	WHERE
	    erc.material_type_c IN ('18', '15') 
	  AND erc.is_qualified_c = 1 
	  AND eser1c.ea_stage_ea_receipt_1ea_stage_ida = ?
`
	rows, err := db.Query(stageReceiptListSql, stageId)

	defer rows.Close()

	stageReceiptList := make(map[int]stageReceipt, 20)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return stageReceiptList
	}

	index := 0
	stageReceipt := stageReceipt{}
	for rows.Next() {
		err := rows.Scan(&stageReceipt.ReceiptId, &stageReceipt.MaterialType, &stageReceipt.Qualified, &stageReceipt.upperCdsStl, &stageReceipt.lowerCdsStl, &stageReceipt.upperBarCode, &stageReceipt.lowerBarCode)
		if err != nil {
			fmt.Printf("scan faled, err:%v \n", err)
		}
		stageReceiptList[index] = stageReceipt
		index++
	}

	return stageReceiptList
}
