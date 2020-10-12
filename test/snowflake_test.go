package test

import (
	"github.com/fjlyx97/short_url/services/snowflake"
	"testing"
	"time"
)

func TestNewSnowflake(t *testing.T) {
	snowFlake, err := snowflake.NewSnowflake(1602228331822, 1, 1)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond * 10)
		id, err := snowFlake.NextId()
		if err != nil {
			t.Error(err)
		}
		t.Log(id)
	}
}
