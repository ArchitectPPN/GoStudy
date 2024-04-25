package QcResult

import "database/sql"

type QcResult struct {
	WorkOrderId string `json:"workOrderId"` // 任务单id
	QcResult    []QcDetail
}

type QcDetail struct {
	CouldBuccinator int        `json:"couldBuccinator"`
	QcItem          string     `json:"qcItem"`
	QcObjects       []QcItem   `json:"qcObjects"`
	QcType          string     `json:"QcType"`
	RiskDetail      RiskDetail `json:"riskDetail"`
}

type QcItem struct {
	BarCode    string `json:"barCode,omitempty"`
	CdsUrl     string `json:"cdsUrl,omitempty"`
	Jaw        string `json:"jaw,omitempty"`
	QcResult   int    `json:"qcResult"`
	QcStandard string `json:"qcStandard"`
}

type RiskDetail struct {
	PromptRisk int `json:"promptRisk"`
}

type TaskInfo struct {
	TaskName          string
	TaskId            string
	WorkType          string
	BiteQcType        string
	QcItem            string
	TaskReceiptDetail map[int]TaskReceiptInfo
}

type TaskReceiptInfo struct {
	Jaw         sql.NullString
	ReceiptType string
	ReceiptInfo sql.NullString
}
