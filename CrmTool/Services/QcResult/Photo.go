package QcResult

type PhotoDetail struct {
	UnqualifiedReason       string                 `json:"unqualifiedReason,omitempty"`       // 不合格原因
	MesUnqualifiedReportUrl string                 `json:"mesUnqualifiedReportUrl,omitempty"` // 不合格报告的MES下载地址
	PhotoDefectiveDetails   [8]PhotoQcDetailObject `json:"photoDefectiveDetails,omitempty"`   // 照片质检结果
	QcResult                int                    `json:"qcResult,omitempty"`
}

type PhotoQcDetailObject struct {
	Attachs             []string `json:"attachs,omitempty"`             // 附件
	DefectiveReason     string   `json:"defectiveReason,omitempty"`     // 不合格原因
	DefectiveReasonCode string   `json:"defectiveReasonCode,omitempty"` // 不合格状态码
	DefectiveRemark     string   `json:"defectiveRemark,omitempty"`     // 不合格描述
	Position            string   `json:"position,omitempty"`            // 不合格照片位置
}
