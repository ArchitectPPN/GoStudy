package QcResult

import (
	"github.com/google/uuid"
)

func createModelQcRes(taskInfo TaskInfo) []QcDetail {
	var QcObjects []QcDetail
	// 模型质检结果
	var QcModelDetail QcDetail

	var qcItem QcItem
	for _, value := range taskInfo.TaskReceiptDetail {

		if value.ReceiptType == "14" || value.ReceiptType == "15" || value.ReceiptType == "21" {
			forModel(value, &qcItem)
			qcItem.QcStandard = taskInfo.BiteQcType
			QcModelDetail.QcObjects = append(QcModelDetail.QcObjects, qcItem)
		}
	}

	// 填入
	QcModelDetail.RiskDetail = RiskDetail{PromptRisk: 0}
	QcModelDetail.QcType = "DIGITAL_QC"
	QcModelDetail.CouldBuccinator = 1
	QcModelDetail.QcItem = "model"

	QcObjects = append(QcObjects, QcModelDetail)

	return QcObjects
}

func forModel(receiptDetail TaskReceiptInfo, qcItem *QcItem) {
	if receiptDetail.ReceiptType == "15" {
		// 设置barcode
		qcItem.BarCode = receiptDetail.ReceiptInfo.String
	} else {
		qcItem.CdsUrl = getCdsUrl(receiptDetail.ReceiptInfo.String)
	}

	// 设置质检结果
	qcItem.QcResult = 1
	// 设置颌位
	qcItem.Jaw = receiptDetail.Jaw.String
}

func getCdsUrl(cdsUuid string) string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return ""
	}

	if cdsUuid != "" {
		return "https://cds.dev.eainc.com:8001/shidaits/rest/cds/case/document/" + cdsUuid + "/download"
	}

	return "https://cds.dev.eainc.com:8001/shidaits/rest/cds/case/document/" + newUUID.String() + "/download"
}
