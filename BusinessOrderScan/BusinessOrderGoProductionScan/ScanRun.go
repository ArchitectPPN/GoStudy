package BusinessOrderGoProductionScan

import (
	"database/sql"
	"fmt"
)

var totalNum int

func Run(dbObj *sql.DB, limit int) {
	caseList := make(map[int]string, 20)
	caseList = GetBusinessOrderList(dbObj, limit)

	if len(caseList) == 0 {
		fmt.Printf("处理完毕!处理条数: %d \n", totalNum)
		return
	}

	businessOrder := BusinessOrderInfo{}
	receiptList := make(map[int]BusinessOrderReceipt)
	businessOrderExtInfo := BusinessOrderExtInfo{}
	businessOrderInUse := make(map[string]int, 2)
	for _, value := range caseList {
		totalNum++
		businessOrder = GetBusinessOrderGetBusinessOrderInfo(dbObj, value)
		_ = businessOrder.BusinessOrderId.Scan(value)

		// 获取业务单据加工选项
		businessOrderExtInfo = GetBusinessOrderExtInfo(dbObj, value)

		// 获取业务单据的收货记录
		receiptList = GetBusinessRelationReceipt(dbObj, value)
		// 没有数据时, 跳过
		if len(receiptList) == 0 {
			fmt.Printf("业务单据下没有收货记录 %s 业务单据编号: %s \n", value, businessOrder.BusinessOrderCode.String)
			continue
		}

		// 设置上下颌未使用
		businessOrderInUse["U"] = 0
		businessOrderInUse["L"] = 0
		for _, item := range receiptList {
			if item.inUse == 1 && item.receiptJaw.String == "L" {
				businessOrderInUse["L"] = 1
			} else if item.inUse == 1 && item.receiptJaw.String == "U" {
				businessOrderInUse["U"] = 1
			}
		}

		/**
		 * 1. 加工上颌, 上颌被使用了
		 * 2. 加工下颌, 下颌被使用了
		 * 3. 加工全颌, 上下颌均被使用
		 */
		if businessOrderExtInfo.requiredJaw.String == "3" && businessOrderInUse["L"] == 1 && businessOrderInUse["U"] == 1 {
			productionId := GetBusinessOrderProduction(dbObj, value)
			if productionId == "" {
				fmt.Printf("加工全颌, 阶段下未找到非取消的生产加工单 业务单据: %s \n", businessOrder.BusinessOrderCode.String)
			}
		} else if businessOrderExtInfo.requiredJaw.String == "1" && businessOrderInUse["U"] == 1 {
			productionId := GetBusinessOrderProduction(dbObj, value)
			if productionId == "" {
				fmt.Printf("加工上颌, 阶段下未找到非取消的生产加工单 业务单据: %s \n", businessOrder.BusinessOrderCode.String)
			}
		} else if businessOrderExtInfo.requiredJaw.String == "2" && businessOrderInUse["L"] == 1 {
			productionId := GetBusinessOrderProduction(dbObj, value)
			if productionId == "" {
				fmt.Printf("加工下颌, 阶段下未找到非取消的生产加工单 业务单据: %s \n", businessOrder.BusinessOrderCode.String)
			}
		} else {
			fmt.Printf("加工颌位 %s 阶段下没有在使用的上下颌模型 业务单据: %s \n", businessOrderExtInfo.requiredJaw.String, businessOrder.BusinessOrderCode.String)
		}
	}

	limit += 100
	Run(dbObj, limit)
}
