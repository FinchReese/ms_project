package repo

import (
	"context"
)

type ProjectAuthNodeRepo interface {
	// 根据auth_id获取节点url
	GetProjectAuthNodeList(ctx context.Context, authId int64) ([]string, error)
}
