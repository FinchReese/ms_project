package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"unicode"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type fieldInfo struct {
	Field string
	Type  string
}

type structInfo struct {
	StructName string
	FieldInfos []*fieldInfo
}

type msgInfo struct {
	MsgName    string
	FieldInfos []*fieldInfo
}

func connectMysql() (*gorm.DB, error) {
	//配置MySQL连接参数
	username := "root"      //账号
	password := "Test_1234" //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "msproject"   //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("connect mysql error: %v\n", err)
		return nil, err
	}
	return db, nil
}

func GenStruct(tableName string, structName string) {
	db, err := connectMysql()
	if err != nil {
		return
	}
	var fieldInfoList []*fieldInfo
	db.Raw(fmt.Sprintf("describe %s", tableName)).Scan(&fieldInfoList)
	for idx, fieldInf := range fieldInfoList {
		fieldInf.Type = ConvertDbFieldToStructField(fieldInf.Type)
		fieldInf.Field = ToUpperCamelCase(fieldInf.Field)
		fmt.Println(idx, fieldInf)
	}
	tmpl, err := template.ParseFiles("./struct.tpl")
	if err != nil {
		fmt.Printf("parse template error: %v\n", err)
		return
	}
	sr := structInfo{StructName: structName, FieldInfos: fieldInfoList}
	// 创建一个新文件
	file, err := os.Create("./gen/project_member.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tmpl.Execute(file, sr)
	if err != nil {
		fmt.Printf("execute template error: %v\n", err)
		return
	}
}

func ConvertDbFieldToStructField(dbField string) string {
	if strings.Contains(dbField, "bigint") {
		return "int64"
	}
	if strings.Contains(dbField, "varchar") {
		return "string"
	}
	if strings.Contains(dbField, "text") {
		return "string"
	}
	if strings.Contains(dbField, "tinyint") {
		return "int"
	}
	if strings.Contains(dbField, "int") {
		return "int"
	}
	if strings.Contains(dbField, "double") {
		return "float64"
	}
	return ""
}

func ConvertDbFieldToProtoField(dbField string) string {
	if strings.Contains(dbField, "bigint") {
		return "int64"
	}
	if strings.Contains(dbField, "varchar") {
		return "string"
	}
	if strings.Contains(dbField, "text") {
		return "string"
	}
	if strings.Contains(dbField, "tinyint") {
		return "int32"
	}
	if strings.Contains(dbField, "int") {
		return "int32"
	}
	if strings.Contains(dbField, "double") {
		return "double"
	}
	return ""
}

// ToUpperCamelCase 将字符串转换为大驼峰形式
func ToUpperCamelCase(s string) string {
	var result strings.Builder
	capitalizeNext := true
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if capitalizeNext {
				result.WriteRune(unicode.ToUpper(r))
				capitalizeNext = false
			} else {
				result.WriteRune(unicode.ToLower(r))
			}
		} else {
			capitalizeNext = true
		}
	}
	return result.String()
}

func GenProtoMsg(tableName string, msgName string) {
	db, err := connectMysql()
	if err != nil {
		return
	}
	var fieldInfoList []*fieldInfo
	db.Raw(fmt.Sprintf("describe %s", tableName)).Scan(&fieldInfoList)
	for idx, fieldInf := range fieldInfoList {
		fieldInf.Type = ConvertDbFieldToProtoField(fieldInf.Type)
		fieldInf.Field = ToUpperCamelCase(fieldInf.Field)
		fmt.Println(idx, fieldInf)
	}
	var fm template.FuncMap = make(map[string]any)
	fm["Add"] = func(v int, add int) int {
		return v + add
	}
	t := template.New("project_member.tpl")
	t.Funcs(fm)
	tmpl, err := t.ParseFiles("./protoc.tpl")
	if err != nil {
		fmt.Printf("parse template error: %v\n", err)
		return
	}
	sr := msgInfo{MsgName: msgName, FieldInfos: fieldInfoList}
	// 创建一个新文件
	file, err := os.Create("./gen/project_member.proto")
	if err != nil {
		fmt.Printf("parse template error: %v\n", err)
		return
	}
	defer file.Close()
	err = tmpl.Execute(file, sr)
	if err != nil {
		fmt.Printf("execute template error: %v\n", err)
		return
	}
}

func main() {
	GenStruct("ms_project_member", "ProjectMember")
	GenProtoMsg("ms_project_member", "ProjectMemberMessage")
}
