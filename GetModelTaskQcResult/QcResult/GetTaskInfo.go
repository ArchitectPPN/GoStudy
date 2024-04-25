package QcResult

import (
	"database/sql"
	"fmt"
)

func GetTaskInfoByTaskName(dbObj *sql.DB, taskName string) TaskInfo {
	var taskInfo TaskInfo
	getTaskInfoSql := "SELECT id, name FROM ea_tasks WHERE name = ? AND deleted = 0"
	err := dbObj.QueryRow(getTaskInfoSql, taskName).Scan(&taskInfo.TaskId, &taskInfo.TaskName)
	if err != nil {
		return taskInfo
	}

	return taskInfo
}

func GetTaskQcInfo(dbObj *sql.DB, info *TaskInfo) {
	getTaskQcInfo := "SELECT work_type,biteQcType,qc_item FROM ea_tasks_extersion WHERE task_id = ?"
	err := dbObj.QueryRow(getTaskQcInfo, info.TaskId).Scan(&info.WorkType, &info.BiteQcType, &info.QcItem)
	if err != nil {
		return
	}
}

func GetTaskReceipt(dbObj *sql.DB, info *TaskInfo) {
	getTaskReceiptSql := "SELECT jaw, receipt_type, receipt_info FROM task_receipt WHERE task_id = ?"
	rows, err := dbObj.Query(getTaskReceiptSql, info.TaskId)
	if err != nil {
		return
	}
	var taskReceipt TaskReceiptInfo
	info.TaskReceiptDetail = make(map[int]TaskReceiptInfo, 5)
	index := 0
	for rows.Next() {
		err := rows.Scan(&taskReceipt.Jaw, &taskReceipt.ReceiptType, &taskReceipt.ReceiptInfo)
		if err != nil {
			fmt.Printf("scan faled, err:%v \n", err)
		}

		info.TaskReceiptDetail[index] = taskReceipt

		index++
	}
}
