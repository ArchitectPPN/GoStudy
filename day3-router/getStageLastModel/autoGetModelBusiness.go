package getStageLastModel

import (
	"database/sql"
	"example/mysqlDb"
	"fmt"
)

type stageLastModel struct {
}

type Receipt struct {
	ReceiptId    string
	MaterialType sql.NullInt16
	Qualified    sql.NullInt16
	GoNext       sql.NullInt16
	BindTime     sql.NullString
	UCdsUrl      sql.NullString
	LCdsUrl      sql.NullString
}

type StageAllReceipt struct {
	NoResReceipt    map[string]Receipt
	QcPassReceipt   map[string]Receipt
	UnQcPassReceipt map[string]Receipt
}

func (stageAllReceipt *stageLastModel) getStageReceipt(stageId string) (returnAllReceipt StageAllReceipt) {
	// 获取阶段下所有的收货记录
	receiptList := make(map[string]Receipt, 10)
	receiptList = stageAllReceipt.getAllReceiptByStageId(stageId)

	returnAllReceipt = stageAllReceipt.dealStageLastModel(receiptList)

	return returnAllReceipt
}

func (stageAllReceipt *stageLastModel) dealStageLastModel(receiptList map[string]Receipt) StageAllReceipt {
	var stageAllReceiptList StageAllReceipt
	stageAllReceiptList.QcPassReceipt = make(map[string]Receipt, 10)
	stageAllReceiptList.NoResReceipt = make(map[string]Receipt, 10)
	stageAllReceiptList.UnQcPassReceipt = make(map[string]Receipt, 10)

	// 根据收货记录类型不同, 做不同的处理
	fmt.Println("阶段--- 开始")
	for _, receiptVal := range receiptList {
		switch receiptVal.MaterialType.Int16 {
		case 14:
			stageAllReceiptList = stageAllReceipt.mouseScanDeal(receiptVal, stageAllReceiptList)
		case 15:
			stageAllReceiptList = stageAllReceipt.silicaDeal(receiptVal, stageAllReceiptList)
		}
	}

	return stageAllReceiptList
}

func (stageAllReceipt *stageLastModel) mouseScanDeal(receipt Receipt, stageAllReceiptList StageAllReceipt) StageAllReceipt {
	if receipt.GoNext.Int16 == 1 && receipt.Qualified.Int16 != 1 {
		// 质检合格
		stageAllReceiptList.QcPassReceipt[receipt.ReceiptId] = receipt
	} else if receipt.Qualified.Int16 == 2 {
		// 质检不通过
		stageAllReceiptList.UnQcPassReceipt[receipt.ReceiptId] = receipt
	} else {
		// 未质检
		stageAllReceiptList.NoResReceipt[receipt.ReceiptId] = receipt
	}

	return stageAllReceiptList
}

func (stageAllReceipt *stageLastModel) silicaDeal(receipt Receipt, stageAllReceiptList StageAllReceipt) StageAllReceipt {
	if receipt.GoNext.Int16 == 1 && receipt.Qualified.Int16 != 1 {
		// 质检合格
		stageAllReceiptList.QcPassReceipt[receipt.ReceiptId] = receipt
	} else if receipt.Qualified.Int16 == 2 {
		// 质检不通过
		stageAllReceiptList.QcPassReceipt[receipt.ReceiptId] = receipt
	} else {
		// 未质检
		stageAllReceiptList.QcPassReceipt[receipt.ReceiptId] = receipt
	}

	return stageAllReceiptList
}

func (stageAllReceipt *stageLastModel) getAllReceiptByStageId(stageId string) (receiptList map[string]Receipt) {
	// 获取阶段下所有的收货记录
	db := mysqlDb.InitDB()

	sqlStr := "SELECT erc.id_c as receiptId, erc.material_type_c, erc.is_qualified_c, erc.go_next_c, erc.barcode_u_status_c, erc.barcode_l_status_c, eser1c.date_modified " +
		" FROM ea_stage_ea_receipt_1_c eser1c " +
		"         INNER JOIN ea_receipt er ON er.id = eser1c.ea_stage_ea_receipt_1ea_receipt_idb AND er.deleted = 0" +
		"         INNER JOIN ea_receipt_cstm erc ON erc.id_c = er.id AND erc.material_type_c IN (14, 21, 15)" +
		" WHERE eser1c.ea_stage_ea_receipt_1ea_stage_ida = ?"

	rows, err := db.Query(sqlStr, stageId)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	defer rows.Close()
	receiptList = make(map[string]Receipt, 10)
	for rows.Next() {
		var receipt Receipt
		err := rows.Scan(&receipt.ReceiptId, &receipt.MaterialType, &receipt.Qualified, &receipt.GoNext, &receipt.UCdsUrl, &receipt.LCdsUrl, &receipt.BindTime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		receiptList[receipt.ReceiptId] = receipt
	}

	return receiptList
}

func (stageAllReceipt *stageLastModel) New() *StageAllReceipt {
	return &StageAllReceipt{}
}
