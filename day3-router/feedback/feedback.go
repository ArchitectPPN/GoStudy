package feedback

import "fmt"

type params struct {
	ReportId string
}

type feedBack struct {
	DesignId    string
	DesignCode  string
	DesignState int
	CaseId      string
}

func (feedBack *feedBack) checkParams(feed feedBack) {

}

func New() *feedBack {
	return &feedBack{}
}

func (feedBack *feedBack) SetFeedback(designId string, designCode string) {
	feedBack.DesignId = designId
	feedBack.DesignCode = designCode
}

func (feedBack *feedBack) ShowFeedBack() {
	fmt.Println("DesignId =", feedBack.DesignId, "DesignCode =", feedBack.DesignCode)
}
