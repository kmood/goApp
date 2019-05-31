package dao

import (
	"golang.org/x/net/context"
	"testing"
)

func TestNew(t *testing.T) {
	dao := New()
	err := dao.Ping(context.Background())
	if err != nil {
		t.Error(err)
	}
}

