package main

import (
	"fmt"
	"math/rand"
)

type ReceiptInfo struct {
	ReceiptId    string
	MaterialType int8
	Barcode      string
	Quality      int8
	Sort         int8
}

type StageModel struct {
	QcPassModel map[string]*ReceiptInfo
	UnQcModel   map[string]*ReceiptInfo
	UnPassModel map[string]*ReceiptInfo
}

func main() {
	var stageModel StageModel
	rand.Seed(2)
	// 设置质检合格的模型
	stageModel.QcPassModel = make(map[string]*ReceiptInfo, 10)
	stageModel.QcPassModel["uuid-opms-yu21-p9k8"] = &ReceiptInfo{"uuid-opms-yu21-p9k8", 14, "GJ20221111011", 1, int8(rand.Intn(100))}
	stageModel.QcPassModel["uuid-opms-yu21-p9k7"] = &ReceiptInfo{"uuid-opms-yu21-p9k7", 15, "GJ20221111012", 2, int8(rand.Intn(100))}
	stageModel.QcPassModel["uuid-opms-yu21-p9k6"] = &ReceiptInfo{"uuid-opms-yu21-p9k6", 21, "GJ20221111013", 1, int8(rand.Intn(100))}

	setStageReceiptModel(&stageModel)

	fmt.Println(stageModel.UnQcModel, stageModel.UnPassModel)

	for stageKey, stageValue := range stageModel.QcPassModel {
		fmt.Printf("stage_key %s receiptId: %s \n", stageKey, stageValue.ReceiptId)
	}

	receipt := getStageLastModel(stageModel)

	//fmt.Printf("stage Model %v", *stageLastModel)
	fmt.Println(receipt.ReceiptId, receipt.Quality, receipt.Barcode, receipt.Sort, receipt.MaterialType)
}

func getStageLastModel(model StageModel) *ReceiptInfo {
	var receipt *ReceiptInfo

	for _, stageModelValue := range model.QcPassModel {
		if receipt == nil {
			receipt = stageModelValue
		} else if stageModelValue != nil && receipt.Sort < stageModelValue.Sort {
			receipt = stageModelValue
		}
	}

	return receipt
}

func setStageReceiptModel(stageModel *StageModel) {
	// 设置未质检的模型
	stageModel.UnQcModel = make(map[string]*ReceiptInfo, 10)
	stageModel.UnQcModel["uuid-opms-yu21-p9k5"] = &ReceiptInfo{"uuid-opms-yu21-p9k5", 14, "GJ20221111010", 0, int8(rand.Intn(100))}
	stageModel.UnQcModel["uuid-opms-yu21-p9k4"] = &ReceiptInfo{"uuid-opms-yu21-p9k4", 15, "GJ20221111009", 0, int8(rand.Intn(100))}
	stageModel.UnQcModel["uuid-opms-yu21-p9k3"] = &ReceiptInfo{"uuid-opms-yu21-p9k3", 21, "GJ20221111008", 0, int8(rand.Intn(100))}

	// 设置质检不合格的模型
	stageModel.UnPassModel = make(map[string]*ReceiptInfo, 10)
	stageModel.UnPassModel["uuid-opms-yu21-p9k2"] = &ReceiptInfo{"uuid-opms-yu21-p9k2", 14, "GJ20221111007", 2, int8(rand.Intn(100))}
	stageModel.UnPassModel["uuid-opms-yu21-p9k1"] = &ReceiptInfo{"uuid-opms-yu21-p9k1", 15, "GJ20221111006", 2, int8(rand.Intn(100))}
	stageModel.UnPassModel["uuid-opms-yu21-p9k0"] = &ReceiptInfo{"uuid-opms-yu21-p9k0", 21, "GJ20221111005", 2, int8(rand.Intn(100))}
}
