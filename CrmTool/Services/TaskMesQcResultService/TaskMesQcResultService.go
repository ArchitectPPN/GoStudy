package TaskMesQcResultService

import (
	"CrmTool/Const"
	"CrmTool/Request"
	"CrmTool/Services/QcResult"
	"CrmTool/Services/TaskService"
	"CrmTool/models/Stage"
	"CrmTool/models/TaskInfo"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

// GetTaskQcResultMsg 获取任务单质检结果消息
func GetTaskQcResultMsg(taskRequest Request.GetQcResultMsgRequest) string {
	// 获取taskId
	taskId := TaskService.GetTaskIdByCode(taskRequest.TaskCode)
	if taskId == "" {
		fmt.Println("任务单未查询到, 请确认任务单编号是否正确, TaskCode" + taskRequest.TaskCode)
		return ""
	}
	fmt.Println("查询到的任务单, taskId:", taskId)

	// 获取任务单关联的收货记录
	taskReceiptList := TaskService.GetTaskReceiptByTaskId(taskId)
	if len(taskReceiptList) == 0 {
		fmt.Println("任务单未关联收货记录")
		return ""
	}

	// 获取任务单关联的阶段信息
	_, stageInfo := TaskService.GetTaskRelationStageId(taskId)

	// 组装质检结果
	qcMsg := assembleQcMsg(taskRequest, taskReceiptList, stageInfo)

	marshal, err := json.Marshal(qcMsg)
	if err != nil {
		fmt.Println("转换失败", err)
		return ""
	}

	return string(marshal)
}

// AssembleQcMsg 组装质检结果
func assembleQcMsg(taskRequest Request.GetQcResultMsgRequest, taskReceiptList []*TaskInfo.TaskReceipt, stageInfo Stage.EaStageCstm) QcResult.QcResult {
	fmt.Println("开始组装质检数据")
	modelTask := QcResult.QcResult{WorkOrderId: taskRequest.TaskCode}

	// 模型质检结果
	var ModelQcDetail QcResult.QcDetail

	// 咬合质检结果
	var OcclusionQcDetail *QcResult.QcDetail
	OcclusionQcDetail = &QcResult.QcDetail{}

	// 照片质检结果
	var PhotoQcDetail QcResult.QcDetail

	// 模型上下颌质检结果, 直接初始化上下颌
	var ModelQcItem = make([]QcResult.QcItem, 0)
	ModelQcDetail.QcObjects = ModelQcItem

	// 生成质检结果
	for _, tmpVal := range taskReceiptList {
		fmt.Println("开始处理收货记录", tmpVal, string(rune(Const.MOUSE_SCAN)), string(rune(Const.SILICA)), string(rune(Const.STL)))
		tmReceiptType, err := strconv.Atoi(tmpVal.ReceiptType)
		if err != nil {
			fmt.Println("转换出错了,", err)
			continue
		}

		fmt.Printf("%T %v \n", tmReceiptType, tmReceiptType)

		switch tmReceiptType {
		case Const.STL:
			// 处理咬质检结果
			getQcBiteResult(OcclusionQcDetail, tmpVal, taskRequest, stageInfo)
			// 处理模型质检结果
			getModelQcDetail(&ModelQcDetail, tmpVal, taskRequest)
		case Const.MOUSE_SCAN, Const.SILICA:
			if !taskRequest.IsQcModel {
				// 任务单未质检模型, 退出
				continue
			}
			fmt.Println("匹配到的数据为", tmpVal.ReceiptType)
			getModelQcDetail(&ModelQcDetail, tmpVal, taskRequest)
		case Const.OCCLUSION, Const.DIGITAL_OCCLUSION:
			if !taskRequest.IsQcBite {
				fmt.Println("任务单未质检咬合, taskCode: ", taskRequest.TaskCode)
				// 任务单未质检咬合, 退出
				continue
			}
			// 获取咬合质检结果
			getQcBiteResult(OcclusionQcDetail, tmpVal, taskRequest, stageInfo)
		case Const.PHOTO:
			if !taskRequest.IsQcPhoto {
				// 任务单未质检照片, 退出
				continue
			}
			PhotoQcDetail.QcItem = "photo"
			PhotoQcDetail.QcType = "DIGITAL_QC"
			PhotoQcDetail.PhotoDetail = getQcPhotoResult(taskRequest)
		}
	}

	// 添加模型质检结果
	if taskRequest.IsQcModel {
		modelTask.QcResult = append(modelTask.QcResult, ModelQcDetail)
	}

	// 添加咬合质检结果
	if taskRequest.IsQcBite {
		modelTask.QcResult = append(modelTask.QcResult, *OcclusionQcDetail)
	}

	// 添加照片质检结果
	if taskRequest.IsQcPhoto {
		modelTask.QcResult = append(modelTask.QcResult, PhotoQcDetail)
	}

	return modelTask
}

// getQcBiteResult 获取模型质检结果
func getQcBiteResult(occlusionQcDetail *QcResult.QcDetail, taskReceipt *TaskInfo.TaskReceipt, taskRequest Request.GetQcResultMsgRequest, stageInfo Stage.EaStageCstm) {
	// 咬合质检结果
	biteTaskQcItem := QcResult.QcItem{}

	// 模型类型
	materialType, _ := strconv.Atoi(taskReceipt.ReceiptType)

	// 默认为非A6
	biteTaskQcItem.QcStandard = "notA6"

	// 根据阶段信息来确定是否使用A6/天使杆
	if stageInfo.StageUseA6OrAngelPole == "1" {
		biteTaskQcItem.QcStandard = "A6"
	} else if stageInfo.StageUseA6OrAngelPole == "3" {
		biteTaskQcItem.QcStandard = "AngelArm"
	}

	// 处理质检结果
	if taskRequest.BiteQcResult == Const.VALID {
		biteTaskQcItem.QcResult = Const.VALID
	} else if taskRequest.BiteQcResult == Const.INVALID {
		// 缺陷部位照片的MES下载地址
		occlusionQcDetail.DefectivePartPhotos = []string{"https://mes2-pro.oss-cn-shanghai.aliyuncs.com/2023-04-11/C01004045834/%E4%B8%AD%E6%9C%9F%E9%98%B6%E6%AE%B53/dc22b3c9-ad4b-4a75-a71e-a17ed9197e0d/patient_model_upload/L8P3ZBdUSkoaDPSCVgCz.png"}
		// 缺陷情况描述
		occlusionQcDetail.DefectiveRemark = "该病例无下颌前伸位咬合/缺少下颌前伸位咬合记录。"
		// 不合格报告下载地址
		occlusionQcDetail.MesUnqualifiedReportUrl = "https://mes2-pro.oss-cn-shanghai.aliyuncs.com/2023-04-11/C01004045834/%E4%B8%AD%E6%9C%9F%E9%98%B6%E6%AE%B53/dc22b3c9-ad4b-4a75-a71e-a17ed9197e0d/patient_model_upload/2dbf0525-7dfd-4b54-965f-5319aa296efd.docx"

		// 不合格原因
		occlusionQcDetail.UnqualifiedReason = "无A6咬合"
		// 不合格原因Code
		occlusionQcDetail.UnqualifiedReasonCode = "BITE12"

		biteTaskQcItem.QcResult = Const.INVALID
	}

	// 处理cds/barcode
	switch materialType {
	case Const.DIGITAL_OCCLUSION, Const.STL:
		// 颌位
		biteTaskQcItem.Jaw = taskReceipt.Jaw
		if materialType == Const.DIGITAL_OCCLUSION {
			// 数据咬合默认为上颌
			biteTaskQcItem.Jaw = "U"
		}
		// cdsUrl地址
		biteTaskQcItem.CdsUrl = getCdsUrl(taskReceipt.ReceiptInfo)
	case Const.OCCLUSION:
		biteTaskQcItem.BarCode = taskReceipt.ReceiptInfo
	}

	occlusionQcDetail.QcItem = Const.QC_ITEM_BITE
	occlusionQcDetail.QcType = Const.QC_TYPE
	occlusionQcDetail.QcObjects = append(occlusionQcDetail.QcObjects, biteTaskQcItem)

	fmt.Println("occlusionInfo: ", occlusionQcDetail, "咬合质检任务单: ", biteTaskQcItem)
}

// getModelQcItemResult 获取模型质检结果
func getModelQcItemResult(taskReceipt *TaskInfo.TaskReceipt, taskRequest Request.GetQcResultMsgRequest) QcResult.QcItem {
	modelTaskQcItem := QcResult.QcItem{}

	materialType, _ := strconv.Atoi(taskReceipt.ReceiptType)

	// 1 上颌质检通过 2 下颌质检通过 3 上下颌均质检通过
	modelTaskQcItem.QcResult = Const.VALID
	if taskRequest.ModelQcResult == 3 {
		modelTaskQcItem.QcResult = Const.VALID
	} else if taskRequest.ModelQcResult == 2 && taskReceipt.Jaw == Const.UPPER || taskRequest.ModelQcResult == 1 && taskReceipt.Jaw == Const.LOWER {
		// 模型质检通过
		modelTaskQcItem.QcResult = Const.INVALID
	}

	// 质检标准设置为不是a6标准
	modelTaskQcItem.QcStandard = Const.QC_STANDARD_NOT_A6
	// 设置颌位
	modelTaskQcItem.Jaw = taskReceipt.Jaw

	// 处理cds/barcode
	switch materialType {
	case Const.MOUSE_SCAN, Const.STL:
		runMode := beego.AppConfig.String("runmode")
		switch runMode {
		case "dev":
			modelTaskQcItem.CdsUrl = "http://crm-web.dev.eainc.com/crm/" + "case/document/" + taskReceipt.ReceiptInfo + "/download"
		case "sit2":
			modelTaskQcItem.CdsUrl = "http://crm-web.sit2.eainc.com/crm/" + "case/document/" + taskReceipt.ReceiptInfo + "/download"
		case "sit":
			modelTaskQcItem.CdsUrl = "http://crm-web.sit.eainc.com/crm/" + "case/document/" + taskReceipt.ReceiptInfo + "/download"
		default:
			modelTaskQcItem.CdsUrl = "http://crm-web.dev.eainc.com/crm/" + "case/document/" + taskReceipt.ReceiptInfo + "/download"
		}
	case Const.SILICA:
		modelTaskQcItem.BarCode = taskReceipt.ReceiptInfo
	}

	return modelTaskQcItem
}

// getQcPhotoResult 获取照片质检结果
func getQcPhotoResult(taskRequest Request.GetQcResultMsgRequest) *QcResult.PhotoDetail {
	// 声明照片质检结果
	photoQcResult := QcResult.PhotoDetail{}

	// 照片质检通过
	if taskRequest.PhotoQcResult == Const.VALID {
		photoQcResult.QcResult = Const.VALID
	} else {
		photoQcResult.QcResult = Const.INVALID
	}

	// 添加不合格原因
	if photoQcResult.QcResult == Const.INVALID {
		photoQcResult.MesUnqualifiedReportUrl = "https://mes-sit.oss-cn-shanghai.aliyuncs.com/2023-08-17/C01001789331/%E6%96%B0%E7%97%85%E4%BE%8B%E9%98%B6%E6%AE%B5/91de1d39-9e34-42fb-a18c-99a180ff0df6/patient_model_upload/b9aec052-bfd0-4013-82a8-d5ec83827ff7.docx"
	}

	var photoQcDetailObject [8]QcResult.PhotoQcDetailObject
	// 照片每个质检项

	if photoQcResult.QcResult == Const.INVALID {
		photoQcDetailObject[0] = QcResult.PhotoQcDetailObject{
			Attachs:             []string{"https://mes-sit.oss-cn-shanghai.aliyuncs.com/2023-04-03/C01003990045/%E6%96%B0%E7%97%85%E4%BE%8B%E9%98%B6%E6%AE%B5/37bd2aad-5f06-4510-8f36-1c91f03f3298/patient_model_upload/jSZpv7MW224h4GWusew5.jpg"},
			DefectiveReason:     "耳廓后缘未拍全",
			DefectiveReasonCode: "PHOTO16",
			DefectiveRemark:     "耳廓后缘未拍全",
			Position:            "1",
		}
	} else {
		photoQcDetailObject[0] = QcResult.PhotoQcDetailObject{Position: "1"}
	}

	photoQcDetailObject[1] = QcResult.PhotoQcDetailObject{Position: "2"}
	photoQcDetailObject[2] = QcResult.PhotoQcDetailObject{Position: "3"}
	photoQcDetailObject[3] = QcResult.PhotoQcDetailObject{Position: "4"}
	photoQcDetailObject[4] = QcResult.PhotoQcDetailObject{Position: "5"}
	photoQcDetailObject[5] = QcResult.PhotoQcDetailObject{Position: "6"}
	photoQcDetailObject[6] = QcResult.PhotoQcDetailObject{Position: "7"}
	photoQcDetailObject[7] = QcResult.PhotoQcDetailObject{Position: "8"}

	photoQcResult.PhotoDefectiveDetails = photoQcDetailObject

	return &photoQcResult
}

// getModelQcDetail 组装模型质检结果
func getModelQcDetail(modelQcDetail *QcResult.QcDetail, taskReceipt *TaskInfo.TaskReceipt, taskRequest Request.GetQcResultMsgRequest) {

	modelQcDetail.QcItem = "model"
	modelQcDetail.QcType = "DIGITAL_QC"
	modelQcDetail.CouldBuccinator = 3
	modelQcDetail.QcObjects = append(modelQcDetail.QcObjects, getModelQcItemResult(taskReceipt, taskRequest))

	// 模型质检不通过时, 添加不通过报告 3 全颌
	if taskRequest.ModelQcResult != 3 {
		modelQcDetail.DefectivePartPhotos = []string{"https://mes-sit.oss-cn-shanghai.aliyuncs.com/2023-08-16/C01001792155/%E6%96%B0%E7%97%85%E4%BE%8B%E9%98%B6%E6%AE%B5/00ffe14b-588c-4b39-ba91-ad587f7f3539/patient_model_upload/0kmUATOSuS9zkJ5w4FXS.jpg"}
		modelQcDetail.DefectiveRemark = time.Now().Format("2006-01-02 15:04:05") + "!!! Begoo生成 !!! 我是不通过描述!"
		modelQcDetail.MesUnqualifiedReportUrl = "https://mes-sit.oss-cn-shanghai.aliyuncs.com/2023-08-16/C01001792155/%E6%96%B0%E7%97%85%E4%BE%8B%E9%98%B6%E6%AE%B5/00ffe14b-588c-4b39-ba91-ad587f7f3539/patient_model_upload/ab32f8a7-5bde-44f2-b9cd-ef6b1a6168da.docx"
		modelQcDetail.UnqualifiedReason = "硅胶断裂、断开 ---- " + time.Now().Format("2006-01-02 15:04:05") + "!!! Begoo 生成 !!! "
		modelQcDetail.UnqualifiedReasonCode = "MODEL04"
	}

	// 下颌前伸位 咬合质检不合格时,模型质检结果中添加不合格报告
	if taskRequest.BiteQcResult == Const.INVALID {
		modelQcDetail.MesUnqualifiedReportUrl = "https://mes2-pro.oss-cn-shanghai.aliyuncs.com/2023-04-11/C01004045834/%E4%B8%AD%E6%9C%9F%E9%98%B6%E6%AE%B53/dc22b3c9-ad4b-4a75-a71e-a17ed9197e0d/patient_model_upload/2dbf0525-7dfd-4b54-965f-5319aa296efd.docx"
	}

	// 处理模型质检结果
	var RiskDetail QcResult.RiskDetail
	RiskDetail.PromptRisk = 0

	fmt.Println("模型有风险:", taskRequest.ModelHasRisk)
	// 模型有风险
	if taskRequest.ModelHasRisk {
		// 风险描述
		RiskDetail.RiskDesc = "上颌17萌出部分未扫全。"
		RiskDetail.PromptRisk = 1
		// 风险照片描述
		RiskDetail.RiskDescPhotos = []string{"https://mes2-pro.oss-cn-shanghai.aliyuncs.com/2023-04-11/C01004045834/%E4%B8%AD%E6%9C%9F%E9%98%B6%E6%AE%B53/dc22b3c9-ad4b-4a75-a71e-a17ed9197e0d/patient_model_upload/JLBfa1JzkVu2jVBZxK01.png"}
		// 风险报告
		RiskDetail.RiskReportUrl = "https://mes2-pro.oss-cn-shanghai.aliyuncs.com/2023-04-11/C01004045834/%E4%B8%AD%E6%9C%9F%E9%98%B6%E6%AE%B53/dc22b3c9-ad4b-4a75-a71e-a17ed9197e0d/patient_model_upload/e2f0ff58-c4d8-4643-a437-0112a0adab61.docx"
	}

	modelQcDetail.RiskDetail = &RiskDetail
}

// getCdsUrl 获取cdsUrl
func getCdsUrl(cdsUuid string) string {
	runMode := beego.AppConfig.String("runmode")
	switch runMode {
	case "dev":
		return "http://crm-web.dev.eainc.com/crm/" + "case/document/" + cdsUuid + "/download"
	case "sit2":
		return "http://crm-web.sit2.eainc.com/crm/" + "case/document/" + cdsUuid + "/download"
	case "sit":
		return "http://crm-web.sit.eainc.com/crm/" + "case/document/" + cdsUuid + "/download"
	default:
		return "http://crm-web.dev.eainc.com/crm/" + "case/document/" + cdsUuid + "/download"
	}
}
