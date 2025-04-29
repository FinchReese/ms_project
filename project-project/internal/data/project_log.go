package data

// ProjectLog 项目日志表
type ProjectLog struct {
	Id           int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`                 // 主键ID
	MemberCode   int64  `json:"memberCode" gorm:"column:member_code;default:0"`               // 操作人id
	Content      string `json:"content" gorm:"column:content;type:text"`                      // 操作内容
	Remark       string `json:"remark" gorm:"column:remark;type:text"`                        // 备注
	Type         string `json:"type" gorm:"column:type;type:varchar(255);default:create"`     // 操作类型
	CreateTime   int64  `json:"createTime" gorm:"column:create_time"`                         // 添加时间
	SourceCode   int64  `json:"sourceCode" gorm:"column:source_code;default:0"`               // 任务id
	ActionType   string `json:"actionType" gorm:"column:action_type;type:varchar(30)"`        // 场景类型
	ToMemberCode int64  `json:"toMemberCode" gorm:"column:to_member_code;default:0"`          // 目标成员id
	IsComment    int8   `json:"isComment" gorm:"column:is_comment;type:tinyint(1);default:0"` // 是否评论，0：否
	ProjectCode  int64  `json:"projectCode" gorm:"column:project_code"`                       // 项目编码
	Icon         string `json:"icon" gorm:"column:icon;type:varchar(20)"`                     // 图标
	IsRobot      int8   `json:"isRobot" gorm:"column:is_robot;type:tinyint(1);default:0"`     // 是否机器人
}

// TableName 指定表名
func (ProjectLog) TableName() string {
	return "ms_project_log"
}
