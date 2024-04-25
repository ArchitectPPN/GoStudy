package ScanReceiptNoStl

import (
	"database/sql"
	"fmt"
)

func getCaseInfo(db *sql.DB, stageId string) caseInfo {
	getCaseInfoSql := `
	SELECT 
		ec.name, 
		ec.id, 
		ecc.state_c
	FROM ea_case_ea_stage_1_c eces1c 
	INNER JOIN ea_case ec ON ec.id = eces1c.ea_case_ea_stage_1ea_case_ida AND ec.deleted = 0
	INNER JOIN ea_case_cstm ecc ON ec.id = ecc.id_c 
	WHERE 
	    eces1c.ea_case_ea_stage_1ea_stage_idb = ?
`
	var caseInfoDetail caseInfo

	err := db.QueryRow(getCaseInfoSql, stageId).Scan(&caseInfoDetail.caseCode, &caseInfoDetail.caseId, &caseInfoDetail.caseStatus)
	if err != nil {
		fmt.Println("查询病例信息出错:", err, " 阶段信息: ", stageId)
		return caseInfoDetail
	}

	return caseInfoDetail
}

// 检查病例是否有任务单
func checkCaseHasTask(db *sql.DB, caseId string, stageId string) bool {
	var taskId string
	checkTaskSql := `
	SELECT 
	    et.id 
	FROM ea_case_ea_tasks_1_c ecet1c 
	INNER JOIN ea_tasks et ON et.id = ecet1c.ea_case_ea_tasks_1ea_tasks_idb AND et.deleted = 0
	INNER JOIN ea_tasks_cstm etc ON etc.id_c = et.id AND etc.type_c IN ('1107a5', '1107a7', '1107a8', '1303b6', '1307e1') AND etc.ea_stage_id_c = ?
	WHERE ecet1c.ea_case_ea_tasks_1ea_case_ida = ?
`
	err := db.QueryRow(checkTaskSql, stageId, caseId).Scan(&taskId)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		fmt.Println("查询病例关联的任务单失败, errInfo:", err)
		return false
	}

	if taskId != "" {
		return true
	}

	return false
}
