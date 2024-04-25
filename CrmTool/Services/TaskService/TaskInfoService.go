package TaskService

import (
	"CrmTool/models/Stage"
	"CrmTool/models/TaskInfo"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetTaskExternalByTaskId 获取任务单的质检信息
func GetTaskExternalByTaskId(taskId string) TaskInfo.EaTasksExtersion {
	// 初始化orm
	ormObj := orm.NewOrm()
	taskExternalInfo := TaskInfo.EaTasksExtersion{TaskId: taskId}
	err := ormObj.Read(&taskExternalInfo, "task_id")
	if err != nil {
		if err == orm.ErrNoRows {
			beego.Info("未查询到数据, err:", err)
			return taskExternalInfo
		}

		beego.Info("数据库出错, err:", err)
		return taskExternalInfo
	}

	beego.Info("查询到的数据 taskInfo:", taskExternalInfo)
	return taskExternalInfo
}

// GetTaskReceiptByTaskId 根据任务单id获取任务单关联的收货记录
func GetTaskReceiptByTaskId(tasKId string) []*TaskInfo.TaskReceipt {
	// 初始化orm
	ormObj := orm.NewOrm()

	var allTaskReceipt []*TaskInfo.TaskReceipt
	num, err := ormObj.QueryTable("task_receipt").Filter("task_id", tasKId).All(&allTaskReceipt)
	if err != nil {
		fmt.Println("Returned Rows Num: ", num, err)
		return allTaskReceipt
	}

	for key, val := range allTaskReceipt {
		fmt.Println("key: ", key, "val", val)
	}

	return allTaskReceipt
}

// GetTaskIdByCode 根据任务单Code获取任务单id
func GetTaskIdByCode(taskCode string) string {
	ormObj := orm.NewOrm()

	task := TaskInfo.EaTasks{Name: taskCode}
	err := ormObj.Read(&task, "name")
	if err != nil {
		if err == orm.ErrNoRows {
			fmt.Println("未查询到数据taskCode:", taskCode)
			return ""
		}

		fmt.Println("数据库查询出错:", err)
		return ""
	}

	fmt.Println("任务单:", task)

	return task.Id
}

// GetTaskRelationStageId 获取任务单关联阶段信息
func GetTaskRelationStageId(tasKId string) (error, Stage.EaStageCstm) {
	ormObj := orm.NewOrm()

	taskInfo := TaskInfo.EaTasksCstm{Id: tasKId}
	err := ormObj.Read(&taskInfo, "id_c")
	if err != nil {
		if err == orm.ErrNoRows {
			fmt.Println("未查询到数据taskId:", taskInfo.Id)
			return err, Stage.EaStageCstm{}
		}
	}

	// 查询阶段信息
	stageInfo := Stage.EaStageCstm{StageId: taskInfo.StageId}
	err = ormObj.Read(&stageInfo, "id_c")
	if err != nil {
		if err == orm.ErrNoRows {
			fmt.Println("未查询到数据taskId:", taskInfo.Id)
			return err, Stage.EaStageCstm{}
		}
	}

	return err, stageInfo
}
