package utils

import (
	"reflect"
	"strings"
	"unicode"
)

// GeneratorTableSql 根据结构体生成sql ，默认情况下以大写字母前添加‘_’分隔
func GeneratorTableSql(class interface{}) string {
	typeof := reflect.TypeOf(class)
	numField := typeof.NumField()
	name := typeof.Name()
	tableName := ConverTableName(name)
	sql := "create table " + tableName + " ( "
	for i := 0; i < numField; i++ {
		structField := typeof.Field(i)
		fieldName := strings.ToLower(structField.Name)
		typeName := structField.Type.Name()
		comment_size := structField.Tag.Get("comment_size")
		csArr := strings.Split(comment_size, "_")
		comment := csArr[0]
		sql += fieldName + " " + getDbDataType(typeName)
		if len(csArr[1]) != 0 {
			sql += "(" + csArr[1] + ")  "
		}
		sql += " comment '" + comment + "',"
	}
	sql = strings.TrimSuffix(sql, ",")
	sql += " ) "
	return sql
}

func ConverTableName(name string)string {
	var nameNew = ""
	for i, str := range name {
		//fmt.Println(reflect.TypeOf(str))
		if i == 0 {
			nameNew += string(str)
			continue
		}
		if unicode.IsUpper(str) {
			nameNew += "_" + strings.ToLower(string(str))
		} else {
			nameNew += string(str)
		}
	}
	tableName := strings.ToLower(nameNew)
	return tableName
}
func getDbDataType(typeName string) string {
	fieldType := "varchar"
	switch typeName {
	case "Int", "Int8", "Int16", "Int32", "Int64":
		fieldType = "integer"
		break
	case "Float32", "Float64":
		fieldType = "decimal"
		break
	}
	return fieldType
}
