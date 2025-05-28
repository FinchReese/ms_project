package interceptor

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"test.com/project-common/encrypt"
)

type Cache interface {
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	// 添加元素到指定集合
	SetAdd(ctx context.Context, key string, value string) error
}

type RespCacheConfig struct {
	Resp   any           // 指向回复消息结构体指针
	Expire time.Duration // 缓存过期时间
}

type MethodToConfigMap map[string]RespCacheConfig

type ServiceInterceptor struct {
	methodToConfigMap MethodToConfigMap // fullmethod到缓存配置的映射
	cache             Cache             // 缓存接口
}

func NewServiceInterceptor(methodToConfigMap MethodToConfigMap, cache Cache) *ServiceInterceptor {
	return &ServiceInterceptor{
		methodToConfigMap: methodToConfigMap,
		cache:             cache,
	}
}

func (s *ServiceInterceptor) Intercept(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	respCacheConfig, ok := s.methodToConfigMap[info.FullMethod]
	if !ok { // 不需要缓存
		return handler(ctx, req)
	}
	//先查询是否有缓存，有的话直接返回，否则继续处理先请求 将回复消息存入缓存
	cacheCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	marshal, _ := json.Marshal(req)
	reqJson := encrypt.Md5(string(marshal))
	cacheKey := info.FullMethod + "::" + reqJson
	respJson, _ := s.cache.Get(cacheCtx, cacheKey)
	if respJson != "" { // 有缓存直接回复缓存
		json.Unmarshal([]byte(respJson), respCacheConfig.Resp)
		zap.L().Info(info.FullMethod + "应用缓存")
		return respCacheConfig.Resp, nil
	}
	resp, err := handler(ctx, req)
	if err != nil {
		return resp, err
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		zap.L().Error(info.FullMethod+" 放入缓存失败", zap.Error(err))
		return nil, err
	}
	err = s.cache.Put(cacheCtx, cacheKey, string(bytes), respCacheConfig.Expire)
	if err != nil {
		zap.L().Error(info.FullMethod+" 放入缓存失败", zap.Error(err))
		return nil, err
	}
	zap.L().Info(info.FullMethod + " 放入缓存成功")
	// 将缓存key存入task集合
	if strings.HasPrefix(info.FullMethod, "/task.service.v1") {
		s.cache.SetAdd(cacheCtx, "task", cacheKey)
	}
	return resp, err
}
