package repo

import (
	"context"
)

type Cache interface {
	// 根据key获取集合元素
	GetMembers(ctx context.Context, key string) ([]string, error)
	// 根据key清空集合
	ClearSet(ctx context.Context, key string) error
	// 根据key删除记录
	Delete(ctx context.Context, key string) error
}
