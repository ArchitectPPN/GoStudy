package ScanReceiptNoStl

import (
	"Byte/MySqlDb"
	"database/sql"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"os"
	"time"
)

var fileHandle *os.File
var ddTalk *dingtalk.DingTalk

func Handle(db *sql.DB, mysqlConfig *MySqlDb.MySqlConfig) {
	var modelHasNotStl string
	modelHasNotStl = "G:\\laragon\\www\\dnmp-master\\www\\GoStudy\\Byte\\Cache\\modelHasNotStl_" + time.Now().Format("20060102")
	// 文件创建失败, 退出
	if !createFile(modelHasNotStl) {
		return
	}

	fileHandle = openFile(modelHasNotStl)

	// 初始化dingtalk
	ddTalk = initDingTalk()
	ddSendTextMsg("开始扫描")

	startIndex := 0
	dealStageTotalNum := 0
	for {
		stageList := getStageList(db, startIndex)

		for _, tempStageId := range stageList {
			if tempStageId == "" {
				continue
			}
			//fmt.Println("阶段id:[", value, "]")
			// 获取阶段信息
			stageDetail := getStageInfo(db, tempStageId)
			// 获取阶段下的所有的收货记录
			dealStageTotalNum++
			//fmt.Println("当前处理条数: [", dealStageTotalNum, "]")
			//fmt.Println("阶段是否结束: [", stageDetail.produce.String, "]")
			if stageDetail.produce.String == "1" {
				//fmt.Println("阶段:[", value, "] 阶段已结束, 无需处理!")
				continue
			}

			// 检查病例是否结束/暂停/不收治
			caseInfoDetail := getCaseInfo(db, tempStageId)
			if caseInfoDetail.caseStatus.String == "^8^" || caseInfoDetail.caseStatus.String == "^9^" || caseInfoDetail.caseStatus.String == "^10^" {
				continue
			} else if caseInfoDetail.caseId.String == "" {
				fmt.Println("病例信息未查询到, stageId:", tempStageId)
				continue
			} else if checkCaseHasTask(db, caseInfoDetail.caseId.String, tempStageId) {
				fmt.Println("病例数据太旧, 忽略, caseId:", caseInfoDetail.caseId.String, " stageId: ", tempStageId)
				continue
			}

			stageReceiptList := getStageReceiptList(db, tempStageId)
			for _, tmpStageReceipt := range stageReceiptList {
				dealOneReceipt(tmpStageReceipt, mysqlConfig)
			}
		}

		if stageList[0] == "" {
			break
		}

		startIndex += 10000
	}

	fmt.Println("总处理阶段条数: [", dealStageTotalNum, "]")
	// 关闭文件
	err := fileHandle.Close()
	if err != nil {
		fmt.Println("关闭文件失败: ", modelHasNotStl)
		return
	}

	fmt.Println("扫描完毕: success~")
}

// 阶段下收货记录不存在上下颌时, 报警
func dealOneReceipt(stageReceiptDetail stageReceipt, mysqlConfig *MySqlDb.MySqlConfig) {
	// 硅胶条形码为空时, 跳过
	if stageReceiptDetail.upperBarCode.String == "" && stageReceiptDetail.lowerBarCode.String == "" {
		return
	}

	webUrl := getUrl(mysqlConfig)

	// 上下颌条码都为空时, 才计算
	if stageReceiptDetail.lowerCdsStl.String == "" && stageReceiptDetail.upperCdsStl.String == "" {
		hasQuestionFile := "硅胶质检通过, 但是没有stl地址; 收货记录链接: " + webUrl + stageReceiptDetail.ReceiptId.String + "\n"
		writeResultToFile(hasQuestionFile)
	}
}

// 创建文件
func createFile(fileName string) bool {
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println("文件创建失败!")
		return false
	}

	fmt.Println("文件创建成功!")

	return true
}

// 打开文件
func openFile(openFileName string) *os.File {
	file, err := os.OpenFile(openFileName, os.O_WRONLY|os.O_APPEND, 0200)
	if err != nil {
		fmt.Println("打开文件出错, err", err)
		return file
	}

	return file
}

// 将结果写入文件
func writeResultToFile(hasQuestion string) {
	_, err := fileHandle.WriteString(hasQuestion)
	if err != nil {
		fmt.Println("文件写入失败")
		return
	}
}

// 初始化钉钉
func initDingTalk() *dingtalk.DingTalk {
	dingToken := "30e727816845b120387412da50d164380802748d947cdbf6d5d948ec05901deb"
	dingSecret := "SECaf679071ac68e5a4cc38a0d2f3c0e1bef91e656a41307f51526291c878babe47"
	return dingtalk.InitDingTalkWithSecret(dingToken, dingSecret)
}

// 钉钉发送通知
func ddSendTextMsg(contentDingString string) {
	err := ddTalk.SendTextMessage(contentDingString)

	if err != nil {
		fmt.Println("发送钉钉通知失败:", err)
		return
	}
}

// 获取请求url
func getUrl(config *MySqlDb.MySqlConfig) string {
	var env string

	switch config.GetEnv() {
	case "sit":
		env = config.GetEnv() + "."
		break
	case "dev":
		env = config.GetEnv() + "."
		break
	case "prod":
		env = ""
		break
	}

	return "http://crm-web." + env + "eainc.com/crm/index.php?module=ea_receipt&action=DetailView&record="
}
