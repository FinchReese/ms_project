package interceptor

import (
	"context"
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"test.com/project-common/encrypt"
)

type Cache interface {
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
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
	cacheKey := encrypt.Md5(string(marshal))
	respJson, _ := s.cache.Get(cacheCtx, info.FullMethod+"::"+cacheKey)
	if respJson != "" { // 有缓存直接回复缓存
		json.Unmarshal([]byte(respJson), respCacheConfig.Resp)
		zap.L().Info(info.FullMethod + "应用缓存")
		return respCacheConfig.Resp, nil
	}
	resp, err := handler(ctx, req)
	bytes, _ := json.Marshal(resp)
	s.cache.Put(cacheCtx, info.FullMethod+"::"+cacheKey, string(bytes), respCacheConfig.Expire)
	zap.L().Info(info.FullMethod + " 放入缓存")
	return resp, err
}
