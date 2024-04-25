package TaskController

import (
	"CrmTool/Const"
	"CrmTool/Request"
	"CrmTool/Services/KafkaService"
	"CrmTool/Services/TaskMesQcResultService"
	"CrmTool/Services/TaskService"
	"CrmTool/Services/UploadStl"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type TaskController struct {
	beego.Controller
	TaskRequest Request.GetQcResultMsgRequest
}

// initTaskExternal 初始化任务单质检信息
func (c *TaskController) initTaskExternal() {
	// 获取任务单id
	taskId := TaskService.GetTaskIdByCode(c.TaskRequest.TaskCode)

	fmt.Println("TaskId:", taskId)

	// 获取任务单质检
	taskExternalInfo := TaskService.GetTaskExternalByTaskId(taskId)

	dataJson := taskExternalInfo.QcItem.String
	var arr []string
	_ = json.Unmarshal([]byte(dataJson), &arr)
	fmt.Println("Unmarshaled: ", arr)

	c.TaskRequest.TaskId = taskId

	for _, value := range arr {
		switch value {
		case Const.QC_ITEM_BITE:
			c.TaskRequest.IsQcBite = true
		case Const.QC_ITEM_MODEL:
			c.TaskRequest.IsQcModel = true
		case Const.QC_ITEM_PHOTO:
			c.TaskRequest.IsQcPhoto = true
		}
	}

	fmt.Println(c.TaskRequest)
}

// init 初始化请求参数
func (c *TaskController) init() {
	// 任务单编号
	c.TaskRequest.TaskCode = c.Ctx.Input.Query("task_code")

	// 模型质检结果
	modelQcRes, err := strconv.Atoi(c.Ctx.Input.Query("model_qc_res"))
	if err != nil {
		c.TaskRequest.ModelQcResult = 3
	} else {
		switch modelQcRes {
		case 1:
			c.TaskRequest.ModelQcResult = 1
		case 2:
			c.TaskRequest.ModelQcResult = 2
		default:
			c.TaskRequest.ModelQcResult = 3
		}
	}

	// 咬合质检结果
	taskBiteQcResult, err := strconv.Atoi(c.Ctx.Input.Query("bite_qc_res"))
	if err != nil {
		c.TaskRequest.BiteQcResult = Const.VALID
	} else {
		c.TaskRequest.BiteQcResult = taskBiteQcResult
	}

	// 恒牙萌出质检
	permanentEruption, err := strconv.Atoi(c.Ctx.Input.Query("permanent_eruption"))
	if err != nil {
		c.TaskRequest.PermanentEruption = Const.PERMANENT_ERUPTION_NO
	} else {
		switch permanentEruption {
		case 1:
			c.TaskRequest.PermanentEruption = Const.PERMANENT_ERUPTION_YES
		case 2:
			c.TaskRequest.PermanentEruption = Const.PERMANENT_ERUPTION_NO
		default:
			c.TaskRequest.PermanentEruption = Const.PERMANENT_ERUPTION_DEFAULT
		}
	}

	// 质检唇颊肌
	couldBuccinator, err := strconv.Atoi(c.Ctx.Input.Query("could_buccinator"))
	if err != nil {
		c.TaskRequest.CouldBuccinator = Const.COULD_BUCCINATOR_DEFAULT
	} else {
		switch couldBuccinator {
		case 1:
			c.TaskRequest.CouldBuccinator = Const.COULD_BUCCINATOR_YES
		case 2:
			c.TaskRequest.CouldBuccinator = Const.COULD_BUCCINATOR_NO
		default:
			c.TaskRequest.CouldBuccinator = Const.COULD_BUCCINATOR_DEFAULT
		}
	}

	// 照片质检
	photoRes, err := strconv.Atoi(c.Ctx.Input.Query("photo_qc_res"))
	if err != nil {
		c.TaskRequest.PhotoQcResult = Const.VALID
	} else {
		switch photoRes {
		case 1:
			c.TaskRequest.PhotoQcResult = Const.VALID
		case 2:
			c.TaskRequest.PhotoQcResult = Const.INVALID
		default:
			c.TaskRequest.PhotoQcResult = Const.VALID
		}
	}

	// 模型是否有风险
	modelHasRisk, err := strconv.Atoi(c.Ctx.Input.Query("model_has_risk"))
	if err != nil {
		// 模型默认没有风险
		c.TaskRequest.ModelHasRisk = false
	} else {
		switch modelHasRisk {
		case 1:
			c.TaskRequest.ModelHasRisk = true
		case 2:
			c.TaskRequest.ModelHasRisk = false
		default:
			c.TaskRequest.ModelHasRisk = false
		}
	}

	// RunEnv
	runEnv := c.Ctx.Input.Query("run_env")
	if runEnv == "" {
		c.TaskRequest.RunEnv = "sit"
	} else {
		c.TaskRequest.RunEnv = "dev"
	}
}

func (c *TaskController) start() {
	c.init()
	c.initTaskExternal()
}

func (c *TaskController) GetTaskExternalByTaskCode() {
	c.start()

	if c.TaskRequest.TaskCode == "" {
		c.Ctx.WriteString("未输入taskI")
		return
	}

	TaskService.GetTaskExternalByTaskId(c.TaskRequest.TaskCode)

	c.Ctx.WriteString("当前输入的taskId: " + c.TaskRequest.TaskCode)
}

func (c *TaskController) GetTaskDetailInfoByTaskCode() {
	c.start()
}

// GetTaskMesQcResultJsonByTaskCode 根据任务单编号获取质检结果
func (c *TaskController) GetTaskMesQcResultJsonByTaskCode() {
	c.start()

	// 获取任务单数据
	mesQcMsg := TaskMesQcResultService.GetTaskQcResultMsg(c.TaskRequest)

	// 发送质检结果到kafka中去
	kayakService := KafkaService.KafkaService{}

	var kafkaUrl string
	if c.TaskRequest.RunEnv == "dev" {
		kafkaUrl = "kafka.dev.eainc.com:9092"
	} else {
		kafkaUrl = "10.6.34.74:9092,10.6.34.75:9092,10.6.34.73:9092"
	}

	kayakService.SetConfig([]string{kafkaUrl}, "MES_CRM_QC_RESULT_V2").Send(mesQcMsg)

	c.Ctx.WriteString("质检结果: " + mesQcMsg)
}

// GetUploadStlResult 获取任务单上传stl地址
func (c *TaskController) GetUploadStlResult() {
	c.start()

	// 获取任务单数据
	uploadReceiptStl := UploadStl.GetUploadStlResult(c.TaskRequest)

	var kafkaUrl string
	if c.TaskRequest.RunEnv == "dev" {
		kafkaUrl = "kafka.dev.eainc.com:9092"
	} else {
		kafkaUrl = "10.6.34.74:9092,10.6.34.75:9092,10.6.34.73:9092"
	}
	// 发送质检结果到kafka中去
	kayakService := KafkaService.KafkaService{}
	kayakService.SetConfig([]string{kafkaUrl}, "MES_CRM_UPLOAD_STL").Send(uploadReceiptStl)

	c.Ctx.WriteString("上传收货记录Stl地址: " + uploadReceiptStl)
}

// GetTaskReceiptByTaskCode 根据任务单编号获取收货记录
func (c *TaskController) GetTaskReceiptByTaskCode() {
	c.start()
	taskReceipt := TaskService.GetTaskReceiptByTaskId(c.TaskRequest.TaskCode)

	for key, value := range taskReceipt {
		fmt.Println("key:", key, value.ReceiptId, value.ReceiptType)
	}

	beego.Info("任务单所有的收货记录: ", taskReceipt)

	c.Ctx.WriteString("当前输入的taskId: ")
}

// GetStlQcResult 下颌前伸位质检结果
func (c *TaskController) GetStlQcResult() {

}

// GetMouseScanQcResult 下颌前伸位质检结果
func (c *TaskController) GetMouseScanQcResult() {

}

// GetSilicaQcResult 硅胶质检结果获取
func (c *TaskController) GetSilicaQcResult() {
	c.start()

	// 获取任务单数据
	mesQcMsg := TaskMesQcResultService.GetTaskQcResultMsg(c.TaskRequest)

	c.Ctx.WriteString(mesQcMsg)
}

// GetAllPassQcResult 获取质检通过的硅胶质检结果
func (c *TaskController) GetAllPassQcResult() {

}

func (c *TaskController) TestDemo() {
	c.start()

	c.Ctx.WriteString("mesQcMsg")
}
