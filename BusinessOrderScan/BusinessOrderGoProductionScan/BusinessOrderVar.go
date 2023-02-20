package BusinessOrderGoProductionScan

import "database/sql"

type CaseInfo struct {
	CaseId       sql.NullString
	CaseCode     sql.NullString
	TeenagerInfo sql.NullString
	CaseState    sql.NullString
}

type BusinessOrderInfo struct {
	BusinessOrderCode   sql.NullString
	BusinessOrderId     sql.NullString
	BusinessOrderStatus sql.NullString
}

type BusinessOrderExtInfo struct {
	requiredJaw sql.NullString
}

type BusinessOrderReceipt struct {
	receiptId     sql.NullString
	receiptType   sql.NullString
	receiptJaw    sql.NullString
	qualification int
	goNext        int
	inUse         int
}
