package snowflake

import (
	"testing"
	"time"
)

func TestNewSnowflake(t *testing.T) {
	snowFlake, err := NewSnowflake(1602228331822, 1, 1)
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
