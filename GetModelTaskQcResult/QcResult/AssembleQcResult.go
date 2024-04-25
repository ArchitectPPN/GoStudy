package QcResult

func GetExampleQcResult(info TaskInfo) QcResult {
	modelQcResult := QcResult{}

	// 质检结果
	//{BarCode: "GJ20220102369", CdsUrl: "http:\\www.baidu.com", Jaw: "U", QcResult: 1, QcStandard: "A6"}, {BarCode: "GJ20220102370", CdsUrl: "http:\\www.baidu.com", Jaw: "L", QcResult: 1, QcStandard: "A6"}
	//var qcItem QcItem
	//for i := 0; i <= 1; i++ {
	//	qcItem.QcResult = 1
	//	if i == 0 {
	//		qcItem.Jaw = "U"
	//	} else {
	//		qcItem.Jaw = "L"
	//	}
	//	qcItem.CdsUrl = getCdsUrl()
	//	qcItem.BarCode = "barCode" + string(rune(i))
	//	QcObjects[i] = qcItem
	//}
	//
	//qcDetail := QcDetail{CouldBuccinator: 1, QcItem: "model", QcType: "DIGITAL_QC", RiskDetail: RiskDetail{PromptRisk: 0}}
	//qcDetail.QcObjects = append(qcDetail.QcObjects, QcObjects...)
	////
	//modelQcResult.QcResult = append(modelQcResult.QcResult, qcDetail)
	// 设置质检工单id
	modelQcResult.WorkOrderId = info.TaskName

	// 等于质检
	if info.WorkType == "qc" {
		modelQcResult.QcResult = createModelQcRes(info)
	}

	return modelQcResult
}
