package orm

import "testing"

func TestNewMySQL(t *testing.T) {
	config := &Config{
		DSN:         "root:123456@tcp(localhost:3307)/test?charset=utf8",
		Active:      10,
		Idle:        5,
		IdleTimeout :1000,
	}
	db := NewMySQL(config)
	ping := db.DB().Ping()
	if ping != nil {
		t.Error(ping)
	}
}
