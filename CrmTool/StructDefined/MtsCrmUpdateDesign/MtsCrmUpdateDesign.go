package MtsCrmUpdateDesign

type MtsCrmUpdateDesign struct {
	CaseCode            string          `json:"caseCode"`
	StageId             string          `json:"stageId"`
	DesignOrderCode     string          `json:"designOrderCode"`
	DesignOrderStatus   string          `json:"designOrderStatus"`
	DesignId            string          `json:"designId"`
	DesignCode          string          `json:"designCode"`
	DesignStatus        string          `json:"designStatus"`
	DesignRequestId     string          `json:"design_request_id"`
	ExtraInfo           ExtraInfo       `json:"extraInfo"`
	DesignType          string          `json:"design_type"`
	NoAdmission         string          `json:"no_admission,omitempty"`
	TreatedDiagnosisC   string          `json:"treated_diagnosis_c"`
	UpperJawStepBeginC  string          `json:"upper_jaw_step_begin_c"`
	UpperJawStepC       string          `json:"upper_jaw_step_c"`
	UpperJawStepMore    string          `json:"upper_jaw_step_more,omitempty"`
	LowerJawStepBeginC  string          `json:"lower_jaw_step_begin_c"`
	LowerJawStepC       string          `json:"lower_jaw_step_c"`
	LowerJawStepMore    string          `json:"lower_jaw_step_more,omitempty"`
	JawDetails          []JawDetails    `json:"jaw_details"`
	AtdVersion          string          `json:"atdVersion"`
	AddRetentionSuccess string          `json:"addRetentionSuccess"`
	RetentionInfo       []RetentionInfo `json:"retentionInfo"`
	Jaw                 string          `json:"jaw"`
	DesignKind          string          `json:"design_kind"`
	ApplianceType       string          `json:"appliance_type"`
	Difficulty          string          `json:"difficulty"`
	CycleStart          string          `json:"cycle_start"`
	CycleEnd            string          `json:"cycle_end"`
	ProcessRemark       string          `json:"process_remark"`
	IsNeedRetainer      string          `json:"is_need_retainer"`
	HasA6Attachment     string          `json:"has_a6_attachment"`
	IsCbct              string          `json:"is_cbct"`
	HasCompensation     string          `json:"has_compensation,omitempty"`
	IsUseCompensation   string          `json:"is_use_compensation"`
	DesignOpinion       string          `json:"design_opinion,omitempty"`
	DesignOpinionJson   string          `json:"design_opinion_json"`
	RemoteCure          string          `json:"remote_cure,omitempty"`
	RemoteCureCycle     string          `json:"remote_cure_cycle"`
	Property            string          `json:"property"`
	CaseDesignType      string          `json:"case_design_type"`
	AttachmentTemplate  string          `json:"attachment_template"`
	AmendmentsItems     AmendmentsItems `json:"amendmentsItems"`
	OperationStep       string          `json:"operation_step"`
}

type ExtraInfo struct {
	Action string `json:"action"`
}

type JawDetails struct {
	Jaw   string `json:"Jaw"`
	Start string `json:"Start"`
	End   string `json:"End"`
	Type  string `json:"Type"`
}

type RetentionInfo struct {
}

type AmendmentsItems struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	HandleDescription string `json:"handle_description"`
	UserId            string `json:"user_id"`
	UserName          string `json:"user_name"`
}
