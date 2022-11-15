package queryCaseCode

import (
	"example/mysqlDb"
	"fmt"
)

type caseInfo struct {
	CrmCaseCode     string `json:"crmCaseCode"`
	CrmPatientCode  string `json:"crmPatientCode"`
	CrmPatientName  string `json:"crmPatientName"`
	CrmPatientPhone string `json:"crmPatientPhone"`
	DoctorName      string `json:"doctorName"`
	DoctorId        string `json:"doctorId"`
	ApplianceType   string `json:"applianceType"`
	BaId            string `json:"baId"`
	CrmAccountName  string `json:"crmAccountName"`
	CrmAccountId    string `json:"crmAccountId"`
}

func QueryCaseCode(caseCode string) *caseInfo {
	var caseInfo = new(caseInfo)

	queryCaseByCode(caseCode, caseInfo)

	fmt.Println(caseInfo)

	return caseInfo
}

func queryCaseByCode(caseCode string, caseInfo *caseInfo) {
	db := mysqlDb.InitDB()

	sqlStr := "	SELECT " +
		"			ea_case.name as caseCode," +
		"			ea_case_cstm.appliance_type_c as applianceType," +
		"			ea_case_cstm.ba_id_c as baId," +
		"			contacts.last_name as doctorName," +
		"			contacts.id as doctorId," +
		"			ea_patients.name as patientName," +
		"			ea_patients_cstm.code_c as patientCode," +
		"			ea_patients_cstm.phone_mobile_c as phoneMobile," +
		"			accounts.name AS accountName," +
		"			accounts.id as accountId" +
		"		FROM ea_case" +
		"		INNER JOIN ea_case_cstm ON ea_case.id = ea_case_cstm.id_c" +
		"		INNER JOIN contacts_ea_case_1_c ON contacts_ea_case_1_c.contacts_ea_case_1ea_case_idb = ea_case.id AND contacts_ea_case_1_c.deleted=0" +
		"		INNER JOIN contacts ON contacts_ea_case_1_c.contacts_ea_case_1contacts_ida = contacts.id" +
		"		INNER JOIN ea_patients_ea_case_1_c ON ea_patients_ea_case_1_c.ea_patients_ea_case_1ea_case_idb = ea_case.id AND ea_patients_ea_case_1_c.deleted=0" +
		"		INNER JOIN ea_patients ON ea_patients_ea_case_1_c.ea_patients_ea_case_1ea_patients_ida = ea_patients.id" +
		"		INNER JOIN ea_patients_cstm ON ea_patients_cstm.id_c = ea_patients.id" +
		"		INNER JOIN accounts_ea_case_1_c ON accounts_ea_case_1_c.accounts_ea_case_1ea_case_idb = ea_case.id AND accounts_ea_case_1_c.deleted=0" +
		"		INNER JOIN accounts ON accounts_ea_case_1_c.accounts_ea_case_1accounts_ida = accounts.id" +
		"		WHERE ea_case.name = ?"

	err := db.QueryRow(sqlStr, caseCode).Scan(
		&caseInfo.CrmCaseCode,
		&caseInfo.ApplianceType,
		&caseInfo.BaId,
		&caseInfo.DoctorName,
		&caseInfo.DoctorId,
		&caseInfo.CrmPatientName,
		&caseInfo.CrmPatientCode,
		&caseInfo.CrmPatientPhone,
		&caseInfo.CrmAccountName,
		&caseInfo.CrmAccountId)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
}
