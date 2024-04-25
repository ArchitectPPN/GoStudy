package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	getQcRes := NewGetQcRes()
	getQcRes.GetQcRes()
}

type GetQcRes struct {
	mu sync.Mutex
}

func NewGetQcRes() *GetQcRes {
	return &GetQcRes{}
}

func (g *GetQcRes) GetQcRes() {
	var wg sync.WaitGroup

	qc := &Res{
		WorkId: "o989-posi-jh782-shnl0-2iuh",
		Name:   "T20240425001",
	}

	// 组装详情
	wg.Add(1)
	go g.assembleDetail(qc, &wg)

	// 组装质检结果
	wg.Add(1)
	go g.assembleQcRes(qc, &wg)

	// 组装收获记录列表
	wg.Add(1)
	go g.assembleReceiptList(qc, &wg)
	wg.Add(1)
	go g.assembleReceiptList(qc, &wg)

	// 等待所有的协程操作结束
	wg.Wait()

	jsonData, err := json.Marshal(qc)
	if err != nil {
		fmt.Println("转json失败errMsg：", err)

		return
	}
	fmt.Println("质检结果：", string(jsonData))
}

// assembleDetail 组装详情
func (g *GetQcRes) assembleDetail(qc *Res, wg *sync.WaitGroup) {
	// 当协程结束时，减少WaitGroup计数
	defer wg.Done()
	detail := &Detail{
		Id:   "uiuu-9okj-plok-98nh-gbf7",
		Type: "1920c1i",
		Name: "T20240428900",
	}

	qc.Detail = detail
}

// assembleQcRes 组装质检结果
func (g *GetQcRes) assembleQcRes(qc *Res, wg *sync.WaitGroup) {
	// 当协程结束时，减少WaitGroup计数
	defer wg.Done()

	qcRes := &QcRes{
		Id:      "uijh-9ujh-polk-9uhy",
		Quailed: "pass",
	}

	qc.QcRes = qcRes
}

// assembleReceiptList 组装收获记录列表
func (g *GetQcRes) assembleReceiptList(qc *Res, wg *sync.WaitGroup) {
	defer wg.Done()

	// 加锁
	g.mu.Lock()
	defer g.mu.Unlock()

	for i := 0; i < 100; i++ {
		receipt := &Receipt{Id: "polk-joke-p98i-jgb8", MaterialType: "口扫"}
		qc.ReceiptList = append(qc.ReceiptList, receipt)
	}
}

type Res struct {
	WorkId      string     `json:"workId"`
	Name        string     `json:"name"`
	Detail      *Detail    `json:"detail"`
	QcRes       *QcRes     `json:"qcRes"`
	ReceiptList []*Receipt `json:"receiptList"`
}

type Detail struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type QcRes struct {
	Id      string `json:"id"`
	Quailed string `json:"quailed"`
}

type Receipt struct {
	Id           string `json:"id"`
	MaterialType string `json:"materialType"`
}
