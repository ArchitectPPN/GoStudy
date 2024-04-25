package UploadStl

import (
	"CrmTool/Const"
	"CrmTool/Request"
	"CrmTool/Services/TaskService"
	"CrmTool/StructDefined/MesCrmUploadStl"
	"CrmTool/models/TaskInfo"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func GetUploadStlResult(taskRequest Request.GetQcResultMsgRequest) string {
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

	uploadRes := assembleUploadResultStl(taskRequest, taskReceiptList)

	uploadResStr, err := json.Marshal(uploadRes)
	if err != nil {
		fmt.Println("mes 上传slt地址 json 失败, 原因:", err)
		return ""
	}

	return string(uploadResStr)
}

func assembleUploadResultStl(taskRequest Request.GetQcResultMsgRequest, taskReceiptList []*TaskInfo.TaskReceipt) MesCrmUploadStl.MesCrmUploadStl {
	// 初始化MesCrmUploadStl并设置taskCode
	mesUploadStlRes := MesCrmUploadStl.MesCrmUploadStl{WorkOrderId: taskRequest.TaskCode}

	// 开始循环写入stl地址
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
		case Const.SILICA:
			tempDetail := MesCrmUploadStl.Detail{}
			tempDetail.Jaw = tmpVal.Jaw
			tempDetail.CdsUrl = getCdsUrl()
			tempDetail.GenerateFromRemade = false
			tempDetail.BarCode = tmpVal.ReceiptInfo
			mesUploadStlRes.UploadResult = append(mesUploadStlRes.UploadResult, tempDetail)
		case Const.OCCLUSION:
			upperTempDetail := MesCrmUploadStl.Detail{}
			lowerTempDetail := MesCrmUploadStl.Detail{}

			upperTempDetail.Jaw = Const.UPPER
			upperTempDetail.CdsUrl = getCdsUrl()
			upperTempDetail.GenerateFromRemade = false
			upperTempDetail.BarCode = tmpVal.ReceiptInfo

			lowerTempDetail.Jaw = Const.LOWER
			lowerTempDetail.CdsUrl = getCdsUrl()
			lowerTempDetail.GenerateFromRemade = false
			lowerTempDetail.BarCode = tmpVal.ReceiptInfo

			mesUploadStlRes.UploadResult = append(mesUploadStlRes.UploadResult, upperTempDetail)
			mesUploadStlRes.UploadResult = append(mesUploadStlRes.UploadResult, lowerTempDetail)

		}
	}

	return mesUploadStlRes
}

// getCdsUrl 获取cdsUrl
func getCdsUrl() string {
	cdsUrl := uuid.New()

	runMode := beego.AppConfig.String("runmode")
	switch runMode {
	case "dev":
		return "http://cds.dev.eainc.com:8001/shidaits/rest/cds/" + "case/document/" + cdsUrl.String() + "/download"
	case "sit2":
		return "http://cds.sit2.eainc.com:8001/shidaits/rest/cds/" + "case/document/" + cdsUrl.String() + "/download"
	case "sit":
		return "http://cds.sit.eainc.com:8001/shidaits/rest/cds/" + "case/document/" + cdsUrl.String() + "/download"
	default:
		return "http://cds.dev.eainc.com:8001/shidaits/rest/cds/" + "case/document/" + cdsUrl.String() + "/download"
	}
}

func randStringBytesRmndr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
