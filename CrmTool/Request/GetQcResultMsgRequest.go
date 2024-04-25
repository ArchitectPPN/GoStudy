package Request

type GetQcResultMsgRequest struct {
	TaskCode          string // 任务单编号
	TaskId            string // 任务单id
	ModelQcResult     int    // 模型质检结果 1 上颌质检通过 2 下颌质检通过 3 上下颌均质检通过
	PhotoQcResult     int    // 照片质检结果 1 质检通过 2 质检不通过
	BiteQcResult      int    // 咬合质检结果 1 质检通过 2 质检不通过
	PermanentEruption int    // 恒牙萌出质检 0 默认不存在 1 恒牙萌出 2 恒牙未萌出
	CouldBuccinator   int    // 做唇颊肌 0 不处理 1 做唇颊肌 2 不做唇颊肌
	ModelHasRisk      bool   // 模型质检是否有风险

	// 任务单质检项
	IsQcBite  bool // 任务单是否包含质检咬合
	IsQcModel bool // 任务单是否包含质检模型
	IsQcPhoto bool // 任务单是否包含质检照片

	// 环境
	RunEnv string // 执行的环境
}
