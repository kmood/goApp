package utils

import (
	"testing"
)

func TestConverTableName(t *testing.T) {
	tableName := "TestTestTest"
	name := ConverTableName(tableName)
	t.Log(name)
}
func TestGeneratorTableSql(t *testing.T) {

}
