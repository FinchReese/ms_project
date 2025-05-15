package repo

import (
	"context"

	"gorm.io/gorm"
)

type ProjectAuthNodeRepo interface {
	// 根据auth_id获取节点url
	GetProjectAuthNodeList(ctx context.Context, authId int64) ([]string, error)
	// 删除指定auth id的节点（支持事务）
	DeleteProjectAuthNode(ctx context.Context, authId int64, db *gorm.DB) error
	// 给指定auth id的节点添加节点列表（支持事务）
	AddProjectAuthNode(ctx context.Context, authId int64, nodeList []string, db *gorm.DB) error
}
