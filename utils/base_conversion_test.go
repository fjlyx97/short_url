package utils

import "testing"

func TestBase10ToBase62(t *testing.T) {
	var num int64 = 1238721408124
	b62 := Base10ToBase62(num)
	b10 := Base62ToBase10(b62)
	t.Logf("b62 : %s , b10 : %d", b62, b10)
	if b10 != num {
		t.Error("base conversion error")
	}
}
