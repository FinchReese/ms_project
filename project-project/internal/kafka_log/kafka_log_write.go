package kafka_log

import (
	"test.com/project-common/kk"
	"test.com/project-project/config"
)

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw = kk.GetWriter(config.AppConf.KafkaConf.Addr)
	return kw.Close
}

func SendLog(data []byte) {
	kw.Send(kk.LogData{
		Topic: config.AppConf.KafkaConf.Topic,
		Data:  data,
	})
}
