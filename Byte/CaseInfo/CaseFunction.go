package CaseInfo

func NewCaseInfo(caseCode string, product uint32) *caseInfo {
	return &caseInfo{
		caseCode: caseCode,
		product:  product,
	}
}

func (c *caseInfo) GetCaseCode() string {
	return c.caseCode
}

func (c *caseInfo) GetCaseProduct() uint32 {
	return c.product
}

func (c *caseInfo) SetCaseProduct(product uint32) {
	c.product = product
}

func (c *caseInfo) SetCaseCode(caseCode string) {
	c.caseCode = caseCode
}
