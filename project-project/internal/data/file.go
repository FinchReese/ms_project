package data

// File 文件表
type File struct {
	Id               int    `db:"id"`                // 主键
	PathName         string `db:"path_name"`         // 相对路径
	Title            string `db:"title"`             // 名称
	Extension        string `db:"extension"`         // 扩展名
	Size             int    `db:"size"`              // 文件大小
	ObjectType       string `db:"object_type"`       // 对象类型
	OrganizationCode int64  `db:"organization_code"` // 组织编码
	TaskCode         int64  `db:"task_code"`         // 任务编码
	ProjectCode      int64  `db:"project_code"`      // 项目编码
	CreateBy         int64  `db:"create_by"`         // 上传人
	CreateTime       int64  `db:"create_time"`       // 创建时间
	Downloads        int    `db:"downloads"`         // 下载次数
	Extra            string `db:"extra"`             // 额外信息
	Deleted          int    `db:"deleted"`           // 删除标记
	FileUrl          string `db:"file_url"`          // 完整地址
	FileType         string `db:"file_type"`         // 文件类型
	DeletedTime      int64  `db:"deleted_time"`      // 删除时间
}

// TableName 返回表名
func (f *File) TableName() string {
	return "ms_file"
}
