package getStageLastModel

import (
	"fmt"
)

type halfModel struct {
	ReceiptId    string
	MaterialType int
	CdsUrl       string
	Jaw          string
	BarCode      string
	Qualified    int
}

type StageLastModel struct {
	Upper         halfModel
	Lower         halfModel
	CaseId        string
	StageId       string
	ReceiptId     string
	MaterialType  int
	StageHasModel bool
	Qualified     int
}

func GetStageLastModel() (stageModel StageLastModel) {
	getStageLastModel := stageLastModel{}
	stageAllReceipt := getStageLastModel.getStageReceipt("b8fa035c-e68b-925a-03da-635f70c22bb0")

	stageModel = compareReceiptBindTime(stageAllReceipt)

	fmt.Println(stageModel)

	return
}

func compareReceiptBindTime(stageAllReceipt StageAllReceipt) StageLastModel {
	fmt.Println("输出质检通过的收货记录列表 开始")
	fmt.Println(stageAllReceipt.QcPassReceipt)
	fmt.Println("输出质检通过的收货记录列表 结束")

	var receipt Receipt

	for _, listDetail := range stageAllReceipt.QcPassReceipt {
		if receipt.BindTime.String == "" && listDetail.BindTime.String != "" {
			receipt = listDetail
		}
	}

	return packStageLastModel(receipt)
}

func packStageLastModel(receiptInfo Receipt) StageLastModel {
	var stageLastModel StageLastModel
	var upperModel halfModel
	var lowerModel halfModel

	// 阶段是否有模型
	stageLastModel.StageHasModel = true

	// 阶段Id
	stageLastModel.StageId = ""
	// 病例Id
	stageLastModel.CaseId = ""

	// 收货记录Id
	stageLastModel.ReceiptId = ""

	// 设置上下颌数据
	stageLastModel.Lower = lowerModel
	stageLastModel.Upper = upperModel

	// 设置收货记录类型
	stageLastModel.MaterialType = 14

	// 质检结果
	stageLastModel.Qualified = 1

	return stageLastModel
}
