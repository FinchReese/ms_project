package kafka_log

import (
	"context"

	"go.uber.org/zap"
	"test.com/project-common/kk"
	"test.com/project-project/config"
	"test.com/project-project/internal/dao"
	"test.com/project-project/internal/repo"
)

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw = kk.GetWriter(config.AppConf.KafkaConf.Addr)
	return kw.Close
}

func SendLog(data []byte) {
	kw.Send(kk.LogData{
		Topic: "msproject_log",
		Data:  data,
	})
}

func SendCache(data []byte) {
	kw.Send(kk.LogData{
		Topic: "msproject_cache",
		Data:  data,
	})
}

type KafkaCache struct {
	reader *kk.KafkaReader
	cache  repo.Cache
}

func (c *KafkaCache) DeleteCache() {
	for {
		message, err := c.reader.ReadMsg(context.TODO())
		if err != nil {
			zap.L().Error("DeleteCache ReadMessage err", zap.Error(err))
			continue
		}
		zap.L().Info("收到缓存", zap.String("value", string(message.Value)))
		if string(message.Value) == "task" {
			fields, err := c.cache.GetMembers(context.Background(), "task")
			if err != nil {
				zap.L().Error("DeleteCache HKeys err", zap.Error(err))
				continue
			}
			for _, field := range fields {
				c.cache.Delete(context.Background(), field)
			}
			c.cache.ClearSet(context.Background(), "task")
		}
	}

}

func (c *KafkaCache) Close() {
	c.reader.Close()
}

func NewKafkaCache() *KafkaCache {
	reader := kk.GetReader([]string{config.AppConf.KafkaConf.Addr}, "cache_group", "msproject_cache")
	return &KafkaCache{
		reader: reader,
		cache:  dao.Rc,
	}
}
