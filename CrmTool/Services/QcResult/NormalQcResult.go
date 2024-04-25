package QcResult

// QcResult 公用
type QcResult struct {
	QcResult    []QcDetail `json:"qcResults,omitempty"`
	WorkOrderId string     `json:"workOrderId,omitempty"` // 任务单id
}
