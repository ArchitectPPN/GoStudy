package ScanReceiptNoStl

import "database/sql"

// 阶段
type stage struct {
	id        string
	stageName sql.NullString
}

// 阶段信息
type stageInfo struct {
	id      string
	produce sql.NullString
}

// 阶段收货记录详情
type stageReceipt struct {
	ReceiptId    sql.NullString
	MaterialType sql.NullString
	Qualified    sql.NullString
	upperCdsStl  sql.NullString
	lowerCdsStl  sql.NullString
	lowerBarCode sql.NullString
	upperBarCode sql.NullString
}
