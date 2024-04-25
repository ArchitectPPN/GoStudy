package ProductionService

import (
	"CrmTool/Request/ProductionRequest"
	"CrmTool/models/Case"
	"CrmTool/models/Production"
	"CrmTool/models/Stage"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type CreateProductionService struct {
}

// CreateProduction 创建生产加工单
func (create *CreateProductionService) CreateProduction(request *ProductionRequest.ProductionRequest) {
	ormObj := orm.NewOrm()

	// 根据阶段查询病例id
	stageCase := Stage.EaCaseEaStage_1_c{StageId: request.StageId}
	err := ormObj.Read(&stageCase, "ea_case_ea_stage_1ea_stage_idb")
	if err != nil {
		fmt.Println("阶段没有对应的病例")
		return
	}
	fmt.Println("阶段对应的病例为: " + stageCase.CaseId)

	// 根据病例id查询当前的产品编号
	caseInfo := Case.EaCaseCstm{Id: stageCase.CaseId}
	err = ormObj.Read(&caseInfo, "id_c")
	if err != nil {
		fmt.Println("病例不存在")
		return
	}

	fmt.Println("执行的结果, 病例信息: ", caseInfo)

	// 开始创建生产加工单

	// 获取当前时间
	t := time.Now()

	// 获取主键id
	pId := uuid.New().String()

	// 获取生产加工单编号
	productionNo := "M" + t.Format("2006") + strconv.FormatInt(t.Unix(), 10)

	// 构建主表数据
	productionMain := Production.EaProduction{
		Id:           pId,
		DateEntered:  t.Format("2006-01-02 15:04:05"),
		DateModified: t.Format("2006-01-02 15:04:05"),
		Name:         productionNo,
	}

	applianceType := request.ApplianceType
	if applianceType == "" {
		applianceType = caseInfo.ApplianceType
	}

	// 构建副表数据
	productionCstm := Production.EaProductionCstm{
		Id:                  pId,
		ExpireDate:          t.Format("2006-01-02"),
		EaCaseId:            caseInfo.Id,
		Products:            applianceType,
		State:               "1",
		OdsAddr:             "406e0c34-b62a-a1ae-bf55-6448c37a1259",
		Num:                 "2",
		ProduceRemark:       "api创建",
		Province:            "450000",
		City:                "450100",
		Area:                "450103",
		PostalCode:          "530000",
		Contact:             "王**内",
		Telephone:           "154****2451",
		Address:             "********",
		ModelType:           "4",
		Tag:                 "1",
		Jaw:                 "3",
		Proportion:          "1",
		UpperSiliconBarcode: "http://cds.dev.eainc.com:8001/shidaits/rest/cds/case/document/73579c55-fbab-4bfc-abda-cc500baa1d00/download",
		LowerSiliconBarcode: "http://cds.dev.eainc.com:8001/shidaits/rest/cds/case/document/da6aef78-ead9-4609-8d74-27c1664d4f87/download",
	}

	// 生产加工单详情
	productionDetail := Production.EaProductiondetail{
		Id:             pId,
		Name:           productionNo + "-1",
		DateEntered:    t.Format("2006-01-02 15:04:05"),
		DateModified:   t.Format("2006-01-02 15:04:05"),
		ModifiedUserId: "1",
		Description:    "",
		CreatedBy:      "1",
		AssignedUserId: "1",
	}

	// 生产加工单详情信息
	productionDetailCstm := Production.EaProductiondetailCstm{
		Id:                pId,
		IsBuchaBcq:        "0",
		UpperTimes:        "1",
		ModelType:         "4",
		Product:           applianceType,
		UpperJawEndStep:   "",
		Thickness:         "0.76",
		LowerTimes:        "1",
		LowerJawEndStep:   "",
		UpperJawBeginStep: "",
		LowerJawBeginStep: "",
		Material:          "Angelalign Retainer",
		CrudeMaterial:     "MC",
		WearStep:          "",
		ProductType:       "",
		Scheme:            "",
		Tag:               "",
		Cnt:               "",
	}

	// 生产加工单关联
	productionDetailRelation := Production.EaProductionEaProductiondetail_1_c{
		Id:                  pId,
		DateModified:        t.Format("2006-01-02 15:04:05"),
		Deleted:             "0",
		ProductionIda:       pId,
		ProductionDetailIdb: pId,
	}

	// 获取病例最新的阶段
	var caseLastStageId string
	err = ormObj.Raw("SELECT ea_case_ea_stage_1ea_stage_idb FROM ea_case_ea_stage_1_c WHERE ea_case_ea_stage_1ea_case_ida = ? AND deleted = 0 ORDER BY date_modified DESC", stageCase.CaseId).QueryRow(&caseLastStageId)
	if err != nil {
		fmt.Println("错误:", err)
	}

	fmt.Println("病例id:", stageCase.CaseId, " stageId:", caseLastStageId)

	productionStage := Production.EaStageEaProduction_1_c{
		Id:           pId,
		DateModified: t.Format("2006-01-02 15:04:05"),
		StageId:      caseLastStageId,
		ProductionId: pId,
	}

	_, err = ormObj.Insert(&productionMain)
	_, err = ormObj.Insert(&productionCstm)
	_, err = ormObj.Insert(&productionStage)
	_, err = ormObj.Insert(&productionDetail)
	_, err = ormObj.Insert(&productionDetailCstm)
	_, err = ormObj.Insert(&productionDetailRelation)
}
