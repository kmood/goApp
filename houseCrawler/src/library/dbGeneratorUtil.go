package library

import (
	"reflect"
	"strings"
)

//根据
func GeneratorTableSql(class interface{},tableName string,tableComment string) string {
	typeof := reflect.TypeOf(class)
	numField := typeof.NumField()
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
	sql += " ) COMMENT='"+tableComment+"' "
	return sql
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
