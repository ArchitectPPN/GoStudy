package TaskInfo

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
)

// EaTasksExtersion 任务单质检信息
type EaTasksExtersion struct {
	Id            int            `orm:"column(id);pk"`
	QcItem        sql.NullString `orm:"column(qc_item)"`
	TaskId        string         `orm:"column(task_id)"`
	WorkType      string         `orm:"column(work_type)"`
	BiteQcType    string         `orm:"column(biteQcType)"`
	ReworkInfo    sql.NullString `orm:"column(rework_info)"`
	IsMatchBite   string         `orm:"column(is_match_bite)"`
	MuscleTrainer int            `orm:"column(muscleTrainer)"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(EaTasksExtersion))
}
