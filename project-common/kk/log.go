package kk

import (
	"encoding/json"
	"test.com/project-common/time_format"
	"time"
)

type FieldMap map[string]any

type KafkaLog struct {
	Type     string
	Action   string
	Time     string
	Msg      string
	Field    FieldMap
	FuncName string
}

func Error(err error, funcName string, fieldMap FieldMap) []byte {
	kl := KafkaLog{
		Type:     "error",
		Action:   "click",
		Time:     time_format.ConvertTimeToString(time.Now()),
		Msg:      err.Error(),
		Field:    fieldMap,
		FuncName: funcName,
	}
	bytes, _ := json.Marshal(kl)
	return bytes
}

func Info(msg string, funcName string, fieldMap FieldMap) []byte {
	kl := KafkaLog{
		Type:     "info",
		Action:   "click",
		Time:     time_format.ConvertTimeToString(time.Now()),
		Msg:      msg,
		Field:    fieldMap,
		FuncName: funcName,
	}
	bytes, _ := json.Marshal(kl)
	return bytes
}
