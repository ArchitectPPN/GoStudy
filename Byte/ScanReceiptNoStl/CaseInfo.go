package ScanReceiptNoStl

import "database/sql"

type caseInfo struct {
	caseId     sql.NullString
	caseStatus sql.NullString
	caseCode   sql.NullString
}
